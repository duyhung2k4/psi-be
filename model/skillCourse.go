package model

import (
	"time"

	"gorm.io/gorm"
)

type SkillCourse struct {
	Id       uint `json:"id" gorm:"primaryKey"`
	CourseId uint `json:"courseId"`
	SkillId  uint `json:"skillId"`

	Course *Course `json:"course" gorm:"foreignKey:CourseId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Skill  *Skill  `json:"skill" gorm:"foreignKey:SkillId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
