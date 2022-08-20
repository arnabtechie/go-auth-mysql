package sql_connector

import (
	"fmt"

	"github.com/arnabtechie/go-ecommerce/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() {
	connection_string := "root:arnabroot@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection_string,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	} else {
		fmt.Println("database connected...")
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
