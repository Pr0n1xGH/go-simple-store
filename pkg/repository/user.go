package repository

import (
	"go-start/pkg/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id uint, newUser *model.User) error {
	var user model.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return err
	}

	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password

	return r.DB.Save(&user).Error
}

func (r *UserRepository) Delete(id uint) error {
	var user model.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
