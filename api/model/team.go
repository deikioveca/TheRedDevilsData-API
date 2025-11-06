package model

type Team struct {
	ID			uint	`gorm:"primaryKey"`
	TeamID		int			
	TeamName	string		
	Code		string		
	Country		string		
	Founded		int			
	National	bool
	VenueID		int			
	VenueName	string		
	Address		string		
	City		string		
	Capacity	int			
	Surface		string			
}


type TeamDTO struct {
	TeamID		int			`json:"id"`
	TeamName	string		`json:"name"`
	Code		string		`json:"code"`
	Country		string		`json:"country"`
	Founded		int			`json:"founded"`
	National	bool		`json:"national"`
}


type VenueDTO struct {
	VenueID		int			`json:"id"`
	VenueName	string		`json:"name"`
	Address		string		`json:"address"`
	City		string		`json:"city"`
	Capacity	int			`json:"capacity"`
	Surface		string		`json:"surface"`
}


type TeamInfoDTO struct {
	Team		TeamDTO		`json:"team"`
	Venue		VenueDTO	`json:"venue"`
}


type TeamResponse struct {
	Results		int				`json:"results"`
	Response	[]TeamInfoDTO	`json:"response"`
}


type ManchesterUnitedTeamDTO struct {
	Name		string	`json:"name"`
	Code		string	`json:"code"`
	Country		string	`json:"country"`
	Founded		int		`json:"founded"`
	VenueName	string	`json:"venue_name"`
	Address		string	`json:"address"`
	City		string	`json:"city"`
	Capacity	int		`json:"capacity"`
	Surface		string	`json:"surface"`
}