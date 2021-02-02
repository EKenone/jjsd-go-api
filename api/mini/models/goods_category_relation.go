package models

import (
	"gorm.io/gorm"
	"time"
)

type GoodsCategoryRelation struct {
	ID         uint   `gorm:"column:id;primaryKey"`
	GoodsId    uint   `gorm:"column:goods_id"`
	CategoryId uint   `gorm:"column:category_id"`
	IsDel      uint   `gorm:"column:is_del"`
	CreatedAt  uint64 `gorm:"column:created_at"`
	UpdatedAt  uint64 `gorm:"column:updated_at"`
	CreatedBy  uint   `gorm:"column:created_by"`
	UpdatedBy  uint   `gorm:"column:updated_by"`
}

func (GoodsCategoryRelation) TableName() string {
	return "sd_goods_category_relation"
}

func (t *GoodsCategoryRelation) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *GoodsCategoryRelation) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
