package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDB() *gorm.DB {
	DB := os.Getenv("DB")
	USER := os.Getenv("user")
	PASS := os.Getenv("pass")
	HOST := os.Getenv("host")
	PORT := os.Getenv("portDB")
	DBNAME := os.Getenv("database")
	if DB == "mysql" {
		URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
		db, err := gorm.Open(DB, URL)
		if err != nil {
			panic(err.Error())
		}
		return db
	}
	if DB == "postgres" {
		//URL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", USER, PASS, HOST, DBNAME)
		//db, err := gorm.Open(DB, URL)
		db, err := gorm.Open("postgres", "host="+HOST+" port="+PORT+" user="+USER+" dbname="+DBNAME+" password="+PASS+" sslmode=disable")
		if err != nil {
			panic(err.Error())
		}
		return db
	}
	return nil
}
