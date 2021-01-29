package models

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	ID             uint    `gorm:"column:id;primaryKey"`
	ShopId         uint    `gorm:"column:shop_id"`
	Name           string  `gorm:"column:name"`
	ShortName      string  `gorm:"column:short_name"`
	Number         string  `gorm:"column:number"`
	Unit           string  `gorm:"column:unit"`
	Format         string  `gorm:"column:format"`
	Stock          float64 `gorm:"column:stock"`
	PurchasePrice  float64 `gorm:"column:purchase_price"`
	RetailPrice    float64 `gorm:"column:retail_price"`
	WholesalePrice float64 `gorm:"column:wholesale_price"`
	ImgSource      string  `gorm:"column:img_source"`
	ProductDate    string  `gorm:"column:product_date"`
	ShelfLife      string  `gorm:"column:shelf_life"`
	IsDel          uint    `gorm:"column:is_del"`
	CreatedAt      uint64  `gorm:"column:created_at"`
	UpdatedAt      uint64  `gorm:"column:updated_at"`
	CreatedBy      uint    `gorm:"column:created_by"`
	UpdatedBy      uint    `gorm:"column:updated_by"`
}

func (Goods) TableName() string {
	return "sd_goods"
}

func (t *Goods) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *Goods) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
