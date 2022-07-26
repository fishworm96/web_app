package routers

import (
	"net/http"

	"web_app/settings"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecover(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}