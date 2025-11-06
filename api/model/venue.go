package model


type Venue struct {
	ID			uint		`gorm:"primaryKey"`
	VenueID		int			
	VenueName	string		
	Address		string		
	City		string		
	Capacity	int			
	Surface		string		
}


type VenueResponse struct {
	Results			int			`json:"results"`
	Response		[]VenueDTO	`json:"response"`
}