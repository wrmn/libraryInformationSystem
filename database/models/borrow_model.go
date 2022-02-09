package models

import "time"

type Borrow struct {
	Id            int        `json:"id" gorm:"type:int;primaryKey;size:32"`
	AdminId       int        `json:"adminId" gorm:"type:int;size:32;not null"`
	MemberId      int        `json:"memberId" gorm:"type:int;size:32;not null"`
	BookId        int        `json:"bookId" gorm:"type:int;size:32;not null"`
	FineStatus    bool       `json:"fineStatus" gorm:"not null"`
	DateOfBorrow  time.Time  `json:"dateOfBorrow" gorm:"type:datetime;not null"`
	DateOfReturn  *time.Time `json:"dateOfReturn" gorm:"type:datetime"`
	DateOfPayment *time.Time `json:"dateOfPayment" gorm:"type:datetime"`
	Admin         User       `gorm:"foreignkey:AdminId"`
	Member        User       `gorm:"foreignkey:MemberId"`
	Book          Book       `gorm:"foreignkey:BookId"`
}
