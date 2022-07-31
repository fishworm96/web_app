package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})
	r.POST("/signUp", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	return r
}
