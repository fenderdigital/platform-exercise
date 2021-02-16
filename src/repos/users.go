package repos

import (
	"gorm.io/gorm"

	"platform-exer/src/models"
)

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{db}
}

type UsersRepo interface {
	Get() (*models.User, error)
	Update(user *models.User) error
}

type usersRepo struct {
	db *gorm.DB
}

func (u *usersRepo) Get() (*models.User, error) {
	return nil, nil
}

func (u *usersRepo) Update(user *models.User) error {
	return u.db.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"FirstName": user.FirstName,
			"LastName":  user.LastName,
		}).Error
	return nil
}