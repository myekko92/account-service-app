package main

import (
	"account-service-app/config"
	"account-service-app/controllers"
	"account-service-app/entity"
	"fmt"
)

func init() {
	db := config.MysqlConnect()
	db.AutoMigrate(&entity.User{}, &entity.Transfer{}, &entity.TopUp{})
}

func main() {
	pilihan := ""
	fmt.Println("===========================")
	fmt.Println("Daftar pilihan ")
	fmt.Println("===========================")
	fmt.Println("1 = Add User")
	fmt.Println("2 = List User")
	fmt.Println("3 = Update User")
	fmt.Println("4 = Delete User")
	fmt.Println("5 = Top-up")
	fmt.Println("6 = Transfer")
	fmt.Println("7 = History Top-up")
	fmt.Println("8 = History Transfer")
	fmt.Println("============================")
	fmt.Println("0 = Exit")
	fmt.Println("============================")
	fmt.Print("Masukkan pilihan menu anda: ")
	fmt.Scanln(&pilihan)
	fmt.Println("============================")

	menuController(pilihan)
}

func menuController(pilihan string) {
	switch pilihan {
	case "1":
		controllers.AddUser()
	case "2":
		controllers.ListUser()
	case "3":
		controllers.UpdateUser()
	case "4":
		controllers.DeleteUser()
	case "5":
		controllers.TopUp()
	case "6":
		controllers.Transfer()
	case "7":
		controllers.HistoryTopUp()
	case "8":
		controllers.HistoryTransfer()
	}

}
