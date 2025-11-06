package service

import (
	"github.com/deikioveca/TheRedDevilsData/api/football_client"
	"gorm.io/gorm"
)

type Service interface {
	DataImporter
	DataProvider
}

type service struct {
	db 		*gorm.DB
	client 	football_client.FootballClient
}


func NewService(db *gorm.DB, client football_client.FootballClient) Service {
	return &service{db: db, client: client}
}