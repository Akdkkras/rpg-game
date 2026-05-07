package main

import (
	"log"

	"github.com/Akdkkras/rpg-game/internal/adapter/cli"
	"github.com/Akdkkras/rpg-game/internal/entity"
	"github.com/Akdkkras/rpg-game/internal/infrastructure/loader"
	"github.com/Akdkkras/rpg-game/internal/repository/memory"
	"github.com/Akdkkras/rpg-game/internal/usecase"
)

func main() {
	locations, err := loader.LoadLocationsFromJSON("tmp")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// TODO: Обдумать, мб можно определять начальную локу напрямую из json
	worldState := entity.NewWorldState("tavern")

	locationRepository := memory.NewMemoryLocationRepository(locations)
	ui := cli.NewTerminalUI()
	worldUseCase := usecase.NewWorldUseCase(worldState, locationRepository, ui)

	worldUseCase.Run()
}
