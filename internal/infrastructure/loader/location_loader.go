package loader

import "github.com/Akdkkras/rpg-game/internal/entity"

func LoadLocationsFromJSON(filePath string) (map[string]*entity.Location, error) {
	// TODO: Заменить на JSON парсинг
	store := map[string]*entity.Location{
		"tavern": {
			Alias:        "tavern",
			Title:        "Таверна",
			Description:  "Тёплое и шумное место с запахом эля и жареного мяса. В углу играют в кости, а за стойкой хозяин протирает кружки.",
			NextLocation: "dark_forest",
			Quests: []entity.Quest{
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
		"dark_forest": {
			Alias:        "dark_forest",
			Title:        "Тёмный лес",
			Description:  "Густые деревья смыкаются над головой, почти не пропуская солнечный свет. Слышен хруст веток и далёкий волчий вой.",
			NextLocation: "abandoned_castle",
			Quests: []entity.Quest{
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
		"abandoned_castle": {
			Alias:        "abandoned_castle",
			Title:        "Заброшенный замок",
			Description:  "Старые каменные стены покрыты мхом и плющом. Внутри царит зловещая тишина, нарушаемая лишь скрипом половиц.",
			NextLocation: "town_square",
			Quests: []entity.Quest{
				{
					Title:       "Призрачный рыцарь",
					Description: "Упокоить дух древнего рыцаря, который до сих пор охраняет пустые залы.",
				},
			},
		},
		"town_square": {
			Alias:        "town_square",
			Title:        "Городская площадь",
			Description:  "Оживлённая площадь с фонтаном в центре. Торговцы зазывают покупателей, дети бегают между взрослыми, а на ступенях ратуши выступает менестрель.",
			NextLocation: "finish",
			Quests: []entity.Quest{
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
	}

	return store, nil
}
