package main

import (
	"fmt"

	"github.com/aimotrens/text2qrcode/app"
	_ "github.com/aimotrens/text2qrcode/docs"
)

// @title Text2QRCode / QR-Code Generator API
// @description Ein einfacher Konverter um einen QRCode aus Text zu erstellen.
// @BasePath /api

// @contact.name aimotrens
// @contact.url https://github.com/aimotrens/text2qrcode

func main() {
	r := app.Setup()

	fmt.Println("Text2QRCode gestartet.")
	r.Run(":8080")
}
