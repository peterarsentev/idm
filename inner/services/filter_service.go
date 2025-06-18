package services

import (
	"context"
	"idm/inner/models"
)

type FilterRepository interface {
	FindAll(ctx context.Context) ([]models.Filter, error)
	Add(ctx context.Context, filter models.Filter) (int64, error)
	FindByID(ctx context.Context, id int64) (models.Filter, error)
}

type KeyRepository interface {
	FindAll(ctx context.Context) ([]models.Key, error)
	FindKeyByFilterID(id int64) ([]models.Key, error)
}

type KeyValueRepository interface {
	FindAll(ctx context.Context) ([]models.KeyValue, error)
	FindValueInKeys(ids []int64) ([]models.KeyValue, error)
}

type FilterService struct {
	filterRepository FilterRepository
	keyRepository    KeyRepository
	kvalueRepository KeyValueRepository
}

func NewFilterService(repoFilter FilterRepository,
	repoKey KeyRepository,
	repoKValue KeyValueRepository) *FilterService {
	return &FilterService{
		filterRepository: repoFilter,
		keyRepository:    repoKey,
		kvalueRepository: repoKValue,
	}
}

func (s *FilterService) GetAll(context context.Context) ([]models.Filter, error) {
	return s.filterRepository.FindAll(context)
}

func (s *FilterService) Add(context context.Context, filter models.Filter) (int64, error) {
	return s.filterRepository.Add(context, filter)
}

func (s *FilterService) FindById(context context.Context, filterId int64) (models.Filter, error) {
	return s.filterRepository.FindByID(context, filterId)
}

func (s *FilterService) FindAll(context context.Context) ([]models.Filter, error) {
	return s.filterRepository.FindAll(context)
}

func (s *FilterService) FindKeyByFilterID(filterID int64) ([]models.Key, error) {
	return s.keyRepository.FindKeyByFilterID(filterID)
}

func (s *FilterService) FindVKeyInKeys(IDS []int64) (map[int64]models.KeyValue, error) {
	rsl := make(map[int64]models.KeyValue)
	values, err := s.kvalueRepository.FindValueInKeys(IDS)
	if err != nil {
		return nil, err
	}
	for _, value := range values {
		rsl[value.KeyID] = value
	}
	return rsl, nil
}
