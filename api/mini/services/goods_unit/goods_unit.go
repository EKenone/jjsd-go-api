package goods_unit

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/models"
)

type GoodUnitService struct {
	Ctx *gin.Context
}

// 单位
func (s *GoodUnitService) List() []string {
	var unit []string
	db := models.DbLink()

	db.Model(&models.GoodsUnit{}).Where("is_del = ?", 0).Order("id DESC").Pluck("title", &unit)

	return unit
}
