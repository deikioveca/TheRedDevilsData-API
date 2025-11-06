package football_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/deikioveca/TheRedDevilsData/api/model"
	"gorm.io/gorm"
)


type FootballClient interface {
	FetchCountries() (*model.CountryResponse, error)
	FetchAllLeaguesForTeam() (*model.LeagueResponse, error)
	FetchTeam() (*model.TeamResponse, error)
	FetchTeamStats(teamID, leagueID int) ([]*model.TeamStatsResponse, error)
	FetchVenues() (*model.VenueResponse, error)
	FetchStandings(leagueID, teamID int) ([]*model.StandingResponse, error)
	FetchFixtures(leagueID, teamID int) ([]*model.FixtureResponse, error)
	FetchInjuries(teamID int) ([]*model.InjuryResponse, error)
	FetchSquad(teamID int) (*model.SquadResponse, error)
}


type footballClient struct {
	httpClient 	*http.Client
	apiKey		string
	baseUrl		string
}


func NewFootballClient(db *gorm.DB) FootballClient {
	apiKey 	:= os.Getenv("API_KEY")
	baseUrl := os.Getenv("BASE_URL")

	return &footballClient{
		httpClient: 	&http.Client{
			Timeout: 	15 * time.Second,
		},
		apiKey: 	apiKey,
		baseUrl: 	baseUrl,
	}
}


func (f *footballClient) get(endpoint string, data any) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", f.baseUrl, endpoint), nil)
	if err != nil {
		return err
	}

	req.Header.Add("x-rapidapi-key", f.apiKey)
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := f.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed request to %s: %s", endpoint, res.Status)
	}

	return json.NewDecoder(res.Body).Decode(&data)
}



func (f *footballClient) FetchCountries() (*model.CountryResponse, error) {
	var data model.CountryResponse
	if err := f.get("/countries", &data); err != nil {
		return nil, err
	}
	return &data, nil
}


func (f *footballClient) FetchAllLeaguesForTeam() (*model.LeagueResponse, error) {
	var data model.LeagueResponse
	if err := f.get("/leagues?team=33", &data); err != nil {
		return nil, err
	}
	return &data, nil
}


func (f *footballClient) FetchTeam() (*model.TeamResponse, error) {
	var data model.TeamResponse
	if err := f.get("/teams?id=33", &data); err != nil {
		return nil, err
	}
	return &data, nil
}


func (f *footballClient) FetchTeamStats(teamID, leagueID int) ([]*model.TeamStatsResponse, error) {
	seasons := []int{2021, 2022, 2023}
	results := make([]*model.TeamStatsResponse, len(seasons))

	var wg sync.WaitGroup
	errCh := make(chan error, len(seasons))

	for i, season := range seasons {
		wg.Add(1)

		go func(i, season int) {
			defer wg.Done()

			var data model.TeamStatsResponse
			endpoint := fmt.Sprintf("/teams/statistics?league=%d&team=%d&season=%d", leagueID, teamID, season)
			if err := f.get(endpoint, &data); err != nil {
				errCh <- fmt.Errorf("season %d: %w", season, err)
				return 
			}
			results[i] = &data
		}(i, season)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		return nil, <- errCh
	}

	return results, nil
}


func (f *footballClient) FetchVenues() (*model.VenueResponse, error) {
	var data model.VenueResponse
	if err := f.get("/venues?country=england", &data); err != nil {
		return nil, err
	}
	return &data, nil
}


func (f *footballClient) FetchStandings(leagueID, teamID int) ([]*model.StandingResponse, error) {
	seasons := []int{2021, 2022, 2023}
	results := make([]*model.StandingResponse, len(seasons))

	var wg sync.WaitGroup
	errCh := make(chan error, len(seasons))

	for i, season := range seasons {
		wg.Add(1)

		go func(i, season int) {
			defer wg.Done()

			var data model.StandingResponse
			endpoint := fmt.Sprintf("/standings?league=%d&season=%d&team=%d", leagueID, season, teamID)
			if err := f.get(endpoint, &data); err != nil {
				errCh <- fmt.Errorf("season %d: %w", season, err)
				return 
			}
			results[i] = &data
		}(i, season)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		return nil, <- errCh
	}

	return results, nil
}


func (f *footballClient) FetchFixtures(leagueID, teamID int) ([]*model.FixtureResponse, error) {
	seasons := []int{2021, 2022, 2023}
	results := make([]*model.FixtureResponse, len(seasons))

	var wg sync.WaitGroup
	errCh := make(chan error, len(seasons))

	for i, season := range seasons {
		wg.Add(1)

		go func(i, season int) {
			defer wg.Done()

			var data model.FixtureResponse
			endpoint := fmt.Sprintf("/fixtures?league=%d&season=%d&team=%d", leagueID, season, teamID)
			if err := f.get(endpoint, &data); err != nil {
				errCh <- fmt.Errorf("season %d: %w", season, err)
				return 
			}
			results[i] = &data
		}(i, season)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		return nil, <- errCh
	}

	return results, nil
}


func (f *footballClient) FetchInjuries(teamID int) ([]*model.InjuryResponse, error) {
	seasons := []int{2021, 2022, 2023}
	results := make([]*model.InjuryResponse, len(seasons))

	var wg sync.WaitGroup
	errCh := make(chan error, len(seasons))

	for i, season := range seasons {
		wg.Add(1)

		go func(i, season int) {
			defer wg.Done()

			var data model.InjuryResponse
			endpoint := fmt.Sprintf("/injuries?season=%d&team=%d", season, teamID)
			if err := f.get(endpoint, &data); err != nil {
				errCh <- fmt.Errorf("season %d: %w", season, err)
				return 
			}
			results[i] = &data
		}(i, season)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		return nil, <- errCh
	}

	return results, nil
}


func (f *footballClient) FetchSquad(teamID int) (*model.SquadResponse, error) {
	var data model.SquadResponse
	endpoint := fmt.Sprintf("/players/squads?team=%d", teamID)
	if err := f.get(endpoint, &data); err != nil {
		return nil, err
	}

	return &data, nil
}