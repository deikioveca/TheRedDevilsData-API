package model

type Standing struct {
	ID		uint		`gorm:"primaryKey"`

	LeagueID    int
	LeagueName  string
	Country     string
	Season      int
	TeamID      int
	TeamName    string
	TeamLogo    string
	Rank        int
	Points      int
	GoalsDiff   int
	GroupName   string
	Form        string
	Status      string
	Description string

	PlayedAll int
	WinsAll   int
	DrawsAll  int
	LosesAll  int
	GoalsForAll int
	GoalsAgainstAll int

	PlayedHome int
	WinsHome   int
	DrawsHome  int
	LosesHome  int
	GoalsForHome int
	GoalsAgainstHome int

	PlayedAway int
	WinsAway   int
	DrawsAway  int
	LosesAway  int
	GoalsForAway int
	GoalsAgainstAway int

	UpdatedAt string
}


type StandingEntry struct {
	Rank        int          	`json:"rank"`
	Team        StandingTeam 	`json:"team"`
	Points      int          	`json:"points"`
	GoalsDiff   int          	`json:"goalsDiff"`
	Group       string       	`json:"group"`
	Form        string       	`json:"form"`
	Status      string       	`json:"status"`
	Description string       	`json:"description"`
	All         StandingStats 	`json:"all"`
	Home        StandingStats 	`json:"home"`
	Away        StandingStats 	`json:"away"`
	Update      string       	`json:"update"`
}


type StandingTeam struct {
	ID   	int    	`json:"id"`
	Name 	string 	`json:"name"`
}


type StandingGoals struct {
	For		int		`json:"for"`
	Against int		`json:"against"`
}


type StandingStats struct {
	Played int          	`json:"played"`
	Win    int          	`json:"win"`
	Draw   int          	`json:"draw"`
	Lose   int          	`json:"lose"`
	Goals  StandingGoals 	`json:"goals"`
}


type StandingInfoDTO struct {
	ID        	int               	`json:"id"`
	Name      	string            	`json:"name"`
	Country   	string            	`json:"country"`
	Season    	int               	`json:"season"`
	Standings	[][]StandingEntry	`json:"standings"`
}


type StandingDTO struct {
	StandingInfo	StandingInfoDTO	`json:"league"`
}


type StandingResponse struct {
	Response		[]StandingDTO	`json:"response"`	
}


type ManchesterUnitedStandingsDTO struct {
	LeagueName  		string	`json:"league_name"`
	Season      		int		`json:"season"`
	TeamName    		string	`json:"team_name"`
	Rank        		int		`json:"rank"`
	Points      		int		`json:"points"`
	GoalsDiff   		int		`json:"goals_diff"`
	Description 		string	`json:"description"`
	PlayedAll 			int		`json:"played_all"`
	WinsAll   			int		`json:"wins_all"`
	DrawsAll  			int		`json:"draws_all"`
	LosesAll  			int		`json:"loses_all"`
	GoalsForAll 		int		`json:"goals_for_all"`
	GoalsAgainstAll 	int		`json:"goals_against_all"`
	PlayedHome 			int		`json:"played_home"`
	WinsHome   			int		`json:"wins_home"`
	DrawsHome  			int		`json:"draws_home"`
	LosesHome  			int		`json:"loses_home"`
	GoalsForHome 		int		`json:"goals_for_home"`
	GoalsAgainstHome 	int		`json:"goals_against_home"`
	PlayedAway 			int		`json:"played_away"`
	WinsAway   			int		`json:"wins_away"`
	DrawsAway  			int		`json:"draws_away"`
	LosesAway  			int		`json:"loses_away"`
	GoalsForAway 		int		`json:"goals_for_away"`
	GoalsAgainstAway 	int		`json:"goals_against_away"`
}