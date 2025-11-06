package model

type Fixture struct {
	ID        uint   `gorm:"primaryKey"`

	FixtureID  int
	Referee    string
	Timezone   string
	Date       string
	Timestamp  int64
	PeriodFirst  *int64
	PeriodSecond *int64

	VenueID   int
	VenueName string
	VenueCity string

	StatusLong   string
	StatusShort  string
	StatusElapsed int
	StatusExtra   *string

	LeagueID   int
	LeagueName string
	Country    string
	Season     int
	Round      string
	Standings  bool

	HomeTeamID   int
	HomeTeamName string
	HomeTeamLogo string
	HomeWinner   *bool

	AwayTeamID   int
	AwayTeamName string
	AwayTeamLogo string
	AwayWinner   *bool

	GoalsHome int
	GoalsAway int

	HalftimeHome  *int
	HalftimeAway  *int
	FulltimeHome  *int
	FulltimeAway  *int
	ExtratimeHome *int
	ExtratimeAway *int
	PenaltyHome   *int
	PenaltyAway   *int
}


type FixtureScore struct {
	Halftime  FixtureGoals `json:"halftime"`
	Fulltime  FixtureGoals `json:"fulltime"`
	Extratime FixtureGoals `json:"extratime"`
	Penalty   FixtureGoals `json:"penalty"`
}


type FixtureGoals struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}


type FixtureTeam struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Logo   string  `json:"logo"`
	Winner *bool   `json:"winner"`
}


type FixtureTeams struct {
	Home FixtureTeam `json:"home"`
	Away FixtureTeam `json:"away"`
}


type FixtureLeague struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Logo      string `json:"logo"`
	Flag      string `json:"flag"`
	Season    int    `json:"season"`
	Round     string `json:"round"`
	Standings bool   `json:"standings"`
}


type FixtureStatus struct {
	Long    string  `json:"long"`
	Short   string  `json:"short"`
	Elapsed int     `json:"elapsed"`
	Extra   *string `json:"extra"`
}


type FixtureVenue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}


type FixturePeriods struct {
	First  *int64 `json:"first"`
	Second *int64 `json:"second"`
}


type FixtureInfo struct {
	ID        int            `json:"id"`
	Referee   string         `json:"referee"`
	Timezone  string         `json:"timezone"`
	Date      string         `json:"date"`
	Timestamp int64          `json:"timestamp"`
	Periods   FixturePeriods `json:"periods"`
	Venue     FixtureVenue   `json:"venue"`
	Status    FixtureStatus  `json:"status"`
}


type FixtureDTO struct {
	Fixture FixtureInfo  `json:"fixture"`
	League  FixtureLeague `json:"league"`
	Teams   FixtureTeams  `json:"teams"`
	Goals   FixtureGoals  `json:"goals"`
	Score   FixtureScore  `json:"score"`
}


type FixtureResponse struct {
	Response []FixtureDTO `json:"response"`
}


type ManchesterUnitedFixturesDTO struct {
	Referee    			string	`json:"referee"`
	Date       			string	`json:"date"`
	VenueName 			string	`json:"venue_name"`
	VenueCity 			string	`json:"venue_city"`
	StatusLong   		string	`json:"status_long"`
	StatusShort  		string	`json:"status_short"`
	StatusElapsed 		int		`json:"status_elapsed"`
	StatusExtra   		string	`json:"status_extra"`
	LeagueName 			string	`json:"league_name"`
	Country    			string	`json:"country"`
	Season     			int		`json:"season"`
	Round      			string	`json:"round"`
	HomeTeamName 		string	`json:"home_team_name"`
	HomeWinner   		bool	`json:"home_winner"`
	AwayTeamName 		string	`json:"away_team_name"`
	AwayWinner   		bool	`json:"away_winner"`
	GoalsHome 			int		`json:"goals_home"`
	GoalsAway 			int		`json:"goals_away"`
	HalftimeHome  		int		`json:"half_time_home"`
	HalftimeAway  		int		`json:"half_time_away"`
	FulltimeHome  		int		`json:"full_time_home"`
	FulltimeAway  		int		`json:"full_time_away"`
	ExtratimeHome 		int		`json:"extra_time_home"`
	ExtratimeAway 		int		`json:"extra_time_away"`
	PenaltyHome   		int		`json:"penalty_home"`
	PenaltyAway   		int		`json:"penalty_away"`
}