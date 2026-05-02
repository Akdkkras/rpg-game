package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	// TODO: Возможно, алиас не нужен на уровне сущности, а только на уровне файла конфигураций игры (пока что исп-ся только в кач-ве ключа в мапах на эту стр-ру)
	Alias        string // Требования: уникальны
	Title        string
	Description  string
	NextLocation string
	Quests       []Quest
}

type Quest struct {
	Title       string
	Description string
	// Alias    string // Пока что не нужен, но, возможно, пригодится в дальнейшем для связи с конкретным квестом
	// TODO: По шагам внутри квеста будет запускаться новый, внутренний движок (как сейчас по локациям)
	// QuestSteps 	[]QuestStep
}

type Game struct {
	Locations  []Location
	StartAlias string
}

func startQuest(quest Quest) bool {
	// TODO: реализовать
	return false
}

func main() {
	game := Game{
		Locations: []Location{
			{
				Alias:        "tavern",
				Title:        "Таверна",
				Description:  "Тёплое и шумное место с запахом эля и жареного мяса. В углу играют в кости, а за стойкой хозяин протирает кружки.",
				NextLocation: "dark_forest",
				Quests: []Quest{
					{
						Title:       "Крысы в подвале",
						Description: "Хозяин таверны просит избавиться от крыс в подвале.",
					},
					{
						Title:       "Потерянное кольцо",
						Description: "Посетитель обронил золотое кольцо где-то у барной стойки. Нужно его найти.",
					},
					{
						Title:       "Пьяная драка",
						Description: "Разнимите двух завсегдатаев, пока они не разнесли всю мебель.",
					},
				},
			},
			{
				Alias:        "dark_forest",
				Title:        "Тёмный лес",
				Description:  "Густые деревья смыкаются над головой, почти не пропуская солнечный свет. Слышен хруст веток и далёкий волчий вой.",
				NextLocation: "abandoned_castle",
				Quests: []Quest{
					{
						Title:       "Сбор грибов",
						Description: "Собрать редкие светящиеся грибы для местного алхимика.",
					},
					{
						Title:       "Волчья стая",
						Description: "Прогнать стаю свирепых волков, перегородивших тропу.",
					},
				},
			},
			{
				Alias:        "abandoned_castle",
				Title:        "Заброшенный замок",
				Description:  "Старые каменные стены покрыты мхом и плющом. Внутри царит зловещая тишина, нарушаемая лишь скрипом половиц.",
				NextLocation: "town_square",
				Quests: []Quest{
					{
						Title:       "Призрачный рыцарь",
						Description: "Упокоить дух древнего рыцаря, который до сих пор охраняет пустые залы.",
					},
				},
			},
			{
				Alias:        "town_square",
				Title:        "Городская площадь",
				Description:  "Оживлённая площадь с фонтаном в центре. Торговцы зазывают покупателей, дети бегают между взрослыми, а на ступенях ратуши выступает менестрель.",
				NextLocation: "finish",
				Quests: []Quest{
					{
						Title:       "Украденное яблоко",
						Description: "Поймать мальчишку, который стащил яблоко с прилавка торговца.",
					},
					{
						Title:       "Песня менестреля",
						Description: "Помочь менестрелю вспомнить забытые слова для его новой баллады.",
					},
				},
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
