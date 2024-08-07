package data

import (
	"errors"
	"gorm.io/gorm"
	"svc/proxy-service/internal/common"
	"time"
)

type SiteConfig struct {
	Id        int       `gorm:"primaryKey;autoIncrement;column:id"`
	SiteID    int       `gorm:"column:site_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdateAt  time.Time `gorm:"column:created_at"`
}

func (SiteConfig) TableName() string {
	return "site_configuration"
}

func FindSiteConfig(condition interface{}) SiteConfig {
	db := DB
	var model SiteConfig
	err := db.Where(condition).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Site config not found"})
	}
	return model
}

func FindSiteConfigs(condition interface{}) []SiteConfig {
	db := DB
	var model []SiteConfig
	err := db.Where(condition).Find(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Site config not found"})
	}
	return model
}
