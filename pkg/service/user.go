package service

import (
	"errors"
	"go-start/pkg/model"
	"go-start/pkg/repository"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	Repo      *repository.UserRepository
	Validator *validator.Validate
}

func NewUserService(repo *repository.UserRepository, validator *validator.Validate) *UserService {
	return &UserService{
		Repo:      repo,
		Validator: validator,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	err := s.Validator.Struct(user)
	if err != nil {
		return err
	}

	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(id uint, user *model.User) error {
	return s.Repo.Update(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}
	if user.Password != password {
		return nil, errors.New("неверный пароль")
	}
	return user, nil
}
