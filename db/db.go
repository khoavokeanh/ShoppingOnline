package db

import (
	"fmt"
	"gorm/entity"

	"gorm/env"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	var c env.Config
	d := c.GetConf()

	//db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ShoppingOnline password=cuocdoi123 sslmode=disable")
	db, err = gorm.Open("postgres", d)

	if err != nil {
		fmt.Println("cnn fail")
		panic(err)
	} else {
		fmt.Println("Cnn success")
	}

	autoMigration()
}
func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
