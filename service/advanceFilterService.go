package service

import (
	"psi/model"
	"psi/repository"
)

type advanceFilterService struct {
	repo repository.AdvanceFilterRepo
}

type AdvanceFilterService interface {
	AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error)
}

func (a *advanceFilterService) AdvanceFilter(payload model.AdvanceFilterPayload) (interface{}, int64, error) {
	data, total, err := a.repo.AdvanceFilter(payload)

	return data, total, err
}

func NewAdvanceFilterService() AdvanceFilterService {
	repo := repository.NewAdvanceFilterRepo()
	return &advanceFilterService{
		repo: repo,
	}
}
