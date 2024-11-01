package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB(dbConfig map[interface{}]interface{}) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		int(dbConfig["port"].(int)),
		dbConfig["dbname"],
		dbConfig["charset"],
		dbConfig["parseTime"],
		dbConfig["loc"],
	)

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql connection failed: %v", err)
	}
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Mysql connection successful")
}
