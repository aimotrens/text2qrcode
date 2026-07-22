package internal

import (
	"net/http"

	"github.com/aimotrens/text2qrcode/internal/healthcheck"
	"github.com/aimotrens/text2qrcode/internal/text2qrcode"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler("/swagger/index.html", http.StatusTemporaryRedirect))
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("GET /api/healthcheck", healthcheck.HealthCheck())
	mux.HandleFunc("GET /api/text2qrcode/encode", text2qrcode.EncodeWithQueryString())
	mux.HandleFunc("POST /api/text2qrcode/encode", text2qrcode.Encode())

	return mux
}
