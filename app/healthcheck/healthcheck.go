package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Healthcheck
// @Summary Gibt immer "OK" zurück
// @Produce  text/plain
// @Success 200 {string} OK
// @Router /healthcheck [get]
func HealthCheck(g *gin.Context) {
	g.String(http.StatusOK, "OK")
}
