package db

import (
	"fmt"
	"gorm/entity"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

type config struct {
	sslmode  string `yaml:"SSLMODE"`
	host     string `yaml:"POSTGRES_HOST"`
	port     uint   `yaml:"ports"`
	user     string `yaml:"POSTGRES_USER"`
	password string `yaml:"POSTGRES_PASSWORD"`
	dbname   string `yaml:"POSTGRES_DB"`
}

func (c *config) getConf() *config {

	yamlFile, err := ioutil.ReadFile("docker-compose.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func Init() {
	var c config
	c.getConf()
	fmt.Println(c)

	//db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ShoppingOnline password=cuocdoi123 sslmode=disable")
	db, err = gorm.Open("postgres", c)

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
