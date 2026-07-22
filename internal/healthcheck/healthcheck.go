package healthcheck

import (
	"net/http"

	"github.com/aimotrens/text2qrcode/pkg/api"
)

// @Tags Healthcheck
// @Summary Gibt immer "OK" zurück
// @Produce  text/plain
// @Success 200 {string} OK
// @Router /healthcheck [get]
func HealthCheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := api.JsonOK(w, "OK")
		if err != nil {
			api.JsonErr(w, http.StatusInternalServerError, err)
		}
	}
}
