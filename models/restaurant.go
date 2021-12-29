package models

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null" validate:"required,max=255"`
	Cnpj      string         `json:"cnpj" gorm:"not null;size:14;unique" validate:"required,min=14,max=14"`
	IsOpen    bool           `json:"is_open"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func (r Restaurant) Validate() []error {
	err, trans := SetValidationPtBr()
	if err != nil {
		newErr := []error{err}
		return newErr
	}
	err = Validate.Struct(r)
	errs := TranslateError(err, trans)
	return errs
}
