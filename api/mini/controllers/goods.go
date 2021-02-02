package controllers

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/services/goods"
	"strconv"
)

func (c *Controller) GoodsKeywordList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")

	if keyword == "" {
		ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": nil})
		return
	}

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
	data, err, emp := service.GoodsNumber(number)
	if emp {
		ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": nil})
		return
	}
	if err != nil {
		ctx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": data})
}

func (c *Controller) GoodsUpdateNumber(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	number := ctx.PostForm("number")

	if err != nil {
		ctx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	service := goods.GoodService{Ctx: ctx}
	err = service.UpdateNumber(id, number)

	if err != nil {
		ctx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok"})
}

func (c *Controller) GoodsAdd(ctx *gin.Context) {
	var form goods.AddForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	go func() {
		service := goods.GoodService{Ctx: ctx}
		service.GoodsAdd(form)
	}()

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok"})
}

func (c *Controller) GoodsAttr(ctx *gin.Context) {
	service := goods.GoodService{Ctx: ctx}
	attr := service.GoodsAttr()
	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": attr})
}
