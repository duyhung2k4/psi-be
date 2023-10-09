package model

import (
	"time"

	"gorm.io/gorm"
)

type LanguageCertificateCourse struct {
	Id                    uint   `json:"id" gorm:"primaryKey"`
	CourseId              uint   `json:"courseId"`
	LanguageCertificateId uint   `json:"languageCertificateId"`
	TargetPoint           string `json:"targetPoint"`

	Course              *Course              `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LanguageCertificate *LanguageCertificate `json:"languageCertificate" gorm:"foreignKey:LanguageCertificateId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
