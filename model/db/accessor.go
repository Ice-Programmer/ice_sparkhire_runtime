package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
