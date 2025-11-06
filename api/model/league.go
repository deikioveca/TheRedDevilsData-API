package model


type League struct {
	ID			uint	`gorm:"primaryKey"`
	LeagueID	int
	Name		string
	Type		string
	Country		string
	CountryCode	string
	Year		int		
	Start		string	
	End			string	
	Current		bool
	TeamID     	int	
}


type SeasonDTO struct {
	Year	int		`json:"year"`
	Start	string	`json:"start"`
	End		string	`json:"end"`
	Current	bool	`json:"current"`
}


type LeagueInfoDTO struct {
	LeagueID	int			`json:"id"`
	Name		string		`json:"name"`
	Type		string		`json:"type"`
}


type LeagueDTO struct {
	League		LeagueInfoDTO	`json:"league"`
	Country		CountryDTO		`json:"country"`
	Seasons		[]SeasonDTO		`json:"seasons"`
}


type LeagueResponse struct {
	Results		int				`json:"results"`
	Response	[]LeagueDTO		`json:"response"`
}


type ManchesterUnitedLeaguesDTO struct {
	Name		string	`json:"name"`
	Type		string	`json:"type"`
	Country		string	`json:"country"`
	Year		int		`json:"year"`
	Start		string	`json:"start"`
	End			string	`json:"end"`
}