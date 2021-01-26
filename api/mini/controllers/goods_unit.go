package controllers

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/services/goods_unit"
)

func (c *Controller) GoodsUnitList(ctx *gin.Context) {
	service := goods_unit.GoodUnitService{Ctx: ctx}
	unit := service.List()

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": unit})
}
