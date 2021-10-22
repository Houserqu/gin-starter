package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int
	Name        string
	Email       string
	Age         int
	Birthday    time.Time
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
