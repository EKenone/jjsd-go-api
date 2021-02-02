package route

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/controllers"
	"jjsd-go-api/api/mini/middleware"
)

func Init(eg *gin.Engine) *gin.Engine {
	c := controllers.Controller{}
	eg.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "OK"})
		return
	})

	eg.GET("wx-fans/get-session", c.GetSession)
	eg = middleware.Init(eg)

	eg.GET("goods/number", c.GoodsNumber)
	eg.GET("goods/keyword-list", c.GoodsKeywordList)
	eg.POST("goods/update-number", c.GoodsUpdateNumber)

	eg.POST("wx-fans/set-user-info", c.SetUserInfo)
	eg.GET("goods/attr", c.GoodsAttr)
	eg.POST("goods/add", c.GoodsAdd)

	return eg
}
