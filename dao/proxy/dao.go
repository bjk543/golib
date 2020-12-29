package proxy

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
	CreateProxy(symbol string) (proxy *Proxy, err error)
	CreateProxies(addrs []string) (err error)
	GetProxy() (proxy []Proxy, err error)
	GetProxyAll() (proxy []Proxy, err error)
	SaveProxy(proxy []Proxy) ([]Proxy, error)
}

func CreateConn(user, pass, host, port, dbName string) Conn {

	db := dao.CreateConn(user, pass, host, port, dbName)
	if db == nil {
		return nil
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
	if err := d.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&us).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) GetProxyAll() (proxy []Proxy, err error) {
	var slice1 = []Proxy{}
	if err := d.db.Where("active = ? ", true).Find(&slice1).Error; err != nil {
		return nil, err
	}
	return slice1, nil
}

func (d *DAO) GetProxy() (proxy []Proxy, err error) {
	var slice1 = []Proxy{}
	if err := d.db.Order("fail asc,success_ptt desc,success asc").Where("active = ? and success_ptt > 0", true).Find(&slice1).Error; err != nil {
		return nil, err
	}
	return slice1, nil
}

func (d *DAO) SaveProxy(proxies []Proxy) (res []Proxy, err error) {
	for _, v := range proxies {
		if err := d.db.Save(&v).Error; err != nil {
			logger.Log("ERROR", fmt.Sprintf("%s", err))
		}
	}

	return proxies, nil
}
