package models

import "time"

type Employee struct {
	Id             int       `json:"id" gorm:"type:int;size:32;primaryKey"`
	EmployeeNumber string    `json:"employeeNumber" gorm:"type:char(18);not null"`
	Name           string    `json:"name" gorm:"type:varchar(32);not null"`
	Gender         string    `json:"gender" gorm:"type:char(1);not null"`
	PlaceOfBirth   string    `json:"placeOfBirth" gorm:"type:varchar(32);not null"`
	DateOfBirth    time.Time `json:"dateOfBirth" gorm:"type:date;not null"`
	Address1       string    `json:"address1" gorm:"type:varchar(64);not null"`
	Address2       string    `json:"address2" gorm:"type:varchar(64)"`
	Division       string    `json:"division" gorm:"type:char(1);not null"`
	Position       string    `json:"position" gorm:"type:varchar(16);not null"`
}
