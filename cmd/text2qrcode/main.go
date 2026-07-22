package main

import (
	"fmt"
	"net/http"

	_ "github.com/aimotrens/text2qrcode/docs"
	"github.com/aimotrens/text2qrcode/internal"
)

// @title Text2QRCode / QR-Code Generator API
// @description Ein einfacher Konverter um einen QRCode aus Text zu erstellen.
// @BasePath /api

// @contact.name aimotrens
// @contact.url https://github.com/aimotrens/text2qrcode
func main() {
	mux := internal.Setup()

	fmt.Println("Text2QRCode gestartet.")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
