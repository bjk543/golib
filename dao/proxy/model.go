package proxy

import "time"

type Proxy struct {
	Addr       string    `gorm:"primary_key;type:varchar(24);"`
	Success    uint64    `gorm:"type:decimal(5);default:0"`
	SuccessPtt uint64    `gorm:"type:decimal(5);default:0"`
	Fail       uint64    `gorm:"type:decimal(5);default:0"`
	Active     bool      `gorm:"default:true"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:now()"`
	LastSuc    time.Time `gorm:"type:timestamp;"`
}
