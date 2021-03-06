package goods

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jjsd-go-api/api/mini/models"
	"jjsd-go-api/api/mini/services/goods_category"
	"jjsd-go-api/api/mini/services/goods_format"
	"jjsd-go-api/api/mini/services/goods_unit"
	"strconv"
	"strings"
	"sync"
)

type GoodService struct {
	Ctx *gin.Context
}

type KeywordList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (s *GoodService) KeywordList(keyword string) ([]KeywordList, error, bool) {
	db := models.DbLink()
	var goods []models.Goods
	keyword = "%" + keyword + "%"
	err := db.Debug().Where("shop_id = ?", s.Ctx.Query("shop_id")).Where("name LIKE ? OR short_name LIKE ?", keyword, keyword).Limit(50).Order("id desc").Find(&goods).Error

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
			ID:   goods[i].ID,
			Name: goods[i].Name,
			Unit: goods[i].Unit,
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

func (s *GoodService) GoodsNumber(number string) (NumberShow, error, bool) {
	db := models.DbLink()
	var goods models.Goods
	err := db.Debug().Where("shop_id = ? AND number = ? AND is_del = ?", s.Ctx.Query("shop_id"), number, 0).First(&goods).Error

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

func (s *GoodService) UpdateNumber(id int, number string) error {
	db := models.DbLink()

	err := db.Debug().Model(&models.Goods{}).Where("id = ?", id).Update("number", number).Error

	return err
}

type AddForm struct {
	ShopId         uint    `form:"shop_id"`
	Name           string  `form:"name"`
	ShortName      string  `form:"short_name"`
	Number         string  `form:"number"`
	Unit           string  `form:"unit"`
	Format         string  `form:"format"`
	WholesalePrice float64 `form:"wholesale_price"`
	RetailPrice    float64 `form:"retail_price"`
	Category       string  `form:"category"`
}

// 添加商品
func (s *GoodService) GoodsAdd(form AddForm) {

	goods := models.Goods{
		ShopId:         form.ShopId,
		Name:           form.Name,
		ShortName:      form.ShortName,
		Number:         form.Number,
		Unit:           form.Unit,
		Format:         form.Format,
		WholesalePrice: form.WholesalePrice,
		RetailPrice:    form.RetailPrice,
	}

	db := models.DbLink()
	db.Select("ShopId", "Name", "ShortName", "Number", "Unit", "Format", "WholesalePrice", "RetailPrice", "CreatedAt", "UpdatedAt").Create(&goods)

	category := strings.Split(form.Category, ",")

	for _, v := range category {
		cid, _ := strconv.Atoi(v)
		db.Create(&models.GoodsCategoryRelation{
			GoodsId:    goods.ID,
			CategoryId: uint(cid),
		})
	}

}

type AttrResp struct {
	Format   []string                        `json:"format"`
	Unit     []string                        `json:"unit"`
	Category []goods_category.CategorySelect `json:"category"`
}

func (s *GoodService) GoodsAttr() AttrResp {
	var attr AttrResp

	g := sync.WaitGroup{}
	g.Add(3)

	go func() {
		defer g.Done()
		formatService := goods_format.GoodFormatService{Ctx: s.Ctx}
		attr.Format = formatService.List()
	}()

	go func() {
		defer g.Done()
		unitService := goods_unit.GoodUnitService{Ctx: s.Ctx}
		attr.Unit = unitService.List()
	}()

	go func() {
		defer g.Done()
		categoryService := goods_category.GoodCategoryService{Ctx: s.Ctx}
		attr.Category = categoryService.List()
	}()

	g.Wait()

	return attr

}
