package services

import "diplomGo/internal/repository"

type DatabaseService struct {
	db repository.Database
}

func newDatabaseService(db repository.Database) *DatabaseService {
	return &DatabaseService{db}
}
