package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "3306"
	db_name := "blogspot"
	username := "root"
	password := ""

	url := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db_name + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic("Can't connect DB")
	}
	DB.AutoMigrate()
}
