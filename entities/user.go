package entities

import (
	"time"
)

type User struct {
	ID        uint
	FullName  string
	Username  string
	Password  string
	Salt      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
