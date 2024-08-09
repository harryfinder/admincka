package usecase

import "github.com/activ-capital/partner-service/internal/entity"

type Usecase interface {
}

type usecase struct {
	entity entity.Entity
}

func New(entity entity.Entity) Usecase {
	return &usecase{
		entity: entity,
	}
}
