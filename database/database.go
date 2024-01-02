package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	errConfig := config.ReadInConfig()
	if errConfig != nil {
		panic("Invalid config!")
	}

	dsn := "root:@tcp(127.0.0.1:" + config.GetString("DATABASE_PORT") + ")/" + config.GetString("DATABASE_HOST") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to datase.")
}
