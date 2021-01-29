package models

import (
	"gorm.io/gorm"
	"time"
)

type WxFansShopRelation struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	ShopId    uint   `gorm:"column:shop_id"`
	FansId    uint   `gorm:"column:fans_id"`
	IsDel     uint   `gorm:"column:is_del"`
	CreatedAt uint64 `gorm:"column:created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at"`
	CreatedBy uint   `gorm:"column:created_by"`
	UpdatedBy uint   `gorm:"column:updated_by"`
}

func (WxFansShopRelation) TableName() string {
	return "sd_wx_fans_shop_relation"
}

func (t *WxFansShopRelation) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *WxFansShopRelation) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
