package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Akdkkras/rpg-game/internal/entity"
	"github.com/Akdkkras/rpg-game/internal/infrastructure/loader"
	"github.com/Akdkkras/rpg-game/internal/repository/memory"
)

type Game struct {
	Locations  []entity.Location
	StartAlias string
}

func startQuest(quest entity.Quest) bool {
	// TODO: реализовать
	return false
}

func main() {

	// TODO: Обдумать, где и как определять начальную локацию
	currentLocationAlias := "tavern"
	locations, err := loader.LoadLocationsFromJSON("tmp")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	locationRepository := memory.NewMemoryLocationRepository(locations)
	_ = locationRepository

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== ДОБРО ПОЖАЛОВАТЬ В ТЕКСТОВУЮ RPG ИГРУ ===\n")

	for {
		// Получаем текущую локацию
		currentLocation, ok := locations[currentLocationAlias]
		if !ok {
			if currentLocationAlias == "finish" {
				fmt.Println("🎉 ПОЗДРАВЛЯЕМ! Вы прошли игру! 🎉")
				break
			}

			// TODO: можно избежать подобных проверок, провалидировов данные на все требования в одной функции перед в входом в движок (хз есть ли в этом смысл)
			log.Fatal("Ошибка: стартовая локация не найдена")
		}

		// Выводим локацию в едином формате
		fmt.Println("========================================")
		fmt.Printf("📍 ЛОКАЦИЯ: %s\n", currentLocation.Title)
		fmt.Println("========================================")
		fmt.Printf("%s\n\n", currentLocation.Description)

		// Выводим квесты
		fmt.Println("Доступные квесты:")
		for _, quest := range currentLocation.Quests {
			fmt.Printf("  - %s\n", quest.Title)
		}
		fmt.Println()

		// Ожидаем команду от пользователя
		fmt.Println("Доступные команды:")
		fmt.Println("  1. следующая локация")
		fmt.Println("  2. выполнить квест")
		fmt.Println("  3. завершить игру")
		fmt.Print("\n> ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		// Обрабатываем команды
		switch command {
		case "1":
			// Переходим к следующей локации
			currentLocationAlias = currentLocation.NextLocation
			fmt.Println("\n✨ Вы отправляетесь дальше... ✨\n")

		case "2":
			// Переходим к выбору квеста
			fmt.Println("\n⚔️ Выберите квест для прохождения: ⚔️\n")
			for idx, quest := range currentLocation.Quests {
				fmt.Printf("  %d. %s\n", idx+1, quest.Title)
			}
			fmt.Print("\n> ")

			questNumberString, _ := reader.ReadString('\n')
			questNumberString = strings.TrimSpace(questNumberString)
			questNumber, err := strconv.Atoi(questNumberString)
			if err != nil {
				fmt.Println("❌ Некорректный ввод. Попробуйте снова.")
				break
			}

			questIdx := questNumber - 1
			if questIdx < 0 || questIdx > len(currentLocation.Quests) {
				fmt.Println("❌ Некорректный ввод. Попробуйте снова.")
				break
			}

			startQuest(currentLocation.Quests[questIdx])

		case "3":
			fmt.Println("\n👋 Спасибо за игру! До свидания!")
			return

		default:
			fmt.Println("\n❌ Неизвестная команда. Попробуйте снова.\n")
		}
	}
}
