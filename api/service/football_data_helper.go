package service

import (
	"errors"

	"github.com/deikioveca/TheRedDevilsData/api/model"
	"gorm.io/gorm"
)

var (
	ErrTeamStatsNotFound = errors.New("team stats for this season not found")

	ErrLineupNotFound = errors.New("lineups for this season not found")
)


func safeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}


func safeString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}


func safeBool(ptr *bool) bool {
	if ptr != nil {
		return *ptr
	}
	return false
}


func (s *service) getTeamStatsBySeason(season int) (*model.TeamStats, *model.ManchesterUnitedTeamStatsDTO, error) {
	var teamStats model.TeamStats
	if err := s.db.Where("season = ?", season).First(&teamStats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, ErrTeamStatsNotFound
		}
		return nil, nil, err
	}

	manchesterUnitedTeamStatsDTO := &model.ManchesterUnitedTeamStatsDTO{
		Name: 		teamStats.TeamName,
		League: 	teamStats.LeagueName,
		Season: 	teamStats.Season,
		Form: 		teamStats.Form,
	}

	return &teamStats, manchesterUnitedTeamStatsDTO, nil
}


func (s *service) getLineupsBySeason(season int) ([]model.Lineup, error) {
	var lineup []model.Lineup
	if err := s.db.Where("season = ?", season).Find(&lineup).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLineupNotFound
		}
		return nil, err
	}

	return lineup, nil
}