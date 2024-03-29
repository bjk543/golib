package dao

import (
	"fmt"
	"log"
	"os"
	"time"

	log1 "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	LOG_LEVEL_DB = ""
	LogLevel     logger.LogLevel
)

func init() {
	LOG_LEVEL_DB = os.Getenv("LOG_LEVEL_DB")

}

func Close(db *gorm.DB) {
	sql, err := db.DB()
	if err == nil && sql != nil {
		sql.Close()
	}
}

func CreateConn(user, pass, host, port, dbName string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch LOG_LEVEL_DB {
	case "SILENT":
		LogLevel = logger.Silent
	case "INFO":
		LogLevel = logger.Info
	case "WARN":
		LogLevel = logger.Warn
	case "ERROR":
		LogLevel = logger.Error
	default:
		LogLevel = logger.Silent
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  LogLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	for i := 0; i < 5; i++ {
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Taipei", host, port, user, dbName, pass)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
		// db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, pass))
		if err != nil {
			log1.WithFields(log1.Fields{
				"host":    host,
				"port":    port,
				"user":    user,
				"name":    dbName,
				"message": err.Error(),
			}).Println("Can not connect to database")
			time.Sleep(time.Duration(i) * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return nil
	}

	ddb, err := db.DB()
	if err != nil {
		log1.WithFields(log1.Fields{
			"host":    host,
			"port":    port,
			"user":    user,
			"name":    dbName,
			"message": err.Error(),
		}).Println("Can not DB returns *sql.DB")
	}

	// https://github.com/go-gorm/gorm/issues/1822
	if 1 == 2 { // don`t use ,slow query ?
		ddb.SetConnMaxLifetime(60 * time.Second) //这个时间和lb的idle超时短就行了
		ddb.SetMaxIdleConns(0)                   //不要使用连接池
	}

	return db
}
