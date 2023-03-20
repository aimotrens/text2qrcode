package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	docs "github.com/aimotrens/text2qrcode/docs"
	"github.com/aimotrens/text2qrcode/text2qrcode"
)

func main() {
	docs.SwaggerInfo.Title = "QR-Code Generator"
	docs.SwaggerInfo.BasePath = "/api"

	r := gin.Default()
	text2qrcode.SetRoutes(r)

	fmt.Println("Text2QRCode gestartet.")
	r.Run(":8080")
}
