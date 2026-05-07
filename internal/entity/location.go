package entity

type Location struct {
	// TODO: Возможно, алиас не нужен на уровне сущности, а только на уровне файла конфигураций игры (пока что исп-ся только в кач-ве ключа в мапах на эту стр-ру)
	Alias             string `json:"alias"` // Требования: уникальны
	Title             string `json:"title"`
	Description       string `json:"description"`
	NextLocationAlias string `json:"nextLocation"`
	Quests            []Quest
}
