package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Location struct {
	Alias        string
	Title        string
	Description  string
	NextLocation string
}

type Game struct {
	Locations  []Location
	StartAlias string
}

func main() {
	game := Game{
		Locations: []Location{
			{
				Alias:        "tavern",
				Title:        "Таверна",
				Description:  "Тёплое и шумное место с запахом эля и жареного мяса. В углу играют в кости, а за стойкой хозяин протирает кружки.",
				NextLocation: "dark_forest",
			},
			{
				Alias:        "dark_forest",
				Title:        "Тёмный лес",
				Description:  "Густые деревья смыкаются над головой, почти не пропуская солнечный свет. Слышен хруст веток и далёкий волчий вой.",
				NextLocation: "abandoned_castle",
			},
			{
				Alias:        "abandoned_castle",
				Title:        "Заброшенный замок",
				Description:  "Старые каменные стены покрыты мхом и плющом. Внутри царит зловещая тишина, нарушаемая лишь скрипом половиц.",
				NextLocation: "town_square",
			},
			{
				Alias:        "town_square",
				Title:        "Городская площадь",
				Description:  "Оживлённая площадь с фонтаном в центре. Торговцы зазывают покупателей, дети бегают между взрослыми, а на ступенях ратуши выступает менестрель.",
				NextLocation: "finish",
			},
		},
		// TODO: стоит додумать, откуда получать стартовую локацию
		StartAlias: "tavern",
	}

	locations := make(map[string]Location)
	for _, location := range game.Locations {
		if _, ok := locations[location.Alias]; ok {
			// TODO: можно избежать подобных проверок, провалидировов данные на все требования в одной функции перед в входом в движок (хз есть ли в этом смысл)
			log.Fatal("Ошибка: Location.Alias обязан быть уникальным значением")
		}
		locations[location.Alias] = location
	}

	currentLocationAlias := game.StartAlias
	if currentLocationAlias == "" {
		// TODO: можно избежать подобных проверок, провалидировов данные на все требования в одной функции перед в входом в движок (хз есть ли в этом смысл)
		log.Fatal("Ошибка: стартовая локация не задана")
	}
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

		// Ожидаем команду от пользователя
		fmt.Println("Доступные команды:")
		fmt.Println("  - следующая локация (или 'дальше', 'вперед')")
		fmt.Println("  - завершить игру (или 'выход', 'стоп')")
		fmt.Print("\n> ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		// Обрабатываем команды
		switch command {
		case "следующая локация", "дальше", "вперед":
			// Переходим к следующей локации
			currentLocationAlias = currentLocation.NextLocation
			fmt.Println("\n✨ Вы отправляетесь дальше... ✨\n")

		case "завершить игру", "выход", "стоп":
			fmt.Println("\n👋 Спасибо за игру! До свидания!")
			return

		default:
			fmt.Println("\n❌ Неизвестная команда. Попробуйте снова.\n")
		}
	}
}
