package driver

import (
	"fmt"
	"gokomodo-api/configs"
	"gokomodo-api/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entities.Seller{})
	err = db.AutoMigrate(&entities.Buyer{})
	err = db.AutoMigrate(&entities.Product{})
	err = db.AutoMigrate(&entities.Order{})

	if err != nil {
		log.Info("error auto migrate", err)
	}

}
