package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	qrcode "github.com/skip2/go-qrcode"

	docs "github.com/aimotrens/text2qrcode/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type qrcodeRequest struct {
	Text   string
	Ecl    string
	Height int
	Width  int
}

func main() {
	docs.SwaggerInfo.Title = "QR-Code Generator"
	docs.SwaggerInfo.BasePath = "/api"

	r := gin.Default()

	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/healthcheck", HealthCheck)
		t2q := api.Group("/text2qrcode")
		{
			t2q.GET("/encode", EncodeWithPathParam)
			t2q.POST("/encode", Encode)
		}
	}

	fmt.Println("Html2Pdf started.")
	r.Run(":8080")
}

// Convert HTML to PDF
// @Summary Kodiert den übergebenen Text in einen QR-Code mit 500x500px und ECL=M
// @Param text query string true "Text"
// @Produce  image/png
// @Success 200 {file} binary
// @Success 500 {string} string
// @Router /text2qrcode/encode [get]
func EncodeWithPathParam(g *gin.Context) {
	encode(g, g.Query("text"))
}

// Convert HTML to PDF
// @Summary Kodiert den übergebenen Text mit den angegebenen Parameter in einen QR-Code
// @Param request body string true "Text"
// @Accept text/plain
// @Produce  image/png
// @Success 200 {file} binary
// @Success 500 {string} string
// @Router /text2qrcode/encode [post]
func Encode(g *gin.Context) {
	data := qrcodeRequest{}
	err := g.BindJSON(data)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}
	encode(g, data.Text)
}

func encode(g *gin.Context, text string) {
	png, err := qrcode.Encode(text, qrcode.Medium, 500)

	if err != nil {
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	g.Data(http.StatusOK, "image/png", png)
}

// Healthcheck
// @Summary Gibt immer "OK" zurück
// @Produce  text/plain
// @Success 200 {string} OK
// @Router /healthcheck [get]
func HealthCheck(g *gin.Context) {
	g.String(http.StatusOK, "OK")
}
