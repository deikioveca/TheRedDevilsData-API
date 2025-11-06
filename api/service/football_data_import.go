package service

import (
	"github.com/deikioveca/TheRedDevilsData/api/model"
)


type DataImporter interface {
	SaveCountries() (*model.CountryResponse, error)
	SaveLeaguesForTeam() (*model.LeagueResponse, error)
	SaveTeam() (*model.TeamResponse, error)
	SaveTeamStats() ([]*model.TeamStatsResponse, error)
	SaveVenues() (*model.VenueResponse, error)
	SaveStandings() ([]*model.StandingResponse, error)
	SaveFixtures() ([]*model.FixtureResponse, error)
	SaveInjuries() ([]*model.InjuryResponse, error)
	SaveSquad() (*model.SquadResponse, error)
}


func (s *service) SaveCountries() (*model.CountryResponse, error) {
	countries, err := s.client.FetchCountries()
	if err != nil {
		return nil, err
	}

	for _, c := range countries.Response {
		country := &model.Country{
			Name: c.Name,
			Code: c.Code,
		}
		s.db.Create(&country)
	}

	return countries, nil
}


func (s *service) SaveLeaguesForTeam() (*model.LeagueResponse, error) {
	leagues, err := s.client.FetchAllLeaguesForTeam()
	if err != nil {
		return nil, err
	}

	for _, leagueDTO := range leagues.Response {
		for _, season := range leagueDTO.Seasons {
			league := &model.League{
				LeagueID: 		leagueDTO.League.LeagueID,
				Name: 			leagueDTO.League.Name,
				Type: 			leagueDTO.League.Type,
				Country: 		leagueDTO.Country.Name,
				CountryCode: 	leagueDTO.Country.Code,
				Year: 			season.Year,
				Start: 			season.Start,
				End: 			season.End,
				Current: 		season.Current,
				TeamID: 		33,
			}
			s.db.Create(&league)
		}
	}

	return leagues, nil
}


func (s *service) SaveTeam() (*model.TeamResponse, error) {
	team, err := s.client.FetchTeam()
	if err != nil {
		return nil, err
	}

	for _, res := range team.Response {
		manchesterUnited := &model.Team{
			TeamID: 	res.Team.TeamID,
			TeamName: 	res.Team.TeamName,
			Code: 		res.Team.Code,
			Country: 	res.Team.Country,
			Founded: 	res.Team.Founded,
			National: 	res.Team.National,
			VenueID: 	res.Venue.VenueID,
			VenueName: 	res.Venue.VenueName,
			Address: 	res.Venue.Address,
			City: 		res.Venue.City,
			Capacity: 	res.Venue.Capacity,
			Surface: 	res.Venue.Surface,
		}
		s.db.Create(&manchesterUnited)
	} 

	return team, nil
}


