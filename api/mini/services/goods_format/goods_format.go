package goods_format

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/models"
)

type GoodFormatService struct {
	Ctx *gin.Context
}

// 规格
func (s *GoodFormatService) List() []string {
	var format []string
	db := models.DbLink()

	db.Model(&models.GoodsFormat{}).Where("shop_id = ? AND is_del = ?", s.Ctx.Query("shop_id"), 0).Order("id DESC").Pluck("title", &format)

	return format
}
