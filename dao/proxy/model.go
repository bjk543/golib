package proxy

import "time"

type Proxy struct {
	Addr      string    `gorm:"primary_key;type:varchar(10);"`
	Success   uint64    `gorm:"type:decimal(4);default:0"`
	Fail      uint64    `gorm:"type:decimal(4);default:0"`
	Active    bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
}
