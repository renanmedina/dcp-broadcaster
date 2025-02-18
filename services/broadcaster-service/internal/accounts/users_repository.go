package accounts

import (
	"github.com/renanmedina/dcp-broadcaster/utils"
	"gorm.io/gorm"
)

const USERS_TABLE = "users"

type UsersRepository struct {
	db *gorm.DB
}

func (r UsersRepository) GetById(userId string) (User, error) {
	var user User
	result := r.db.Table(USERS_TABLE).First(&user, userId)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r UsersRepository) GetAllSubscribed() []User {
	var users []User
	result := r.db.Table(USERS_TABLE).Where("subscribed = ?", 1).Find(&users)

	if result.Error != nil {
		return make([]User, 0)
	}

	return users
}

func NewUsersRepository() UsersRepository {
	return UsersRepository{
		utils.GetDatabaseConnection(),
	}
}
