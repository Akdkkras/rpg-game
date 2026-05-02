// TODO: Поскольку Quest не может быть независимой сущ-тью (Location и Quest - агрегат), возможно, не имеет смысла выносить Quest в отдельный файл

package entity

type Quest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// Alias    string // Пока что не нужен, но, возможно, пригодится в дальнейшем для связи с конкретным квестом
	// TODO: По шагам внутри квеста будет запускаться новый, внутренний движок (как сейчас по локациям)
	// QuestSteps 	[]QuestStep
}
