package example

import "time"

type Example struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeleteAt  time.Time
}

// TableName 自定义表明
func (Example) TableName() string {
	return "example"
}
