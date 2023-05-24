package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/silabig1294/testbank/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

// use 
func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===========================================\n", sql)
}

// bring variable DB for connection DB by Gorm use capitalized consonant for public variable
var DB *gorm.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/bank?parseTime=true&loc=Asia%2FBangkok"

	dial := mysql.Open(dsn)

	connection, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to Connect to database")
		os.Exit(2)
		panic("couls not connect to DB")
	}
	DB = connection
	connection.AutoMigrate(&models.User{},&models.Account{},&models.Transaction{})
    // connection.Migrator().CreateTable(&models.Account{})
}
