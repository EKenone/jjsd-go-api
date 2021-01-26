package controllers

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/services/goods_format"
)

func (c *Controller) GoodsFormatList(ctx *gin.Context) {
	service := goods_format.GoodFormatService{Ctx: ctx}
	format := service.List()

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": format})
}
