package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id              uint   `json:"id" gorm:"primaryKey"`
	UserId          uint   `json:"userId" gorm:"unique"`
	ProfileCourseId *uint  `json:"profileCourseId"`
	Phone           string `json:"phone"`

	Credential    *Credential    `json:"credential" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProfileCourse *ProfileCourse `json:"profileCourse" gorm:"foreignKey:ProfileCourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
