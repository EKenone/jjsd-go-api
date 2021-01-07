package controllers

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/services/wx_fans"
)

func (c *Controller) GetSession(ctx *gin.Context) {
	code := ctx.Query("code")

	authObj := c.WcMiniInit().GetAuth()
	res, err := authObj.Code2Session(code)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 401, "msg": err.Error()})
		return
	}

	service := wx_fans.WxFansService{
		Ctx: ctx,
	}
	session := service.GetSession(res.OpenID, res.UnionID, res.SessionKey)

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": session})
}

func (c *Controller) SetUserInfo(ctx *gin.Context) {
	encryptedData := ctx.PostForm("encrypted_data")
	iv := ctx.PostForm("iv")
	//signature := ctx.PostForm("signature")
	sessionKey := ctx.PostForm("session_key")

	authObj := c.WcMiniInit().GetEncryptor()
	data, err := authObj.Decrypt(sessionKey, encryptedData, iv)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 401, "msg": err.Error()})
		return
	}

	service := wx_fans.WxFansService{
		Ctx: ctx,
	}

	err = service.SetUserInfo(data.NickName, data.Gender, data.Province, data.City, data.AvatarURL)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok"})
}
