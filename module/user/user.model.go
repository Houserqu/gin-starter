package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int
	Name        string `form:"name"`
	Email       string `form:"email"`
	Age         int
	Birthday    time.Time
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
