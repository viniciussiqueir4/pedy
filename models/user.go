package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null" validate:"required,min=2"`
	Email     string         `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password  string         `json:"password" gorm:"not null" validate:"required,min=6"`
	Cellphone string         `json:"cellphone" gorm:"not null" validate:"required,min=11"`
	Cpf       string         `json:"cpf" gorm:"not null;unique" validate:"required,min=11"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(password)

	err = user.validate()

	if err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

func NewUser(name string, email string, password string, cellphone string, cpf string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Cellphone: cellphone,
		Cpf:       cpf,
	}
}