func (s *service) SaveTeamStats() ([]*model.TeamStatsResponse, error) {
	stats, err := s.client.FetchTeamStats(33, 39)
	if err != nil {
		return nil, err
	}

	for _, seasonStats := range stats {
		if seasonStats == nil {
			continue
		}

		resp := seasonStats.Response

		teamStats := &model.TeamStats{
			TeamID:     resp.Team.ID,
			TeamName:   resp.Team.Name,
			LeagueID:   resp.League.TeamID, 
			LeagueName: resp.League.Name,
			Country:    resp.League.Country,
			Season:     resp.League.Season,
			Form:       resp.Form,

			PlayedHome:  resp.Fixtures.Played.Home,
			PlayedAway:  resp.Fixtures.Played.Away,
			PlayedTotal: resp.Fixtures.Played.Total,
			WinsHome:    resp.Fixtures.Wins.Home,
			WinsAway:    resp.Fixtures.Wins.Away,
			WinsTotal:   resp.Fixtures.Wins.Total,
			DrawsHome:   resp.Fixtures.Draws.Home,
			DrawsAway:   resp.Fixtures.Draws.Away,
			DrawsTotal:  resp.Fixtures.Draws.Total,
			LosesHome:   resp.Fixtures.Loses.Home,
			LosesAway:   resp.Fixtures.Loses.Away,
			LosesTotal:  resp.Fixtures.Loses.Total,

			GoalsForHome:      resp.Goals.For.Total.Home,
			GoalsForAway:      resp.Goals.For.Total.Away,
			GoalsForTotal:     resp.Goals.For.Total.Total,
			GoalsAgainstHome:  resp.Goals.Against.Total.Home,
			GoalsAgainstAway:  resp.Goals.Against.Total.Away,
			GoalsAgainstTotal: resp.Goals.Against.Total.Total,

			GoalsForAvgHome:      resp.Goals.For.Average.Home,
			GoalsForAvgAway:      resp.Goals.For.Average.Away,
			GoalsForAvgTotal:     resp.Goals.For.Average.Total,
			GoalsAgainstAvgHome:  resp.Goals.Against.Average.Home,
			GoalsAgainstAvgAway:  resp.Goals.Against.Average.Away,
			GoalsAgainstAvgTotal: resp.Goals.Against.Average.Total,

			StreakWins:             resp.Biggest.Streak.Wins,
			StreakDraws:            resp.Biggest.Streak.Draws,
			StreakLoses:            resp.Biggest.Streak.Loses,
			BiggestWinHome:         resp.Biggest.Wins.Home,
			BiggestWinAway:         resp.Biggest.Wins.Away,
			BiggestLoseHome:        resp.Biggest.Loses.Home,
			BiggestLoseAway:        resp.Biggest.Loses.Away,
			BiggestGoalsForHome:    resp.Biggest.Goals.For.Home,
			BiggestGoalsForAway:    resp.Biggest.Goals.For.Away,
			BiggestGoalsAgainstHome: resp.Biggest.Goals.Against.Home,
			BiggestGoalsAgainstAway: resp.Biggest.Goals.Against.Away,

			CleanSheetHome:     resp.CleanSheet.Home,
			CleanSheetAway:     resp.CleanSheet.Away,
			CleanSheetTotal:    resp.CleanSheet.Total,
			FailedToScoreHome:  resp.FailedToScore.Home,
			FailedToScoreAway:  resp.FailedToScore.Away,
			FailedToScoreTotal: resp.FailedToScore.Total,

			PenaltyScoredTotal: resp.Penalty.Scored.Total,
			PenaltyScoredPct:   resp.Penalty.Scored.Percentage,
			PenaltyMissedTotal: resp.Penalty.Missed.Total,
			PenaltyMissedPct:   resp.Penalty.Missed.Percentage,
			PenaltyTotal:       resp.Penalty.Total,
		}

		yellow := 0
		red := 0

		for _, minuteRange := range []model.MinuteCardStat{
			resp.Cards.Yellow.M0_15,
			resp.Cards.Yellow.M16_30,
			resp.Cards.Yellow.M31_45,
			resp.Cards.Yellow.M46_60,
			resp.Cards.Yellow.M61_75,
			resp.Cards.Yellow.M76_90,
			resp.Cards.Yellow.M91_105,
			resp.Cards.Yellow.M106_120,
		} {
			if minuteRange.Total != nil {
				yellow += *minuteRange.Total
			}
		}

		for _, minuteRange := range []model.MinuteCardStat{
			resp.Cards.Red.M0_15,
			resp.Cards.Red.M16_30,
			resp.Cards.Red.M31_45,
			resp.Cards.Red.M46_60,
			resp.Cards.Red.M61_75,
			resp.Cards.Red.M76_90,
			resp.Cards.Red.M91_105,
			resp.Cards.Red.M106_120,
		} {
			if minuteRange.Total != nil {
				red += *minuteRange.Total
			}
		}

		teamStats.YellowCardsTotal = &yellow
		teamStats.RedCardsTotal = &red

		for _, l := range resp.Lineup {
			lineup := &model.Lineup{
				Season: 	resp.League.Season,
				Formation: 	l.Formation,
				Played: 	l.Played,
			}
			s.db.Create(&lineup)
		}

		s.db.Create(&teamStats)
	}

	return stats, nil
}


func (s *service) SaveVenues() (*model.VenueResponse, error) {
	venues, err := s.client.FetchVenues()
	if err != nil {
		return nil, err
	}

	for _, v := range venues.Response {
		venue := &model.Venue{
			VenueID: 	v.VenueID,
			VenueName: 	v.VenueName,
			Address: 	v.Address,
			City: 		v.City,
			Capacity: 	v.Capacity,
			Surface: 	v.Surface,
		}
		s.db.Create(&venue)
	}

	return venues, nil
}


func (s *service) SaveStandings() ([]*model.StandingResponse, error) {
	standings, err := s.client.FetchStandings(39, 33)
	if err != nil {
		return nil, err
	}

	for _, response := range standings {
		for _, standingDTO := range response.Response {
			info := standingDTO.StandingInfo

			for _, group := range info.Standings {
				for _, entry := range group {
					record := model.Standing{
						LeagueID:    info.ID,
						LeagueName:  info.Name,
						Country:     info.Country,
						Season:      info.Season,

						TeamID:      entry.Team.ID,
						TeamName:    entry.Team.Name,

						Rank:        entry.Rank,
						Points:      entry.Points,
						GoalsDiff:   entry.GoalsDiff,
						GroupName:   entry.Group,
						Form:        entry.Form,
						Status:      entry.Status,
						Description: entry.Description,

						PlayedAll:       entry.All.Played,
						WinsAll:         entry.All.Win,
						DrawsAll:        entry.All.Draw,
						LosesAll:        entry.All.Lose,
						GoalsForAll:     entry.All.Goals.For,
						GoalsAgainstAll: entry.All.Goals.Against,

						PlayedHome:       entry.Home.Played,
						WinsHome:         entry.Home.Win,
						DrawsHome:        entry.Home.Draw,
						LosesHome:        entry.Home.Lose,
						GoalsForHome:     entry.Home.Goals.For,
						GoalsAgainstHome: entry.Home.Goals.Against,

						PlayedAway:       entry.Away.Played,
						WinsAway:         entry.Away.Win,
						DrawsAway:        entry.Away.Draw,
						LosesAway:        entry.Away.Lose,
						GoalsForAway:     entry.Away.Goals.For,
						GoalsAgainstAway: entry.Away.Goals.Against,

						UpdatedAt: entry.Update,
					}
					s.db.Create(&record)
				}
			}
		}
	}
	
	return standings, nil
}


