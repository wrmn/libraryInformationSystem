package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           int           `json:"id" gorm:"type:int;primaryKey;size:32"`
	Username     string        `json:"username" gorm:"type:varchar(24);unique;not null"`
	Email        string        `json:"email" gorm:"type:varchar(64);unique;not null"`
	Password     string        `json:"password" gorm:"type:char(32);not null"`
	LastLogin    *time.Time    `json:"lastLogin" gorm:"type:timestamp"`
	Employee     Employee      `gorm:"foreignkey:Id;references:Id"`
	Member       Member        `json:"member" gorm:"foreignkey:Id;references:Id"`
	AssetRecord  []AssetRecord `gorm:"foreignKey:AdminId;association_foreign_key:Id"`
	Visitor      []Visitor     `gorm:"foreignKey:UserId;association_foreign_key:Id"`
	Admin        []Borrow      `gorm:"foreignKey:AdminId;association_foreign_key:Id"`
	BorrowMember []Borrow      `gorm:"foreignKey:MemberId;association_foreign_key:Id"`
}
