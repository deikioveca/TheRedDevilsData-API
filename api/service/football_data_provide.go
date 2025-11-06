package service

import (
	"errors"
	"sort"

	"github.com/deikioveca/TheRedDevilsData/api/model"
	"gorm.io/gorm"
)

var (
	ErrCountryNotFound = errors.New("country with that name not found")

	ErrManchesterUnitedNotFound = errors.New("manchester united not found in database")

	ErrStandingNotFound = errors.New("standings for this season not found")

	ErrFixtureNotFound = errors.New("fixtures for this season not found")
)


type TeamStats interface {
	GetTeamStatsGames(season int) (*model.ManchesterUnitedGamesDTO, error)
	GetTeamStatsGoals(season int) (*model.ManchesterUnitedGoalsDTO, error)
	GetTeamStatsStreak(season int) (*model.ManchesterUnitedStreakDTO, error)
	GetTeamStatsBiggest(season int) (*model.ManchesterUnitedBiggestDTO, error)
	GetTeamStatsCleanSheet(season int) (*model.ManchesterUnitedCleanSheetDTO, error)
	GetTeamStatsFailedToScore(season int) (*model.ManchesterUnitedFailedScoringDTO, error)
	GetTeamStatsPenalty(season int) (*model.ManchesterUnitedPenaltyDTO, error)
	GetTeamStatsCards(season int) (*model.ManchesterUnitedCardsDTO, error)
	GetTeamStatsLineup(season int) (*model.ManchesterUnitedLineupDTO, error)
}


type Venue interface {
	GetVenues() (*model.VenueResponse, error)
	GetVenuesByCity(city string) (*model.VenueResponse, error)
	GetVenuesBiggestAndSmallest() (*model.VenueResponse, error)
}


type DataProvider interface {
	GetCountries() (*model.CountryResponse, error)
	GetCountryByName(countryName string) (*model.CountryDTO, error)

	GetLeagues() ([]*model.ManchesterUnitedLeaguesDTO, error)

	GetTeam() (*model.ManchesterUnitedTeamDTO, error)

	TeamStats

	Venue

	GetStandingsBySeason(season int) (*model.ManchesterUnitedStandingsDTO, error)

	GetFixturesBySeason(season int) ([]*model.ManchesterUnitedFixturesDTO, error)

	GetInjuriesBySeason(season int) (map[int][]*model.ManchesterUnitedInjuriesDTO, error)

	GetSquad() (*model.ManchesterUnitedSquadDTO, error)
}


func (s *service) GetCountries() (*model.CountryResponse, error) {
	var countries []model.Country
	if err := s.db.Find(&countries).Error; err != nil {
		return nil, err
	}

	countryResponse := &model.CountryResponse{Results: len(countries)}
	
	for _, c := range countries {
		countryDTO := &model.CountryDTO{
			Name: c.Name,
			Code: c.Code,
		}
		countryResponse.Response = append(countryResponse.Response, *countryDTO)
	}

	return countryResponse, nil
}


func (s *service) GetCountryByName(countryName string) (*model.CountryDTO, error) {
	var country model.Country
	if err := s.db.Where("name = ?", countryName).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCountryNotFound
		}
		return nil, err
	}

	return &model.CountryDTO{Name: country.Name, Code: country.Code}, nil
}


func (s *service) GetLeagues() ([]*model.ManchesterUnitedLeaguesDTO, error) {
	var leagues []model.League
	if err := s.db.Find(&leagues).Error; err != nil {
		return nil, err
	}

	ManUtdLeagues := []*model.ManchesterUnitedLeaguesDTO{}

	for _, l := range leagues {
		manUtdLeagueDTO := &model.ManchesterUnitedLeaguesDTO{
			Name: 		l.Name,
			Type: 		l.Type,
			Country: 	l.Country,
			Year: 		l.Year,
			Start: 		l.Start,
			End: 		l.End,
		}
		ManUtdLeagues = append(ManUtdLeagues, manUtdLeagueDTO)
	}

	return ManUtdLeagues, nil
}


