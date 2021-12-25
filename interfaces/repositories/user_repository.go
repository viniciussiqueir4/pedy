package repositories

import "pedy/models"

type IUserRepository interface {
	Add(p models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetById(id uint) (models.User, error)
	ExistEmail(email string) (bool, error)
	ExistCpf(cpf string) (bool, error)
	ExistCellphone(cellphone string) (bool, error)
}
