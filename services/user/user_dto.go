package user

type UserDto struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Cellphone string `json:"cellphone"`
	Cpf       string `json:"cpf"`
}
