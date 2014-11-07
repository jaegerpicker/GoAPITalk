package main

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func AutoMigrateModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Task{})
}

func LogWrite(message string, logLevel string) {
	if (logLevel == "DEBUG" || logLevel == "INFO") && (settings.LogLevel == "INFO" || settings.LogLevel == "DEBUG") {
		log.Printf("\n\n\t(%s)Message logged:\n\t===========\n\t%s\n\n", logLevel, message)
	} else if logLevel == "ERROR" {
		log.Printf("\n\n\n\t(%s)ERROR LOGGED: \n====================\n\t%s\n\n\n", logLevel, message)
	}
}
