package article

import (
	"fmt"

	dao "github.com/bjk543/golib/dao"
	logger "github.com/bjk543/golib/log"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DAO struct {
	db *gorm.DB
}

type Conn interface {
	Create(addrs []string) (err error)
	Get() ([]Article, error)
	Save([]Article)
}

func CreateConn(user, pass, host, port, dbName string) Conn {

	db := dao.CreateConn(user, pass, host, port, dbName)
	log.Printf("Connect database: %s:%s/%s %s %s\n", host, port, dbName, user, pass)

	if db == nil {
		return nil
	}

	if err := db.AutoMigrate(&Article{}); err != nil {
		log.WithFields(log.Fields{
			"message": err,
		}).Fatal("Can not migrate")
	}
	return &DAO{db: db}
}

func (d *DAO) Create(addrs []string) (err error) {
	us := make([]Article, len(addrs))
	for idx := range addrs {
		us[idx].Addr = addrs[idx]
		us[idx].Done = false
	}
	if err := d.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&us).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) Get() ([]Article, error) {
	var res = []Article{}
	if err := d.db.Where("done = ? and retry<5", false).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DAO) Save(array []Article) {
	for _, v := range array {
		if err := d.db.Save(&v).Error; err != nil {
			logger.Log("ERROR", fmt.Sprintf("%s", err))
		}
	}
}
