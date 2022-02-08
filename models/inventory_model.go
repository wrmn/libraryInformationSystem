package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Id             int         `json:"id" gorm:"type:int;size:32;primaryKey"`
	RegistrationId int         `json:"registrationId" gorm:"type:int;size:32;not null"`
	SerialNumber   int         `json:"serialNumber" gorm:"type:int;size:32;not null"`
	Name           string      `json:"name" gorm:"type:varchar(32);not null"`
	Category       int         `json:"category" gorm:"type:int;size:32;not null"`
	Status         int         `json:"status" gorm:"type:int;size:32;not null"`
	Description    string      `json:"description" gorm:"type:int;size:256;not null"`
	AssetRecord    AssetRecord `gorm:"foreignkey:RegistrationId"`
}
