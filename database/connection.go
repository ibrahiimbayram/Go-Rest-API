package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	var err error
	const connString = "server=DESKTOP-F621BN8;user id=;password=;port=1433;database=GoProject"

	DB, err = gorm.Open(sqlserver.Open(connString))

	if err != nil {
		log.Fatal(err)
	}

}
