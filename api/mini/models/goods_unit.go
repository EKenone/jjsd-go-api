package models

import (
	"gorm.io/gorm"
	"time"
)

type GoodsUnit struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	Title     string `gorm:"column:title"`
	IsDel     uint   `gorm:"column:is_del"`
	CreatedAt uint64 `gorm:"column:created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at"`
	CreatedBy uint   `gorm:"column:created_by"`
	UpdatedBy uint   `gorm:"column:updated_by"`
}

func (GoodsUnit) TableName() string {
	return "sd_goods_unit"
}

func (t *GoodsUnit) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *GoodsUnit) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
