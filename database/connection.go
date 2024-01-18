package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/digihero"))
	if err != nil {
		panic(err)
	}

	return database

}
