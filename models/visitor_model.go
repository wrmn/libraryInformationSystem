package models

import "time"

type Visitor struct {
	Id      int       `json:"id" gorm:"type:int;size:32;primaryKey"`
	UserId  int       `json:"userId" gorm:"type:int;size:32;"`
	GuestId int       `json:"guestId" gorm:"type:int;size:32;"`
	LoginAt time.Time `json:"loginAt" gorm:"type:datetime"`
	Method  string    `json:"method" gorm:"type:char(1);not null"`
	Purpose string    `json:"purpose" gorm:"type:varchar(32);not null"`
	User    User      `gorm:"foreignkey:UserId"`
}
