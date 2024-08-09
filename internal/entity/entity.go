package entity

import "github.com/activ-capital/partner-service/internal/database"

type Entity interface {
}

type entity struct {
	database database.Database
}

func New(database database.Database) Entity {
	return &entity{database: database}
}
