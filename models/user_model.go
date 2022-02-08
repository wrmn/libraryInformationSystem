package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           int           `json:"id" gorm:"type:int;primaryKey;size:32"`
	Username     string        `json:"username" gorm:"type:varchar(24);not null"`
	Email        string        `json:"email" gorm:"type:varchar(64);unique;not null"`
	Password     []byte        `json:"password" gorm:"not null"`
	LastLogin    time.Time     `json:"lastLogin" gorm:"type:datetime"`
	Employee     Employee      `gorm:"foreignkey:Id;references:Id"`
	Member       Member        `gorm:"foreignkey:Id;references:Id"`
	AssetRecord  []AssetRecord `gorm:"foreignKey:AdminId;association_foreign_key:Id"`
	Visitor      []Visitor     `gorm:"foreignKey:UserId;association_foreign_key:Id"`
	Admin        []Borrow      `gorm:"foreignKey:AdminId;association_foreign_key:Id"`
	BorrowMember []Borrow      `gorm:"foreignKey:MemberId;association_foreign_key:Id"`
}
