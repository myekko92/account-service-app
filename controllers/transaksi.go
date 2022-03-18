package controllers

import (
	"account-service-app/config"
	"account-service-app/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.MysqlConnect()
}

func listUser() {
	users := []entity.User{}
	db.Find(&users)
	fmt.Println("Daftar User")
	fmt.Println("# \tNama     \tHandphone ")
	fmt.Println("_______________________")
	for _, user := range users {
		fmt.Print(user.ID, "\t")
		fmt.Print(user.Nama, "\t\t")
		fmt.Print(user.Phone, "\t")
		fmt.Println()
	}
	fmt.Println("________________________")
}

func findUserByPhone(Phone string) (entity.User, bool) {
	user := entity.User{}
	err := db.Where("Phone = ?", Phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

func findUserByPhoneWithTopup(Phone string) (entity.User, bool) {
	user := entity.User{}
	err := db.Preload("TopUp").Where("Phone = ?", Phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

func findUserByPhoneWithTransferKe(Phone string) (entity.User, bool) {
	user := entity.User{}
	err := db.Preload("TransferKe").Where("Phone = ?", Phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false
	}
	return user, true
}

func TopUp() {
	listUser()
	Phone := ""
	fmt.Print("Ketikkan nomor hp untuk top-up : ")
	fmt.Scanln(&Phone)

	if Phone != "x" {
		user, userExist := findUserByPhone(Phone)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			TopUp()
		}

		var nominalTopUp uint = 0
		fmt.Print("Masukkan jumlah Top-up: ")
		fmt.Scanln(&nominalTopUp)

		user.Saldo = user.Saldo + nominalTopUp
		db.Save(&user)

		topUp := entity.TopUp{
			UserId:  user.ID,
			Nominal: nominalTopUp,
		}
		db.Create(&topUp)

		fmt.Println("Berhasil Top-up!")
		fmt.Println("____________________________")
		fmt.Println("Nama:", user.Nama)
		fmt.Println("Jumlah top-up:", topUp.Nominal)
		fmt.Println("Saldo Sekarang:", user.Saldo)
		fmt.Println("Waktu: ", topUp.CreatedAt)
	}
}

func Transfer() {
	listUser()
	Pengirim := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&Pengirim)

	if Pengirim != "x" {
		userPengirim, userPengirimExist := findUserByPhone(Pengirim)
		if !userPengirimExist {
			fmt.Println("User tidak ditemukan")
			Transfer()
		}

		Penerima := ""
		fmt.Print("Masukan nomor hp penerima : ")
		fmt.Scanln(&Penerima)

		if Penerima != "x" {
			userPenerima, userPenerimaExist := findUserByPhone(Penerima)
			if !userPenerimaExist {
				fmt.Println("User tidak ditemukan")
				Transfer()
			}

			var nominalTransfer uint = 0
			for {
				fmt.Print("Masukkan jumlah Trasfer: ")
				fmt.Scanln(&nominalTransfer)
				if userPengirim.Saldo >= nominalTransfer {
					break
				}
				fmt.Println("!! : Saldo kurang!")
			}
			userPengirim.Saldo = userPengirim.Saldo - nominalTransfer
			userPenerima.Saldo = userPenerima.Saldo + nominalTransfer
			db.Save(&userPengirim)
			db.Save(&userPenerima)

			transfer := entity.Transfer{
				UserId:         userPengirim.ID,
				UserPenerimaId: userPenerima.ID,
				Nominal:        nominalTransfer,
			}
			db.Create(&transfer)

			fmt.Println("----------------------")
			fmt.Println("Transfer Berhasil!")
			fmt.Println("----------------------")
			fmt.Println("Nama Pengirim \t:", userPengirim.Nama)
			fmt.Println("Nama Penerima \t:", userPenerima.Nama)
			fmt.Println("Jumlah transfer \t:", transfer.Nominal)
			fmt.Println("Saldo Anda Sekarang\t:", userPengirim.Saldo)
			fmt.Println("Waktu: ", transfer.CreatedAt)
		}
	}
}

func HistoryTopUp() {
	listUser()
	Phone := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&Phone)

	if Phone != "x" {
		user, userExist := findUserByPhoneWithTopup(Phone)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			HistoryTopUp()
		}

		fmt.Println("------------------------------------------------------")
		fmt.Println("Data Transaksi Topup - ", user.Nama, "| Saldo sekarang:", user.Saldo)
		fmt.Println("------------------------------------------------------")
		for _, topUp := range user.TopUp {
			fmt.Println(topUp.Nominal, "\t", topUp.CreatedAt)
		}

		if len(user.TransferKe) <= 0 {
			fmt.Println("Tidak ada riwayat transaksi")
		}
	}
}

func HistoryTransfer() {
	listUser()

	Phone := ""
	fmt.Print("Ketikkan nomor hp anda: ")
	fmt.Scanln(&Phone)

	if Phone != "q" {
		user, userExist := findUserByPhoneWithTransferKe(Phone)
		if !userExist {
			fmt.Println("User tidak ditemukan")
			HistoryTransfer()
		}

		fmt.Println("------------------------------------------------------")
		fmt.Println("Data Transaksi Kirim Transfer - ", user.Nama, "| Saldo sekarang:", user.Saldo)
		fmt.Println("------------------------------------------------------")
		for _, topUp := range user.TransferKe {
			user := entity.User{}
			db.Find(&user, topUp.UserPenerimaId)
			fmt.Println("Transfer", topUp.Nominal, "\t", "ke", user.Nama, "pada waktu: ", topUp.CreatedAt)
		}

		if len(user.TransferKe) <= 0 {
			fmt.Println("Tidak ada riwayat transaksi")
		}
	}
}
