package services

import (
	"context"
	"idm/inner/models"
	"idm/inner/repositories"
)

type FilterService struct {
	repo *repositories.FilterRepository
}

func NewFilterService(repo *repositories.FilterRepository) *FilterService {
	return &FilterService{repo: repo}
}

func (s *FilterService) GetAll(context context.Context) ([]models.Filter, error) {
	return s.repo.FindAll(context)
}

func (s *FilterService) Add(context context.Context, filter models.Filter) (int64, error) {
	return s.repo.Add(context, filter)
}

func (s *FilterService) FindById(context context.Context, filterId int64) (*models.Filter, error) {
	return s.repo.FindByID(context, filterId)
}
