package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Common struct {
		Id             uint      `gorm:"primary_key"`
		CreatedAt      time.Time `gorm:"autoCreateTime"`
		UpdatedAt      time.Time `gorm:"autoUpdateTime"`
		gorm.DeletedAt `gorm:"index"`
	}
)
