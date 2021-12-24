package models

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Cnpj      string         `json:"cnpj" gorm:"size:14"`
	IsOpen    bool           `json:"is_open"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type RestaurantRepository interface {
	All() ([]Restaurant)
	Find(id int) (Restaurant)
	Delete(id int) (Restaurant)
	Update(id int) (Restaurant)
	Create(id int) (Restaurant)
}