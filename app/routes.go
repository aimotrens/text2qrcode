package app

import (
	"net/http"

	"github.com/aimotrens/text2qrcode/app/healthcheck"
	"github.com/aimotrens/text2qrcode/app/text2qrcode"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.Default()

	// Error Handling, Ausgabe der Fehler als JSON an den Client
	r.Use(func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		// Status 0, da der Status bereits gesetzt wurde
		c.JSON(0, c.Errors)
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	{
		hc := api.Group("/healthcheck")
		{
			hc.GET("/", healthcheck.HealthCheck)
		}

		t2q := api.Group("/text2qrcode")
		{
			t2q.GET("/encode", text2qrcode.EncodeWithQueryString)
			t2q.POST("/encode", text2qrcode.Encode)
		}
	}

	return r
}
