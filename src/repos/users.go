package repos

import (
	"gorm.io/gorm"

	"platform-exer/src/models"
)

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{db}
}

type UsersRepo interface {
	Get(email string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

type usersRepo struct {
	db *gorm.DB
}

// Retrieve the user model
func (u *usersRepo) Get(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Create a new user model
func (u *usersRepo) Create(user *models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Update the specified user model
func (u *usersRepo) Update(user *models.User) error {
	return u.db.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"FirstName": user.FirstName,
			"LastName":  user.LastName,
		}).Error
}

// Delete the specified user model
func (u *usersRepo) Delete(user *models.User) error {
	return u.db.Where("name = ?", "jinzhu").Delete(&user).Error
}
