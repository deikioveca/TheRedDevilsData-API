package model

type TeamStats struct {
	ID			uint	`gorm:"primaryKey"`

	TeamID      int
	TeamName    string
	LeagueID    int
	LeagueName  string
	Country     string
	Season      int
	Form        string

	PlayedHome  int
	PlayedAway  int
	PlayedTotal int
	WinsHome    int
	WinsAway    int
	WinsTotal   int
	DrawsHome   int
	DrawsAway   int
	DrawsTotal  int
	LosesHome   int
	LosesAway   int
	LosesTotal  int

	GoalsForHome     int
	GoalsForAway     int
	GoalsForTotal    int
	GoalsAgainstHome int
	GoalsAgainstAway int
	GoalsAgainstTotal int

	GoalsForAvgHome     string
	GoalsForAvgAway     string
	GoalsForAvgTotal    string
	GoalsAgainstAvgHome string
	GoalsAgainstAvgAway string
	GoalsAgainstAvgTotal string

	StreakWins  int
	StreakDraws int
	StreakLoses int
	BiggestWinHome string
	BiggestWinAway string
	BiggestLoseHome string
	BiggestLoseAway string
	BiggestGoalsForHome int
	BiggestGoalsForAway int
	BiggestGoalsAgainstHome int
	BiggestGoalsAgainstAway int

	CleanSheetHome  int
	CleanSheetAway  int
	CleanSheetTotal int
	FailedToScoreHome  int
	FailedToScoreAway  int
	FailedToScoreTotal int

	PenaltyScoredTotal int
	PenaltyScoredPct   *string
	PenaltyMissedTotal int
	PenaltyMissedPct   *string
	PenaltyTotal       int

	YellowCardsTotal *int
	RedCardsTotal    *int
}


type Lineup struct {
	ID			uint	`gorm:"primaryKey"`
	Season		int
	Formation	string
	Played		int
}


type TeamLeagueDTO struct {
	TeamID		int		`json:"id"`
	Name		string	`json:"name"`
	Country		string	`json:"country"`
	Season		int		`json:"season"`
}


type ManchesterUnitedTeam struct {
	ID		int			`json:"id"`
	Name	string		`json:"name"`
}


type FixturesRecordDTO struct {
	Home		int		`json:"home"`
	Away		int		`json:"away"`
	Total		int		`json:"total"`
}


type FixturesDTO struct {
	Played		FixturesRecordDTO	`json:"played"`
	Wins		FixturesRecordDTO	`json:"wins"`
	Draws		FixturesRecordDTO	`json:"draws"`
	Loses		FixturesRecordDTO	`json:"loses"`
}


type GoalsDTO struct {
	For     GoalsSide `json:"for"`
	Against GoalsSide `json:"against"`
}


type GoalsSide struct {
	Total      GoalsTotalStats           `json:"total"`
	Average    GoalsAverageStats         `json:"average"`
	Minute     map[string]GoalsMinute    `json:"minute"`      
	UnderOver  map[string]GoalsUnderOver `json:"under_over"`  
}


type GoalsTotalStats struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}


type GoalsAverageStats struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}


type GoalsMinute struct {
	Total      *int    `json:"total"`      
	Percentage *string `json:"percentage"` 
}


type GoalsUnderOver struct {
	Over  int `json:"over"`
	Under int `json:"under"`
}


type BiggestDTO struct {
	Streak  BiggestStreak  `json:"streak"`
	Wins    BiggestResult  `json:"wins"`
	Loses   BiggestResult  `json:"loses"`
	Goals   BiggestGoals   `json:"goals"`
}


type BiggestStreak struct {
	Wins  int `json:"wins"`
	Draws int `json:"draws"`
	Loses int `json:"loses"`
}


type BiggestResult struct {
	Home string `json:"home"`
	Away string `json:"away"`
}


type BiggestGoals struct {
	For     BiggestHomeAway `json:"for"`
	Against BiggestHomeAway `json:"against"`
}


type BiggestHomeAway struct {
	Home int `json:"home"`
	Away int `json:"away"`
}


type CleanSheetDTO struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}


type FailedToScoreDTO struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}


type PenaltyDTO struct {
	Scored PenaltyDetail `json:"scored"`
	Missed PenaltyDetail `json:"missed"`
	Total  int           `json:"total"`
}

type PenaltyDetail struct {
	Total      int     `json:"total"`
	Percentage *string `json:"percentage"` 
}


type LineupDTO struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}


type MinuteCardStat struct {
	Total      *int    `json:"total"`
	Percentage *string `json:"percentage"`
}


type CardDistribution struct {
	M0_15    MinuteCardStat `json:"0-15"`
	M16_30   MinuteCardStat `json:"16-30"`
	M31_45   MinuteCardStat `json:"31-45"`
	M46_60   MinuteCardStat `json:"46-60"`
	M61_75   MinuteCardStat `json:"61-75"`
	M76_90   MinuteCardStat `json:"76-90"`
	M91_105  MinuteCardStat `json:"91-105"`
	M106_120 MinuteCardStat `json:"106-120"`
}


type CardsDTO struct {
	Yellow CardDistribution `json:"yellow"`
	Red    CardDistribution `json:"red"`
}


