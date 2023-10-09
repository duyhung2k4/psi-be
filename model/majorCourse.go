package model

import (
	"time"

	"gorm.io/gorm"
)

type MajorCourse struct {
	Id       uint `json:"id" gorm:"primaryKey"`
	MajorId  uint `json:"majorId"`
	CourseId uint `json:"courseId"`

	Course *Course `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Major  *Major  `json:"major" gorm:"foreignKey:MajorId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
