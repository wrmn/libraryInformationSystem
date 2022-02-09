package models

type Guest struct {
	Id          int     `json:"id" gorm:"type:int;primaryKey;size:32"`
	Name        string  `json:"name" gorm:"type:varchar(32);not null"`
	Gender      string  `json:"gender" gorm:"type:char(1);not null"`
	Address     string  `json:"address" gorm:"type:varchar(96);not null"`
	Profession  int     `json:"profession" gorm:"type:int;size:32;not null"`
	Institution string  `json:"institution" gorm:"type:varchar(64); not null"`
	Visitor     Visitor `gorm:"foreignkey:GuestId;references:Id"`
}