func (s *service) GetTeam() (*model.ManchesterUnitedTeamDTO, error) {
	var team model.Team
	if err := s.db.First(&team, 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrManchesterUnitedNotFound
		}
		return nil, err
	}

	ManchesterUnited := &model.ManchesterUnitedTeamDTO{
		Name: 		team.TeamName,
		Code: 		team.Code,
		Country: 	team.Country,
		Founded: 	team.Founded,
		VenueName: 	team.VenueName,
		Address:	team.Address,
		City: 		team.City,
		Capacity: 	team.Capacity,
		Surface: 	team.Surface,
	}

	return ManchesterUnited, nil
}


func (s *service) GetTeamStatsGames(season int) (*model.ManchesterUnitedGamesDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedGamesDTO := &model.ManchesterUnitedGamesDTO{
		Team:	*munTeamStatsDTO,
		PlayedHome: 	teamStats.PlayedHome,
		PlayedAway: 	teamStats.PlayedAway,
		PlayedTotal: 	teamStats.PlayedTotal,
		WinsHome:		teamStats.WinsHome,
		WinsAway:		teamStats.WinsAway,
		WinsTotal:  	teamStats.WinsTotal,
		DrawsHome:  	teamStats.DrawsHome,
		DrawsAway:  	teamStats.DrawsAway,
		DrawsTotal:  	teamStats.DrawsTotal,
		LosesHome:   	teamStats.LosesHome,
		LosesAway:   	teamStats.LosesAway,
		LosesTotal:  	teamStats.LosesTotal,
	}

	return manchesterUnitedGamesDTO, nil
}


func (s *service) GetTeamStatsGoals(season int) (*model.ManchesterUnitedGoalsDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedGoalsDTO := &model.ManchesterUnitedGoalsDTO{
		Team: *munTeamStatsDTO,
		GoalsForHome: 			teamStats.GoalsForHome,
		GoalsForAway: 			teamStats.GoalsForAway,
		GoalsForTotal: 			teamStats.GoalsForTotal,
		GoalsAgainstHome: 		teamStats.GoalsAgainstHome,
		GoalsAgainstAway: 		teamStats.GoalsAgainstAway,
		GoalsAgainstTotal: 		teamStats.GoalsAgainstTotal,
		GoalsForAvgHome: 		teamStats.GoalsForAvgHome,
		GoalsForAvgAway: 		teamStats.GoalsForAvgAway,
		GoalsForAvgTotal: 		teamStats.GoalsForAvgTotal,
		GoalsAgainstAvgHome:	teamStats.GoalsAgainstAvgHome,
		GoalsAgainstAvgAway: 	teamStats.GoalsAgainstAvgAway,
		GoalsAgainstAvgTotal: 	teamStats.GoalsAgainstAvgTotal,
	}

	return manchesterUnitedGoalsDTO, nil
}


func (s *service) GetTeamStatsStreak(season int) (*model.ManchesterUnitedStreakDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedStreakDTO := &model.ManchesterUnitedStreakDTO{
		Team: *munTeamStatsDTO,
		StreakWins: 	teamStats.StreakWins,
		StreakDraws: 	teamStats.StreakDraws,
		StreakLoses: 	teamStats.StreakLoses,
	}

	return manchesterUnitedStreakDTO, nil
}


