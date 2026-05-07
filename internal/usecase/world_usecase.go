package usecase

import (
	"log"

	"github.com/Akdkkras/rpg-game/internal/adapter/cli"
	"github.com/Akdkkras/rpg-game/internal/entity"
	"github.com/Akdkkras/rpg-game/internal/repository"
)

type WorldUseCase struct {
	locationRepo repository.LocationRepository
	worldState   *entity.WorldState
	ui           *cli.TerminalUI // TODO: заменить конкретную реализацию на порт (сначала добавить порт)
}

func NewWorldUseCase(worldState *entity.WorldState, locationRepo repository.LocationRepository, ui *cli.TerminalUI) *WorldUseCase {
	return &WorldUseCase{
		locationRepo: locationRepo,
		worldState:   worldState,
		ui:           ui,
	}
}

// TODO: Удалить отсюда
func startQuest(quest entity.Quest) bool {
	// TODO: реализовать
	return false
}

func (uc *WorldUseCase) Run() {

	uc.ui.PrintGameStart()

	for {
		// Получаем текущую локацию
		currentLocation, err := uc.locationRepo.GetLocationByAlias(uc.worldState.CurrentLocationAlias)
		if err != nil {
			if uc.worldState.CurrentLocationAlias == "finish" {
				uc.ui.PrintGameEnd()
				break
			}

			// TODO: можно избежать подобных проверок, провалидировов данные на все требования в одной функции перед в входом в движок (хз есть ли в этом смысл)
			log.Fatal("Ошибка: стартовая локация не найдена")
		}

		// Выводим локацию в едином формате
		uc.ui.PrintLocationInfo(currentLocation)

		// Ожидаем команду от пользователя
		command := uc.ui.ChooseFromSuggested()

		// Обрабатываем команды
		switch command {
		case "1":
			// Переходим к следующей локации
			uc.worldState.CurrentLocationAlias = currentLocation.NextLocationAlias
			uc.ui.PrintMoveOn()

		case "2":
			// Переходим к выбору квеста
			questIdx, err := uc.ui.ChooseQuestFromLocation(currentLocation)
			if err != nil {
				uc.ui.PrintWrongChoice()
				break
			}
			startQuest(currentLocation.Quests[questIdx])

		case "3":
			uc.ui.PrintGameExit()
			return

		default:
			uc.ui.PrintWrongCommand()
		}
	}
}
