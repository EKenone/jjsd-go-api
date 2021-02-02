package goods_category

import (
	"github.com/gin-gonic/gin"
	"jjsd-go-api/api/mini/models"
)

type GoodCategoryService struct {
	Ctx *gin.Context
}

type CategorySelect struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// 规格
func (s *GoodCategoryService) List() []CategorySelect {
	var category []CategorySelect
	db := models.DbLink()

	db.Model(&models.GoodsCategory{}).Where("shop_id = ? AND is_del = ?", s.Ctx.Query("shop_id"), 0).Order("id DESC").Find(&category)

	return category
}
