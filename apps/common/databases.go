package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:123456789@tcp(localhost:3306)/gentayu"))
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&Book{})
	DB = db
}
