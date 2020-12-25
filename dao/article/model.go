package article

import "time"

type Article struct {
	Addr      string    `gorm:"primary_key;type:varchar(128);"`
	Retry     uint64    `gorm:"type:decimal(2);default:0"`
	Done      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
}
