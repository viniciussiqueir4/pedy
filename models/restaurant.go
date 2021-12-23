package models

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:14"`
	Cnpj      string         `json:"cnpj"`
	IsOpen    bool           `json:"is_open"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type RestaurantRepository interface {
	All() ([]Restaurant)
}