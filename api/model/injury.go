package model

type Injury struct {
	ID uint `gorm:"primaryKey"`

	PlayerID   	int
	PlayerName 	string
	PlayerPhoto string
	Type       	string
	Reason     	string

	TeamID   int
	TeamName string
	TeamLogo string

	FixtureID  			int
	FixtureDate 		string
	FixtureTimestamp 	int64
	FixtureTimezone 	string

	LeagueID   int
	LeagueName string
	Country    string
	Season     int
	LeagueLogo string
	Flag       string
}



type InjuryPlayerDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
}


type InjuryTeamDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Logo  string `json:"logo"`
}


type InjuryFixtureDTO struct {
	ID        int    `json:"id"`
	Timezone  string `json:"timezone"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
}


type InjuryLeagueDTO struct {
	ID      int    `json:"id"`
	Season  int    `json:"season"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
}


type InjuryDTO struct {
	Player  InjuryPlayerDTO  `json:"player"`
	Team    InjuryTeamDTO    `json:"team"`
	Fixture InjuryFixtureDTO `json:"fixture"`
	League  InjuryLeagueDTO  `json:"league"`
}


type InjuryResponse struct {
	Response []InjuryDTO `json:"response"`
}


type ManchesterUnitedInjuriesDTO struct {
	PlayerName 			string	`json:"player_name"`
	Type       			string	`json:"type"`
	Reason     			string	`json:"reason"`
	FixtureDate 		string	`json:"fixture_date"`
	LeagueName 			string	`json:"league_name"`
	Country    			string	`json:"country"`
	Season     			int		`json:"season"`
}