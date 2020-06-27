package dao

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"fmt"
)

var (
	db  *gorm.DB
	err error
)

type Model int

func DB() *gorm.DB {
	return db
}

func Init(uri string, debug bool) {
	db, err = gorm.Open("mysql", uri)

	if err != nil {
		log.Fatalf("Connect mysql [%v] err: %v", uri, err)
	}
	log.Printf("Connect mysql [%v] succ \n", uri)

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(debug)

	db.AutoMigrate(&User{})
}
