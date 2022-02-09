package models

import (
	"time"

	"gorm.io/gorm"
)

type AssetRecord struct {
	gorm.Model
	Id                 int         `json:"id" gorm:"type:int;size:32;primaryKey"`
	AdminId            int         `json:"adminId" gorm:"type:int;size:32;not null"`
	RegistrationNumber string      `json:"registrationNumber" gorm:"type:varchar(16);not null"`
	RegistrationDate   time.Time   `json:"registrationDate" gorm:"type:date;not null"`
	Source             string      `json:"source" gorm:"type:varchar(32);not null"`
	User               User        `gorm:"foreignkey:AdminId"`
	Book               []Book      `gorm:"foreignKey:RegistrationId;association_foreign_key:Id"`
	Inventory          []Inventory `gorm:"foreignKey:RegistrationId;association_foreign_key:Id"`
}
