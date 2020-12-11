package models

import (
	"gorm.io/gorm"
	"time"
)

type WxFans struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	Openid    string `gorm:"column:openid"`
	Unionid   string `gorm:"column:unionid"`
	Nickname  string `gorm:"column:nickname"`
	Gender    uint   `gorm:"column:gender"`
	Province  string `gorm:"column:province"`
	City      string `gorm:"column:city"`
	AvatarUrl string `gorm:"column:avatar_url"`
	IsDel     uint   `gorm:"column:is_del"`
	CreatedAt uint64 `gorm:"column:created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at"`
	CreatedBy uint   `gorm:"column:created_by"`
	UpdatedBy uint   `gorm:"column:updated_by"`
}

func (WxFans) TableName() string {
	return "sd_wx_fans"
}

func (t *WxFans) BeforeCreate(db *gorm.DB) error {
	t.CreatedAt = uint64(time.Now().Unix())
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}

func (t *WxFans) BeforeUpdate(db *gorm.DB) error {
	t.UpdatedAt = uint64(time.Now().Unix())

	return nil
}
