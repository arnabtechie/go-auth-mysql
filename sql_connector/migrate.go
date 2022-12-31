package sql_connector

import (
	"log"
	"os"

	"github.com/arnabtechie/go-ecommerce/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() {
	connection_string := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection_string,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
		panic("cannot connect to database...")
	} else {
		log.Println("database connected...")
	}
	db.AutoMigrate(&models.User{}, &models.Product{})
	log.Println("migration completed...")
}
