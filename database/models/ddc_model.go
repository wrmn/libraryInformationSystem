package models

type Ddc struct {
	Ddc   string `json:"ddc" gorm:"type:char(3);primaryKey"`
	Group string `json:"group" gorm:"type:varchar(32);not null"`
	Name  string `json:"name" gorm:"type:varchar(32);not null"`
	Book  []Book `gorm:"foreignKey:DdcNo;association_foreign_key:Ddc"`
}
