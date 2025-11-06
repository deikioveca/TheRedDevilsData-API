package model

type Squad struct {
	ID        		uint   `gorm:"primaryKey"`
	TeamID    		int
	TeamName  		string
	TeamLogo  		string
	PlayerID  		int
	PlayerName 		string
	Age       		int
	Number    		int
	Position  		string
	PlayerPhoto 	string
}


type SquadTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}


type SquadPlayer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Number   int    `json:"number"`
	Position string `json:"position"`
	Photo    string `json:"photo"`
}


type SquadDTO struct {
	Team    SquadTeam     `json:"team"`
	Players []SquadPlayer `json:"players"`
}



type SquadResponse struct {
	Response []SquadDTO `json:"response"`
}


type ManchesterUnitedFootballerDTO struct {
	PlayerName 		string	`json:"player_name"`
	Age       		int		`json:"age"`
	Number    		int		`json:"number"`
	Position  		string	`json:"position"`
}


type ManchesterUnitedSquadDTO struct {
	SquadDepth		int								`json:"squad_depth"`
	Footballers		[]ManchesterUnitedFootballerDTO	`json:"footballers"`
}