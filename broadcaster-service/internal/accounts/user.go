package accounts

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Username    string
	Name        string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) ToDbMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           u.Id,
		"name":         u.Name,
		"username":     u.Username,
		"phone_number": u.PhoneNumber,
	}
}

func (u *User) ToLogMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           u.Id,
		"name":         u.Name,
		"username":     u.Username,
		"phone_number": u.PhoneNumber,
	}
}
