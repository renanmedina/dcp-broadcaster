package accounts

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey"`
	Username    string
	Name        string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// gorm before create hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()
	return nil
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
