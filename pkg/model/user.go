package model

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=6"`
}
