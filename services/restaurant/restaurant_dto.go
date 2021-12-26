package restaurant

type RestaurantDTO struct {
	Name      string         `json:"name"`
	Cnpj      string         `json:"cnpj"`
	IsOpen    bool           `json:"is_open"`
}