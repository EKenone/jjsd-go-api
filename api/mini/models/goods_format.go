package models

import (
	"gorm.io/gorm"
	"time"
)

type GoodsFormat struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	Title     string `gorm:"column:title"`
	IsDel     uint   `gorm:"column:is_del"`
	CreatedAt uint64 `gorm:"column:created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at"`
	CreatedBy uint   `gorm:"column:created_by"`
	UpdatedBy uint   `gorm:"column:updated_by"`
}

func (GoodsFormat) TableName() string {
	return "sd_goods_format"
}

func (t *GoodsFormat) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *GoodsFormat) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
