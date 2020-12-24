package proxy

import (
	"fmt"
	"time"

	logger "github.com/bjk543/golib/log"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DAO struct {
	db *gorm.DB
}

type Conn interface {
	CreateProxy(symbol string) (proxy *Proxy, err error)
	CreateProxies(addrs []string) (err error)
	GetProxy() (proxy *[]Proxy, err error)
	SaveProxy(proxy []Proxy) (*[]Proxy, error)
}

func CreateConn(user, pass, host, port, dbName string) Conn {
	var db *gorm.DB
	var err error
	log.Printf("Connect database: %s:%s/%s %s %s\n", host, port, dbName, user, pass)
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

	if err := db.AutoMigrate(&Proxy{}); err != nil {
		log.WithFields(log.Fields{
			"message": err,
		}).Fatal("Can not migrate")
	}
	return &DAO{db: db}
}

func (d *DAO) CreateProxy(addr string) (proxy *Proxy, err error) {
	u := Proxy{
		Addr:   addr,
		Active: true,
	}
	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if err := d.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *DAO) CreateProxies(addrs []string) (err error) {
	us := make([]Proxy, len(addrs))
	for idx := range addrs {
		us[idx].Addr = addrs[idx]
		us[idx].Active = true
	}

	// d.db.CreateInBatches(users, 100)
	if err := d.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&us).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) GetProxy() (proxy *[]Proxy, err error) {
	var slice1 = []Proxy{}
	if err := d.db.Where("active = ?", true).Find(&slice1).Error; err != nil {
		return nil, err
	}
	return &slice1, nil
}

func (d *DAO) SaveProxy(proxies []Proxy) (res *[]Proxy, err error) {
	for _, v := range proxies {
		if err := d.db.Save(&v).Error; err != nil {
			logger.Log("ERROR", fmt.Sprintf("%s", err))
		}
	}

	return &proxies, nil
}
