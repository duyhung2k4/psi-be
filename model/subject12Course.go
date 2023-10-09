package model

import (
	"time"

	"gorm.io/gorm"
)

type Subject12Course struct {
	Id             uint       `json:"id" gorm:"primaryKey"`
	Subject12Code  string     `json:"subject12Code"`
	CourseId       uint       `json:"courseId"`
	TargetPoint    string     `json:"targetPoint"`
	Time           *time.Time `json:"time"`
	Schedule       *string    `json:"schedule"`
	CourseDuration string     `json:"courseDuration"`
	Note           string     `json:"note"`

	Course    *Course    `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subject12 *Subject12 `json:"subject12" gorm:"foreignKey:Subject12Code; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
