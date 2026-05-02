package repository

import "github.com/Akdkkras/rpg-game/internal/entity"

type LocationRepository interface {
	GetLocationByAlias(alias string) (*entity.Location, error)
}
