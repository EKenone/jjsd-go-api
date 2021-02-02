package models

import (
	"gorm.io/gorm"
	"time"
)

type GoodsCategory struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	ShopId    uint   `gorm:"column:shop_id"`
	Title     string `gorm:"column:title"`
	IsDel     uint   `gorm:"column:is_del"`
	CreatedAt uint64 `gorm:"column:created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at"`
	CreatedBy uint   `gorm:"column:created_by"`
	UpdatedBy uint   `gorm:"column:updated_by"`
}

func (GoodsCategory) TableName() string {
	return "sd_goods_category"
}

func (t *GoodsCategory) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *GoodsCategory) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
