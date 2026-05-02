package memory

import (
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
	//TODO implement me
	panic("implement me")
}
