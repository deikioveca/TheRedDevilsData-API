package model


type Country struct {
	ID		uint	`gorm:"primaryKey"`
	Name	string
	Code	string
}


type CountryDTO struct {
	Name	string	`json:"name"`
	Code	string	`json:"code"`
}


type CountryResponse struct {
	Results		int				`json:"results"`
	Response	[]CountryDTO	`json:"response"`
}