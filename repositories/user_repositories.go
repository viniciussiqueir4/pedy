package repositories

import (
	"errors"
	"pedy/database"
	"pedy/models"
)

type UserRepository struct {
}

func (s *UserRepository) Add(p models.User) (models.User, error) {
	db := database.GetDatabase()

	err := db.Create(&p).Error

	if err != nil {
		return models.User{}, err
	}
	return p, nil
}

func (s *UserRepository) GetByEmail(email string) (models.User, error) {
	db := database.GetDatabase()

	var user models.User
	err := db.First(&user, "email = ?", email).Error

	if err != nil {
		return models.User{}, errors.New("Invalid credentials")
	}

	return user, nil
}

func (s *UserRepository) GetById(id uint) (models.User, error) {
	db := database.GetDatabase()

	var user models.User

	err := db.First(&user, id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserRepository) ExistEmail(email string) (bool, error) {
	db := database.GetDatabase()

	var user models.User
	result := db.First(&user, "email = ?", email)

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (s *UserRepository) ExistCpf(cpf string) (bool, error) {
	db := database.GetDatabase()

	var user models.User
	result := db.First(&user, "cpf = ?", cpf)

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (s *UserRepository) ExistCellphone(cellophone string) (bool, error) {
	db := database.GetDatabase()

	var user models.User
	result := db.First(&user, "cellphone = ?", cellophone)

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
