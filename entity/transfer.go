package entity

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	UserId         uint
	UserPenerimaId uint
	Nominal        uint
	User           User `gorm:"foreignKey:UserId;references:ID"`
	UserPenerima   User `gorm:"foreignKey:UserPenerimaId;references:ID"`
}

func (t Transfer) TableName() string {
	return "transfer"
}