func (s *service) SaveFixtures() ([]*model.FixtureResponse, error) {
	fixtures, err := s.client.FetchFixtures(39, 33) 
	if err != nil {
		return nil, err
	}

	var allFixtures []model.Fixture

	for _, seasonResp := range fixtures {
		for _, dto := range seasonResp.Response {
			fixture := dto.Fixture
			league 	:= dto.League
			teams 	:= dto.Teams
			goals 	:= dto.Goals
			score 	:= dto.Score

			allFixtures = append(allFixtures, model.Fixture{
				FixtureID:  	fixture.ID,
				Referee:    	fixture.Referee,
				Timezone:   	fixture.Timezone,
				Date:       	fixture.Date,
				Timestamp:  	fixture.Timestamp,
				PeriodFirst:  	fixture.Periods.First,
				PeriodSecond: 	fixture.Periods.Second,
				VenueID:     	fixture.Venue.ID,
				VenueName:   	fixture.Venue.Name,
				VenueCity:   	fixture.Venue.City,
				StatusLong:   	fixture.Status.Long,
				StatusShort:  	fixture.Status.Short,
				StatusElapsed: 	fixture.Status.Elapsed,
				StatusExtra:   	fixture.Status.Extra,

				LeagueID:   league.ID,
				LeagueName: league.Name,
				Country:    league.Country,
				Season:     league.Season,
				Round:      league.Round,
				Standings:  league.Standings,

				HomeTeamID:   teams.Home.ID,
				HomeTeamName: teams.Home.Name,
				HomeTeamLogo: teams.Home.Logo,
				HomeWinner:   teams.Home.Winner,
				AwayTeamID:   teams.Away.ID,
				AwayTeamName: teams.Away.Name,
				AwayTeamLogo: teams.Away.Logo,
				AwayWinner:   teams.Away.Winner,

				GoalsHome:  safeInt(goals.Home),
				GoalsAway:  safeInt(goals.Away),

				HalftimeHome:  score.Halftime.Home,
				HalftimeAway:  score.Halftime.Away,
				FulltimeHome:  score.Fulltime.Home,
				FulltimeAway:  score.Fulltime.Away,
				ExtratimeHome: score.Extratime.Home,
				ExtratimeAway: score.Extratime.Away,
				PenaltyHome:   score.Penalty.Home,
				PenaltyAway:   score.Penalty.Away,
			})
		}
	}

	s.db.Create(&allFixtures)

	return fixtures, nil
}


func (s *service) SaveInjuries() ([]*model.InjuryResponse, error) {
	injuriesResp, err := s.client.FetchInjuries(33)
	if err != nil {
		return nil, err
	}

	var injuries []model.Injury

	for _, seasonResp := range injuriesResp {
		for _, inj := range seasonResp.Response {
			injury := model.Injury{
				PlayerID:   		inj.Player.ID,
				PlayerName: 		inj.Player.Name,
				PlayerPhoto: 		inj.Player.Photo,
				Type:       		inj.Player.Type,
				Reason:     		inj.Player.Reason,

				TeamID:   			inj.Team.ID,
				TeamName: 			inj.Team.Name,
				TeamLogo: 			inj.Team.Logo,

				FixtureID:       	inj.Fixture.ID,
				FixtureDate:     	inj.Fixture.Date,
				FixtureTimestamp: 	inj.Fixture.Timestamp,
				FixtureTimezone: 	inj.Fixture.Timezone,

				LeagueID:   		inj.League.ID,
				LeagueName: 		inj.League.Name,
				Country:    		inj.League.Country,
				Season:     		inj.League.Season,
				LeagueLogo: 		inj.League.Logo,
				Flag:       		inj.League.Flag,
			}

			injuries = append(injuries, injury)
		}
	}

	s.db.Create(&injuries)

	return injuriesResp, nil
}


func (s *service) SaveSquad() (*model.SquadResponse, error) {
	squad, err := s.client.FetchSquad(33)
	if err != nil {
		return nil, err
	}

	var squadEntries []model.Squad

	for _, dto := range squad.Response {
		for _, player := range dto.Players {
			entry := &model.Squad{
				TeamID:       dto.Team.ID,
				TeamName:     dto.Team.Name,
				TeamLogo:     dto.Team.Logo,
				PlayerID:     player.ID,
				PlayerName:   player.Name,
				Age:          player.Age,
				Number:       player.Number,
				Position:     player.Position,
				PlayerPhoto:  player.Photo,
			}
			squadEntries = append(squadEntries, *entry)
		}
	}

	s.db.Create(&squadEntries)

	return squad, nil
}