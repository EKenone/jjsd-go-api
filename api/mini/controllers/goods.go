package controllers

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/services/goods"
)

func (c *Controller) GoodsKeywordList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	service := goods.GoodService{Ctx: ctx}
	data, err, emp := service.KeywordList(keyword)

	if emp {
		ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": nil})
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{"code": 200, "msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": data})
}

func (c *Controller) GoodsNumber(ctx *gin.Context) {
	number := ctx.Query("number")
	service := goods.GoodService{Ctx: ctx}
	data, err, emp := service.Number(number)
	if emp {
		ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": nil})
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{"code": 200, "msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": data})
}
