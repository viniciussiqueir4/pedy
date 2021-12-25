package user

import (
	"errors"
	"pedy/interfaces/repositories"
	"pedy/models"
)

type UserService struct {
	repository repositories.IUserRepository
}

func (s *UserService) CreateUser(data UserDto) (models.User, error) {

	exist, erro := s.repository.ExistEmail(data.Email)

	if erro != nil {
		return models.User{}, erro
	}

	if exist {
		return models.User{}, errors.New("Email already exist")
	}

	existCpf, errCpf := s.repository.ExistCpf(data.Cpf)

	if errCpf != nil {
		return models.User{}, errCpf
	}

	if existCpf {
		return models.User{}, errors.New("Cpf already exist")
	}

	existCel, errCell := s.repository.ExistCellphone(data.Cellphone)

	if errCell != nil {
		return models.User{}, errCell
	}

	if existCel {
		return models.User{}, errors.New("Cellphone already exist")
	}

	newUser := models.NewUser(data.Name,
		data.Email,
		data.Password,
		data.Cellphone,
		data.Cpf)

	err := newUser.Prepare()

	if err != nil {
		return models.User{}, err
	}

	return s.repository.Add(*newUser)
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	return s.repository.GetByEmail(email)
}

func NewUserService(r repositories.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}
