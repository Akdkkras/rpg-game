package memory

import (
	"errors"

	"github.com/Akdkkras/rpg-game/internal/entity"
	"github.com/Akdkkras/rpg-game/internal/repository"
)

type locationRepository struct {
	store map[string]*entity.Location
}

func NewMemoryLocationRepository(store map[string]*entity.Location) repository.LocationRepository {
	return &locationRepository{
		store: store,
	}
}

func (r *locationRepository) GetLocationByAlias(alias string) (*entity.Location, error) {
	location, ok := r.store[alias]
	if !ok {
		return nil, errors.New("локация не найдена")
	}
	return location, nil
}
