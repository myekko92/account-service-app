package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string
	Phone        string
	Saldo        uint
	TopUp        []TopUp    `gorm:"foreignKey:UserId;references:ID"`
	TransferKe   []Transfer `gorm:"foreignKey:UserId;references:ID"`
	TransferDari []Transfer `gorm:"foreignKey:UserPenerimaId;references:ID"`
}

func (u User) TableName() string {
	return "user"
}
