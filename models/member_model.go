package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id            int       `json:"id" gorm:"type:int;primaryKey;size:32"`
	VerifiedAt    time.Time `json:"verifiedAt" gorm:"type:datetime;not null"`
	Name          string    `json:"name" gorm:"type:varchar(32);not null"`
	Gender        string    `json:"gender" gorm:"type:char(1);not null"`
	PlaceOfBirth  string    `json:"placeOfBirth" gorm:"type:varchar(32);not null"`
	DateOfBirth   time.Time `json:"dateOfBirth" gorm:"type:date;not null"`
	Address1      string    `json:"address1" gorm:"type:varchar(64);not null"`
	Address2      string    `json:"address2" gorm:"type:varchar(64)"`
	Profession    string    `json:"profession" gorm:"type:char(1);not null"`
	Institution   string    `json:"institution" gorm:"type:varchar(64)not null"`
	PhoneNo       string    `json:"phoneNo" gorm:"type:varchar(16);not null"`
	IsWhatsapp    bool      `json:"isWhatsapp" gorm:"not null"`
	IdentityNo    string    `json:"identityNo" gorm:"type:varchar(32);not null"`
	IdentityType  string    `json:"identityType" gorm:"type:char(1);not null"`
	IdentityFile  string    `json:"identityFile" gorm:"type:varchar(16);not null"`
	PhotoFile     string    `json:"photoFile" gorm:"not null;type:varchar(16)"`
	AgreementFile string    `json:"agreementFile" gorm:"type:varchar(16)"`
}
