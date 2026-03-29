package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	DBUser     = "root"
	DBPassword = "12345678"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "sparkhire"
)

var DB *gorm.DB

func InitDBGorm() error {
	gormDB, err := newInit()
	if err != nil {
		return err
	}
	DB = gormDB
	return nil
}

func newInit() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
