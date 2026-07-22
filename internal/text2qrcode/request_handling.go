package text2qrcode

import (
	"net/http"
	"strconv"

	"github.com/aimotrens/text2qrcode/pkg/api"
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
func EncodeWithQueryString() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		strErrorCorrection := r.URL.Query().Get("errorCorrection")
		if strErrorCorrection == "" {
			strErrorCorrection = "1"
		}

		errorCorrection, err := strconv.Atoi(strErrorCorrection)
		if err != nil {
			api.JsonErr(w, http.StatusBadRequest, err)
			return
		}

		strSize := r.URL.Query().Get("size")
		if strSize == "" {
			strSize = "250"
		}

		size, err := strconv.Atoi(strSize)
		if err != nil {
			api.JsonErr(w, http.StatusBadRequest, err)
			return
		}

		strWhiteBorder := r.URL.Query().Get("whiteBorder")
		if strWhiteBorder == "" {
			strWhiteBorder = "true"
		}

		whiteBorder, err := strconv.ParseBool(strWhiteBorder)
		if err != nil {
			api.JsonErr(w, http.StatusBadRequest, err)
			return
		}

		encode(w, QRCodeRequest{
			Text:            r.URL.Query().Get("text"),
			ErrorCorrection: errorCorrection,
			Size:            size,
			WhiteBorder:     whiteBorder,
		})
	}
}

// @Tags Text2QRCode
// @Summary Kodiert den übergebenen Text in einen QR-Code. Die Parameter werden als JSON übergeben.
// @Param request body QRCodeRequest true "qrcodeRequest"
// @Accept aplication/json
// @Produce  image/png
// @Success 200 {file} binary
// @Success 500 {string} string
// @Router /text2qrcode/encode [post]
func Encode() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		qrReq := QRCodeRequest{}
		err := api.BindJson(r, &qrReq)
		if err != nil {
			api.JsonErr(w, http.StatusBadRequest, err)
			return
		}

		encode(w, qrReq)
	}
}

func encode(w http.ResponseWriter, qrReq QRCodeRequest) {
	if err := validateInput(qrReq); err != nil {
		api.JsonErr(w, http.StatusBadRequest, err)
		return
	}

	qr, err := qrcode.New(qrReq.Text, qrcode.RecoveryLevel(qrReq.ErrorCorrection))
	if err != nil {
		api.JsonErr(w, http.StatusInternalServerError, err)
		return
	}

	qr.DisableBorder = !qrReq.WhiteBorder

	png, err := qr.PNG(qrReq.Size)
	if err != nil {
		api.JsonErr(w, http.StatusInternalServerError, err)
		return
	}

	api.Data(w, http.StatusOK, "image/png", png)
}
