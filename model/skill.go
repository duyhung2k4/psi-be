package model

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
