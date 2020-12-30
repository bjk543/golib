package dao

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Close(db *gorm.DB) {
	sql, err := db.DB()
	if err == nil && sql != nil {
		sql.Close()
	}
}

func CreateConn(user, pass, host, port, dbName string) *gorm.DB {
	var db *gorm.DB
	var err error
	for i := 0; i < 5; i++ {
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Taipei", host, port, user, dbName, pass)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		// db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, pass))
		if err != nil {
			log.WithFields(log.Fields{
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

	return db
}
