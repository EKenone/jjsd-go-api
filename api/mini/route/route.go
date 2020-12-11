package route

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/controllers"
	"jjsd-go-api/api/mini/middleware"
)

func Init(eg *gin.Engine) *gin.Engine {
	c := controllers.Controller{}
	eg.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "OK"})
		return
	})
	eg.GET("wx-fans/get-session", c.GetSession)

	eg = middleware.Init(eg)

	eg.POST("wx-fans/set-user-info", c.SetUserInfo)

	return eg
}
