package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() *gorm.DB {

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Tidak ada file konfigurasi dotenv tersedia")
	// }

	// dsn := os.Getenv("DSN")
	dsn := "root:Ekko123@tcp(127.0.0.1:3306)/wallet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak dapat terhubung dengan database")
	}
	return db
}