func (s *service) GetTeamStatsBiggest(season int) (*model.ManchesterUnitedBiggestDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedBiggestDTO := &model.ManchesterUnitedBiggestDTO{
		Team:	*munTeamStatsDTO,
		BiggestWinHome: 			teamStats.BiggestWinHome,
		BiggestWinAway: 			teamStats.BiggestWinAway,
		BiggestLoseHome: 			teamStats.BiggestLoseHome,
		BiggestLoseAway: 			teamStats.BiggestLoseAway,
		BiggestGoalsForHome: 		teamStats.BiggestGoalsForHome,
		BiggestGoalsForAway: 		teamStats.BiggestGoalsForAway,
		BiggestGoalsAgainstHome: 	teamStats.BiggestGoalsAgainstHome,
		BiggestGoalsAgainstAway: 	teamStats.BiggestGoalsAgainstAway,
	}

	return manchesterUnitedBiggestDTO, nil
}


func (s *service) GetTeamStatsCleanSheet(season int) (*model.ManchesterUnitedCleanSheetDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedCleanSheetDTO := &model.ManchesterUnitedCleanSheetDTO{
		Team: *munTeamStatsDTO,
		CleanSheetHome: 	teamStats.CleanSheetHome,
		CleanSheetAway: 	teamStats.CleanSheetAway,
		CleanSheetTotal: 	teamStats.CleanSheetTotal,
	}

	return manchesterUnitedCleanSheetDTO, nil
}


func (s *service) GetTeamStatsFailedToScore(season int) (*model.ManchesterUnitedFailedScoringDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedFailedScoringDTO := &model.ManchesterUnitedFailedScoringDTO{
		Team: *munTeamStatsDTO,
		FailedToScoreHome: 	teamStats.FailedToScoreHome,
		FailedToScoreAway: 	teamStats.FailedToScoreAway,
		FailedToScoreTotal: teamStats.FailedToScoreTotal,
	}

	return manchesterUnitedFailedScoringDTO, nil
}


func (s *service) GetTeamStatsPenalty(season int) (*model.ManchesterUnitedPenaltyDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedPenaltyDTO := &model.ManchesterUnitedPenaltyDTO{
		Team: *munTeamStatsDTO,
		PenaltyScoredTotal: 	teamStats.PenaltyScoredTotal,
		PenaltyScoredPct: 		*teamStats.PenaltyScoredPct,
		PenaltyMissedTotal: 	teamStats.PenaltyMissedTotal,
		PenaltyMissedPct: 		*teamStats.PenaltyMissedPct,
		PenaltyTotal: 			teamStats.PenaltyTotal,
	}

	return manchesterUnitedPenaltyDTO, nil
}


func (s *service) GetTeamStatsCards(season int) (*model.ManchesterUnitedCardsDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	manchesterUnitedCardsDTO := &model.ManchesterUnitedCardsDTO{
		Team: *munTeamStatsDTO,
		YellowCardsTotal: 	*teamStats.YellowCardsTotal,
		RedCardsTotal: 		*teamStats.RedCardsTotal,
	}

	return manchesterUnitedCardsDTO, nil
}


func (s *service) GetTeamStatsLineup(season int) (*model.ManchesterUnitedLineupDTO, error) {
	teamStats, munTeamStatsDTO, err := s.getTeamStatsBySeason(season)
	if err != nil {
		return nil, err
	}

	lineups, err := s.getLineupsBySeason(teamStats.Season)
	if err != nil {
		return nil, err
	}

	lineupDTOs := make(map[int][]model.LineupDTO)
	for _, l := range lineups {
		lineupDTOs[teamStats.Season] = append(lineupDTOs[teamStats.Season], model.LineupDTO{Formation: l.Formation, Played: l.Played})
	}

	manchesterUnitedLineupDTO := &model.ManchesterUnitedLineupDTO{
		Team: *munTeamStatsDTO,
		Lineup: lineupDTOs,
	}

	return manchesterUnitedLineupDTO, nil
}


func (s *service) GetVenues() (*model.VenueResponse, error) {
	var venues []model.Venue
	if err := s.db.Find(&venues).Error; err != nil {
		return nil, err
	}

	venueResponse := &model.VenueResponse{Results: len(venues)}
	for _, v := range venues {
		venueDTO := &model.VenueDTO{
			VenueID: 	v.VenueID,
			VenueName: 	v.VenueName,
			Address: 	v.Address,
			City: 		v.City,
			Capacity: 	v.Capacity,
			Surface: 	v.Surface,
		}
		venueResponse.Response = append(venueResponse.Response, *venueDTO)
	}

	return venueResponse, nil
}


