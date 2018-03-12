package db

import (
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/mysql"
	"os"
)

var (
	db *gorm.DB
)

func GetConnection() *gorm.DB {

	if db == nil {
		config := &mysql.Config{
			User:      "root",
			Addr:      "localhost",
			DBName:    "go_simple_chat",
			ParseTime: true,
		}

		var err error
		db, err = gorm.Open("mysql", config.FormatDSN())
		if err != nil {
			os.Exit(1)
		}
	}

	return db
}
