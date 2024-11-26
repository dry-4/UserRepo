package databses

import (
	"fmt"
	"sync"
	"userRepo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

// DBConfig holds db configrations
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Database connection instance
var dbConfig = DBConfig{
	Host:     "127.0.0.1",
	Port:     "3306",
	User:     "root",
	Password: "Sudo@123",
	DBName:   "store",
}

func initDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	once.Do(func() {
		db, err = gorm.Open(mysql.Open(connectionString))
		if err != nil {
			err = fmt.Errorf("failed to connect to Mysql database : %v", err)
		}

		err = db.AutoMigrate(&models.User{})

		if err != nil {
			err = fmt.Errorf("failed to migrate databse: %v", err)
		}
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	initDB()
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}
