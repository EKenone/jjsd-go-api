package controllers

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"jjsd-go-api/api/mini/conf"
)

type Controller struct {
}

func (c *Controller) WcMiniInit() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemcache()
	cfg := &miniConfig.Config{
		AppID:     conf.Conf.Wechat.Appid,
		AppSecret: conf.Conf.Wechat.Secret,
		Cache:     memory,
	}

	return wc.GetMiniProgram(cfg)
}
