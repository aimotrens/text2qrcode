package text2qrcode

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// @Tags Text2QRCode
// Convert HTML to PDF
// @Summary Kodiert den übergebenen Text in einen QR-Code. Die Parameter werden als Query-Parameter übergeben.
// @Param text query string true "Text"
// @Param errorCorrection query int false "ErrorCorrection" default(1)
// @Param size query int false "Size" default(250)
// @Param whiteBorder query bool false "WhiteBorder" default(true)
// @Produce  image/png
// @Success 200 {file} binary
// @Success 500 {string} string
// @Router /text2qrcode/encode [get]
func EncodeWithQueryString(g *gin.Context) {
	errorCorrection, err := strconv.Atoi(g.DefaultQuery("errorCorrection", "1"))
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	size, err := strconv.Atoi(g.DefaultQuery("size", "250"))
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	whiteBorder, err := strconv.ParseBool(g.DefaultQuery("whiteBorder", "true"))
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	encode(g, QRCodeRequest{
		Text:            g.Query("text"),
		ErrorCorrection: errorCorrection,
		Size:            size,
		WhiteBorder:     whiteBorder,
	})
}

// @Tags Text2QRCode
// @Summary Kodiert den übergebenen Text in einen QR-Code. Die Parameter werden als JSON übergeben.
// @Param request body QRCodeRequest true "qrcodeRequest"
// @Accept aplication/json
// @Produce  image/png
// @Success 200 {file} binary
// @Success 500 {string} string
// @Router /text2qrcode/encode [post]
func Encode(g *gin.Context) {
	qrReq := QRCodeRequest{}
	err := g.BindJSON(&qrReq)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	encode(g, qrReq)
}

func encode(g *gin.Context, qrReq QRCodeRequest) {
	if err := validateInput(qrReq); err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	qr, err := qrcode.New(qrReq.Text, qrcode.RecoveryLevel(qrReq.ErrorCorrection))
	if err != nil {
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	qr.DisableBorder = !qrReq.WhiteBorder

	png, err := qr.PNG(qrReq.Size)
	if err != nil {
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	g.Data(http.StatusOK, "image/png", png)
}
