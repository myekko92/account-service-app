package controllers

import (
	"account-service-app/entity"
	"fmt"
)

func AddUser() {
	addUser := entity.User{}

	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&addUser.Nama)
	fmt.Println("Masukkan Nomor HP:")
	fmt.Scanln(&addUser.Phone)
	// fmt.Println("Masukkan Saldo:")
	// fmt.Scanln(&addUser.Saldo)

	tx := db.Save(&addUser)
	if tx.Error != nil {
		panic(tx.Error)
	}

	// fmt.Println("=====================")
	fmt.Println("Add User Successfully")
}

func ListUser() {
	var listUser []entity.User
	tx := db.Find(&listUser)
	if tx.Error != nil {
		panic(tx.Error)
	}
	for _, value := range listUser {
		fmt.Println("ID:", value.ID, "Nama:", value.Nama, "Nomor HP:", value.Phone, "Saldo:", value.Saldo)
	}
}

func UpdateUser() {
	var IDUser uint
	fmt.Print("Input ID User untuk Update:")
	fmt.Scanln(&IDUser)

	user := entity.User{}
	db.Find(&user, IDUser)
	fmt.Println(user)

	var nama, Phone string
	// var saldo uint

	fmt.Println("===============================")
	fmt.Print("1 = Nama [", user.Nama, "] : ")
	fmt.Scanln(&nama)
	if nama != "" {
		user.Nama = nama
	}
	fmt.Print("2 = Nomor HP [", user.Phone, "] : ")
	fmt.Scanln(&Phone)
	if Phone != "" {
		user.Phone = Phone
	}
	db.Save(user)
}

func DeleteUser() {
	var id uint
	fmt.Print("Masukkan ID User yang akan anda hapus:")
	fmt.Scanln(&id)

	user := entity.User{}
	db.Delete(&user, id)
}
