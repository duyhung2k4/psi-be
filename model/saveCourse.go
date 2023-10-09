package model

import (
	"time"

	"gorm.io/gorm"
)

type SaveCourse struct {
	Id        uint `json:"id" gorm:"primaryKey"`
	CourseId  uint `json:"courseId"`
	ProfileId uint `json:"profileId"`

	Course  *Course  `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
