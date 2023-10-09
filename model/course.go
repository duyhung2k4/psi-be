package model

import (
	"time"

	"gorm.io/gorm"
)

type COURSE_TYPE string

const (
	COURSE_TYPE_SUBJECT12            COURSE_TYPE = "subject12"
	COURSE_TYPE_MAJOR                COURSE_TYPE = "major"
	COURSE_TYPE_SKILL                COURSE_TYPE = "skill"
	COURSE_TYPE_LANGUAGE_CENTIFICATE COURSE_TYPE = "languageCentificate"
)

type Course struct {
	Id        uint        `json:"id" gorm:"primaryKey"`
	CreatorId uint        `json:"creatorId"`
	Name      string      `json:"name"`
	Title     string      `json:"title"`
	Code      string      `json:"code" gorm:"unique"`
	Type      COURSE_TYPE `json:"type"`
	Detail    string      `json:"detail"`
	Price     int         `json:"price"`

	Creator *Profile `json:"creator" gorm:"foreignKey:CreatorId; constraint:OnUpdate: CASCADE,OnDelete: SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
