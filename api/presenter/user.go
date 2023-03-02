package presenter

import "time"

type User struct {
	FullName  string    `json:"full_name"`
	Username  string    `json:"username" validate:"required,min=8,max=16"`
	Password  string    `json:"password" validate:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
