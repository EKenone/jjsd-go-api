package wx_fans

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/cache"
	"jjsd-go-api/api/mini/models"
	"strconv"
	"time"
)

type WxFansService struct {
	Ctx *gin.Context
}

type Session struct {
	Token      string `json:"token"`
	FansId     uint   `json:"fans_id"`
	SessionKey string `json:"session_key"`
}

func (s *WxFansService) GetSession(openid string, unionid string, sessionKey string) Session {
	db := models.DbLink()

	fans := models.WxFans{
		Openid:  openid,
		Unionid: unionid,
	}

	db.Debug().Where("openid = ? AND is_del = ?", openid, 0).FirstOrCreate(&fans)
	tokenStr := strconv.Itoa(int(fans.ID)) + "_" + fans.Openid
	data := []byte(tokenStr)
	tokenByte := sha1.Sum(data)

	token := hex.EncodeToString(tokenByte[:])
	fansJson, _ := json.Marshal(fans)

	var ctx = context.Background()
	redis := cache.GetClient()
	redis.Set(ctx, cache.PrefixKey(token), fansJson, time.Hour*2)

	return Session{Token: token, FansId: fans.ID, SessionKey: sessionKey}
}

func (s *WxFansService) SetUserInfo(nickname string, gender int, province string, city string, avatarUrl string) error {
	db := models.DbLink()

	user, _ := s.Ctx.Get("user")
	fans := user.(models.WxFans)
	db.Debug().Model(&fans).Updates(models.WxFans{
		Nickname:  nickname,
		Gender:    uint(gender),
		Province:  province,
		City:      city,
		AvatarUrl: avatarUrl,
	})

	tokenStr := strconv.Itoa(int(fans.ID)) + "_" + fans.Openid
	data := []byte(tokenStr)
	tokenByte := sha1.Sum(data)

	token := hex.EncodeToString(tokenByte[:])
	fansJson, err := json.Marshal(fans)
	if err != nil {
		return err
	}

	var ctx = context.Background()
	redis := cache.GetClient()
	redis.Set(ctx, cache.PrefixKey(token), fansJson, time.Hour*2)

	return nil
}