type TeamStatsInfoDTO struct {
	League			TeamLeagueDTO			`json:"league"`
	Team			ManchesterUnitedTeam	`json:"team"`
	Form			string					`json:"form"`
	Fixtures		FixturesDTO				`json:"fixtures"`
	Goals			GoalsDTO				`json:"goals"`
	Biggest			BiggestDTO				`json:"biggest"`
	CleanSheet		CleanSheetDTO			`json:"clean_sheet"`
	FailedToScore	FailedToScoreDTO		`json:"failed_to_score"`
	Penalty			PenaltyDTO				`json:"penalty"`
	Lineup			[]LineupDTO				`json:"lineups"`
	Cards			CardsDTO				`json:"cards"`
}



type TeamStatsResponse struct {
	Results		int					`json:"results"`
	Response	TeamStatsInfoDTO	`json:"response"`
}


type ManchesterUnitedTeamStatsDTO struct {
	Name		string		`json:"name"`
	League		string		`json:"league"`
	Season		int			`json:"season"`
	Form		string		`json:"form"`
}


type ManchesterUnitedGamesDTO struct {
	Team		ManchesterUnitedTeamStatsDTO	`json:"team"`
	PlayedHome  int								`json:"played_home"`
	PlayedAway  int								`json:"played_away"`
	PlayedTotal int								`json:"played_total"`
	WinsHome    int								`json:"wins_home"`
	WinsAway    int								`json:"wins_away"`
	WinsTotal   int								`json:"wins_total"`
	DrawsHome   int								`json:"draws_home"`
	DrawsAway   int								`json:"draws_away"`
	DrawsTotal  int								`json:"draws_total"`
	LosesHome   int								`json:"loses_home"`
	LosesAway   int								`json:"loses_away"`
	LosesTotal  int								`json:"loses_total"`
}


type ManchesterUnitedGoalsDTO struct {
	Team					ManchesterUnitedTeamStatsDTO	`json:"team"`
	GoalsForHome     		int								`json:"goals_for_home"`
	GoalsForAway     		int								`json:"goals_for_away"`
	GoalsForTotal    		int								`json:"goals_for_total"`
	GoalsAgainstHome 		int								`json:"goals_against_home"`
	GoalsAgainstAway 		int								`json:"goals_against_away"`
	GoalsAgainstTotal 		int								`json:"goals_against_total"`
	GoalsForAvgHome     	string							`json:"goals_for_avg_home"`
	GoalsForAvgAway     	string							`json:"goals_for_avg_away"`
	GoalsForAvgTotal    	string							`json:"goals_for_avg_total"`
	GoalsAgainstAvgHome 	string							`json:"goals_against_avg_home"`
	GoalsAgainstAvgAway 	string							`json:"goals_against_avg_away"`
	GoalsAgainstAvgTotal 	string							`json:"goals_against_avg_total"`
}


type ManchesterUnitedStreakDTO struct {
	Team		ManchesterUnitedTeamStatsDTO	`json:"team"`
	StreakWins  int								`json:"streak_wins"`
	StreakDraws int								`json:"streak_draws"`
	StreakLoses int								`json:"streak_loses"`
}


type ManchesterUnitedBiggestDTO struct {
	Team					ManchesterUnitedTeamStatsDTO	`json:"team"`
	BiggestWinHome 			string							`json:"biggest_win_home"`
	BiggestWinAway 			string							`json:"biggest_win_away"`
	BiggestLoseHome 		string							`json:"biggest_lose_home"`
	BiggestLoseAway 		string							`json:"biggest_lose_away"`
	BiggestGoalsForHome 	int								`json:"biggest_goals_for_home"`
	BiggestGoalsForAway 	int								`json:"biggest_goals_for_away"`
	BiggestGoalsAgainstHome int								`json:"biggest_goals_against_home"`
	BiggestGoalsAgainstAway int								`json:"biggest_goals_against_away"`
}


type ManchesterUnitedCleanSheetDTO struct {
	Team			ManchesterUnitedTeamStatsDTO	`json:"team"`
	CleanSheetHome  int								`json:"clean_sheet_home"`
	CleanSheetAway  int								`json:"clean_sheet_away"`
	CleanSheetTotal int								`json:"clean_sheet_total"`
}


type ManchesterUnitedFailedScoringDTO struct {
	Team				ManchesterUnitedTeamStatsDTO	`json:"team"`
	FailedToScoreHome  	int								`json:"failed_to_score_home"`
	FailedToScoreAway  	int								`json:"failed_to_score_away"`
	FailedToScoreTotal 	int								`json:"failed_to_score_total"`
}


type ManchesterUnitedPenaltyDTO struct {
	Team				ManchesterUnitedTeamStatsDTO	`json:"team"`
	PenaltyScoredTotal 	int								`json:"penalty_scored_total"`
	PenaltyScoredPct   	string							`json:"penalty_scored_percentage"`
	PenaltyMissedTotal 	int								`json:"penalty_missed_total"`
	PenaltyMissedPct   	string							`json:"penalty_missed_percentage"`
	PenaltyTotal       	int								`json:"penalty_total"`
}


type ManchesterUnitedCardsDTO struct {
	Team				ManchesterUnitedTeamStatsDTO	`json:"team"`
	YellowCardsTotal 	int								`json:"yellow_cards_total"`
	RedCardsTotal    	int								`json:"red_cards_total"`
}


type ManchesterUnitedLineupDTO struct {
	Team				ManchesterUnitedTeamStatsDTO	`json:"team"`
	Lineup				map[int][]LineupDTO				`json:"lineups"`
}