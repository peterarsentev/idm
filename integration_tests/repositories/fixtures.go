package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"idm/inner/models"
	"idm/inner/repositories"
)

func clearDb(db *sqlx.DB) error {
	_, err := db.ExecContext(context.Background(), `TRUNCATE TABLE public.dionea_filter CASCADE`)
	return err
}

type FixtureFilter struct {
	filterRepository *repositories.FilterRepository
}

func NewFixtureFilter(repository *repositories.FilterRepository) *FixtureFilter {
	return &FixtureFilter{
		filterRepository: repository,
	}
}

func (f *FixtureFilter) NewFilter(name string) (int64, error) {
	return f.filterRepository.Add(context.Background(), models.Filter{Name: name})
}

type FixtureKey struct {
	keyRepository *repositories.KeyRepository
}

func NewKeyRepository(repository *repositories.KeyRepository) *FixtureKey {
	return &FixtureKey{
		keyRepository: repository,
	}
}

func (f *FixtureKey) NewKey(filterID int64, name string) (int64, error) {
	key := models.Key{
		FilterID: filterID,
		Name:     name,
	}
	return f.keyRepository.Add(context.Background(), key)
}
