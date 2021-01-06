package goods

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jjsd-go-api/api/mini/models"
)

type GoodService struct {
	Ctx *gin.Context
}

type KeywordList struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Format string `json:"format"`
}

func (s *GoodService) KeywordList(keyword string) ([]KeywordList, error, bool) {
	db := models.DbLink()
	var goods []models.Goods
	keyword = "%" + keyword + "%"
	err := db.Debug().Where("name LIKE ? OR short_name LIKE ?", keyword, keyword).Find(&goods).Error

	if err == gorm.ErrRecordNotFound {
		return []KeywordList{}, err, true
	}

	if err != nil {
		return []KeywordList{}, err, false
	}

	var list []KeywordList
	l := len(goods)
	for i := 0; i < l; i++ {
		list = append(list, KeywordList{
			ID:     goods[i].ID,
			Name:   goods[i].Name,
			Format: goods[i].Format,
		})
	}

	return list, err, false
}

type NumberShow struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Number         string  `json:"number"`
	Unit           string  `json:"unit"`
	Format         string  `json:"format"`
	RetailPrice    float64 `json:"retail_price"`
	WholesalePrice float64 `json:"wholesale_price"`
}

func (s *GoodService) Number(number string) (NumberShow, error, bool) {
	db := models.DbLink()
	var goods models.Goods
	err := db.Debug().Where("number = ? AND is_del = ?", number, 0).First(&goods).Error

	if err == gorm.ErrRecordNotFound {
		return NumberShow{}, err, true
	}

	if err != nil {
		return NumberShow{}, err, false
	}

	return NumberShow{
		ID:             goods.ID,
		Name:           goods.Name,
		Number:         goods.Number,
		Unit:           goods.Unit,
		Format:         goods.Format,
		RetailPrice:    goods.RetailPrice,
		WholesalePrice: goods.WholesalePrice,
	}, err, false
}