func (s *service) GetVenuesByCity(city string) (*model.VenueResponse, error) {
	var venues []model.Venue
	if err := s.db.Where("city = ?", city).Find(&venues).Error; err != nil {
		return nil, err
	}

	venueResponse := &model.VenueResponse{Results: len(venues)}
	for _, v := range venues {
		venueDTO := &model.VenueDTO{
			VenueID: 	v.VenueID,
			VenueName: 	v.VenueName,
			Address: 	v.Address,
			City: 		v.City,
			Capacity: 	v.Capacity,
			Surface: 	v.Surface,
		}
		venueResponse.Response = append(venueResponse.Response, *venueDTO)
	}

	return venueResponse, nil
}


func (s *service) GetVenuesBiggestAndSmallest() (*model.VenueResponse, error) {
	var venues []model.Venue
	if err := s.db.Find(&venues).Error; err != nil {
		return nil, err
	}

	biggestAndSmallestCapacities := []int{}
	for _, v := range venues {
		biggestAndSmallestCapacities = append(biggestAndSmallestCapacities, v.Capacity)
	}

	sort.Ints(biggestAndSmallestCapacities)

	var venueDTO []model.VenueDTO
	for _, v := range venues {
		if v.Capacity == biggestAndSmallestCapacities[0] || v.Capacity == biggestAndSmallestCapacities[len(biggestAndSmallestCapacities) - 1] {
			venueDTO = append(venueDTO, model.VenueDTO{
				VenueID: 	v.VenueID,
				VenueName: 	v.VenueName,
				Address: 	v.Address,
				City: 		v.City,
				Capacity: 	v.Capacity,
				Surface: 	v.Surface,
			})
		}
	}

	return &model.VenueResponse{Results: len(venueDTO), Response: venueDTO}, nil
}


func (s *service) GetStandingsBySeason(season int) (*model.ManchesterUnitedStandingsDTO, error) {
	var standing model.Standing
	if err := s.db.Where("season = ?", season).First(&standing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStandingNotFound
		}
		return nil, err
	}

	manchesterUnitedStandingsDTO := &model.ManchesterUnitedStandingsDTO{
		LeagueName: 		standing.LeagueName,
		Season: 			standing.Season,
		TeamName: 			standing.TeamName,
		Rank: 				standing.Rank,
		Points: 			standing.Points,
		GoalsDiff: 			standing.GoalsDiff,
		Description: 		standing.Description,
		PlayedAll: 			standing.PlayedAll,
		WinsAll: 			standing.WinsAll,
		DrawsAll: 			standing.DrawsAll,
		LosesAll: 			standing.LosesAll,
		GoalsForAll: 		standing.GoalsForAll,
		GoalsAgainstAll: 	standing.GoalsAgainstAll,
		PlayedHome: 		standing.PlayedHome,
		WinsHome: 			standing.WinsHome,
		DrawsHome: 			standing.DrawsHome,
		LosesHome: 			standing.LosesHome,
		GoalsForHome: 		standing.GoalsForHome,
		GoalsAgainstHome: 	standing.GoalsAgainstHome,
		PlayedAway: 		standing.PlayedAway,
		WinsAway: 			standing.WinsAway,
		DrawsAway: 			standing.DrawsAway,
		LosesAway: 			standing.LosesAway,
		GoalsForAway: 		standing.GoalsForAway,
		GoalsAgainstAway: 	standing.GoalsAgainstAway,
	}

	return manchesterUnitedStandingsDTO, nil
}


