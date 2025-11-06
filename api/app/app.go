package app

import (
	"log"
	"net/http"

	"github.com/deikioveca/TheRedDevilsData/api/database"
	"github.com/deikioveca/TheRedDevilsData/api/football_client"
	"github.com/deikioveca/TheRedDevilsData/api/handler"
	"github.com/deikioveca/TheRedDevilsData/api/service"
	"gorm.io/gorm"
)


type App struct {
	DB 			*gorm.DB
	Client		football_client.FootballClient
	Service		service.Service
	Handler		*handler.Handler
}


func NewApp() *App {
	db 		:= database.InitDB()
	client 	:= football_client.NewFootballClient(db)
	service := service.NewService(db, client)
	handler := handler.NewHandler(service)

	return &App{
		DB: 		db,
		Client: 	client,
		Service: 	service,
		Handler: 	handler,
	}
}


func (a *App) Run(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /country", 			a.Handler.GetCountries)
	mux.HandleFunc("GET /country/{name}", 	a.Handler.GetCountryByName)

	mux.HandleFunc("GET /league", a.Handler.GetLeagues)

	mux.HandleFunc("GET /team", a.Handler.GetTeam)

	mux.HandleFunc("GET /teamStats/games/{season}", 		a.Handler.GetTeamStatsGames)
	mux.HandleFunc("GET /teamStats/goals/{season}", 		a.Handler.GetTeamStatsGoals)
	mux.HandleFunc("GET /teamStats/streak/{season}", 		a.Handler.GetTeamStatsStreak)
	mux.HandleFunc("GET /teamStats/biggest/{season}", 		a.Handler.GetTeamStatsBiggest)
	mux.HandleFunc("GET /teamStats/cleansheet/{season}", 	a.Handler.GetTeamStatsCleanSheet)
	mux.HandleFunc("GET /teamStats/failedtoscore/{season}", a.Handler.GetTeamStatsFailedToScore)
	mux.HandleFunc("GET /teamStats/penalty/{season}", 		a.Handler.GetTeamStatsPenalty)
	mux.HandleFunc("GET /teamStats/cards/{season}", 		a.Handler.GetTeamStatsCards)
	mux.HandleFunc("GET /teamStats/lineup/{season}", 		a.Handler.GetTeamStatsLineups)

	mux.HandleFunc("GET /venue", 					a.Handler.GetVenues)
	mux.HandleFunc("GET /venue/{city}", 			a.Handler.GetVenuesByCity)
	mux.HandleFunc("GET /venue/biggest&smallest", 	a.Handler.GetVenuesBiggestAndSmallest)

	mux.HandleFunc("GET /standings/{season}", a.Handler.GetStandingsBySeason)

	mux.HandleFunc("GET /fixtures/{season}", a.Handler.GetFixturesBySeason)

	mux.HandleFunc("GET /injuries/{season}", a.Handler.GetInjuriesBySeason)

	mux.HandleFunc("GET /squad", a.Handler.GetSquad)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}