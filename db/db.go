package db

import (
	"fmt"
	"log"
	"shobak/models"
	"shobak/pkg/setting"

	"github.com/kr/pretty"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error

	dbConfig := setting.Config.DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("db.Setup err:%w", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)

	//AutoMigrate
	autoMigrate()
	log.Println("DB successfully connected! ")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, err := db.DB()
	sqlDB.Close()
	if err != nil {
		pretty.Logln("Error on closing the DB: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func autoMigrate() {
	for _, model := range []interface{}{
		(*models.User)(nil),
	} {
		dbSilent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
		err := dbSilent.AutoMigrate(model)
		if err != nil {
			log.Fatalf("create model %s : %s", model, err)
		}
	}
}