func (s *service) GetFixturesBySeason(season int) ([]*model.ManchesterUnitedFixturesDTO, error) {
	var fixtures []model.Fixture
	if err := s.db.Where("season = ?", season).Find(&fixtures).Error; err != nil {
		return nil, err
	}

	manchesterUnitedFixturesDTO := []*model.ManchesterUnitedFixturesDTO{}
	for _, fixture := range fixtures {
		manchesterUnitedFixturesDTO = append(manchesterUnitedFixturesDTO, &model.ManchesterUnitedFixturesDTO{
			Referee: 			fixture.Referee,
    		Date: 				fixture.Date,
    		VenueName: 			fixture.VenueName,
    		VenueCity:    		fixture.VenueCity,
    		StatusLong: 		fixture.StatusLong,
    		StatusShort: 		fixture.StatusShort,
    		StatusElapsed: 		fixture.StatusElapsed,
    		StatusExtra: 		safeString(fixture.StatusExtra),
    		LeagueName: 		fixture.LeagueName,
    		Country: 			fixture.Country,
    		Season: 			fixture.Season,
    		Round: 				fixture.Round,
			HomeTeamName: 		fixture.HomeTeamName,
    		HomeWinner: 		safeBool(fixture.HomeWinner),
    		AwayTeamName: 		fixture.AwayTeamName,
    		AwayWinner: 		safeBool(fixture.AwayWinner),
    		GoalsHome: 			fixture.GoalsHome,
    		GoalsAway: 			fixture.GoalsAway,
    		HalftimeHome: 		safeInt(fixture.HalftimeHome),
			HalftimeAway:    	safeInt(fixture.HalftimeAway),
			FulltimeHome:    	safeInt(fixture.FulltimeHome),
			FulltimeAway:    	safeInt(fixture.FulltimeAway),
			ExtratimeHome:   	safeInt(fixture.ExtratimeHome),
			ExtratimeAway:   	safeInt(fixture.ExtratimeAway),
			PenaltyHome:     	safeInt(fixture.PenaltyHome),
			PenaltyAway:     	safeInt(fixture.PenaltyAway),
		})
	}

	return manchesterUnitedFixturesDTO, nil
}


func (s *service) GetInjuriesBySeason(season int) (map[int][]*model.ManchesterUnitedInjuriesDTO, error) {
	var injuries []model.Injury
	if err := s.db.Where("season = ?", season).Find(&injuries).Error; err != nil {
		return nil, err
	}

	manchesterUnitedInjuriesDTO := []*model.ManchesterUnitedInjuriesDTO{}
	for _, i := range injuries {
		manchesterUnitedInjuriesDTO = append(manchesterUnitedInjuriesDTO, &model.ManchesterUnitedInjuriesDTO{
			PlayerName: 	i.PlayerName,
			Type: 			i.Type,
			Reason: 		i.Reason,
			FixtureDate: 	i.FixtureDate,
			LeagueName: 	i.LeagueName,
			Country: 		i.Country,
			Season: 		i.Season,
		})
	}

	dataDTO := make(map[int][]*model.ManchesterUnitedInjuriesDTO)
	dataDTO[len(injuries)] = manchesterUnitedInjuriesDTO

	return dataDTO, nil
}


func (s *service) GetSquad() (*model.ManchesterUnitedSquadDTO, error) {
	var squad []model.Squad
	if err := s.db.Find(&squad).Error; err != nil {
		return nil, err
	}

	manchesterUnitedSquadDTO := &model.ManchesterUnitedSquadDTO{SquadDepth: len(squad)}
	for _, player := range squad {
		manchesterUnitedSquadDTO.Footballers = append(manchesterUnitedSquadDTO.Footballers, model.ManchesterUnitedFootballerDTO{
			PlayerName: player.PlayerName,
			Age: 		player.Age,
			Number: 	player.Number,
			Position: 	player.Position,
		})
	}

	return manchesterUnitedSquadDTO, nil
}