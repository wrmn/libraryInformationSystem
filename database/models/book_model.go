package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id             int         `json:"id" gorm:"type:int;primaryKey;size:32"`
	RegistrationId int         `json:"registrationId" gorm:"type:int;size:32;not null"`
	SerialNumber   int         `json:"serialNumber" gorm:"type:int;size:32;not null"`
	DdcNo          string      `json:"ddc" gorm:"type:char(3);not null"`
	DdcOrder       int         `json:"ddcOrder" gorm:"type:int;size:32;not null"`
	Title          string      `json:"title" gorm:"type:varchar(96);not null"`
	Author         string      `json:"author" gorm:"type:varchar(32);not null"`
	Publisher      string      `json:"publisher" gorm:"type:varchar(32);not null"`
	Availability   bool        `json:"availability" gorm:"not null"`
	Status         int         `json:"status" gorm:"type:int;size:32;not null"`
	Price          int         `json:"price" gorm:"type:int;size:32;not null"`
	CoverFile      string      `json:"cover" gorm:"type:varchar(24);not null"`
	AssetRecord    AssetRecord `gorm:"foreignkey:RegistrationId"`
	Ddc            Ddc         `gorm:"foreignkey:DdcNo"`
	Borrow         []Borrow    `gorm:"foreignKey:BookId;association_foreign_key:Id"`
}
