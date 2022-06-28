package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("cannot connect to the DB")
	}

	log.Println("successfully connected to the data")
}

func Migrate() {
	Instance.AutoMigrate(&Product{})
	log.Println("data migration for users completed")
}
