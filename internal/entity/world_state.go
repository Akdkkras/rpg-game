package entity

type WorldState struct {
	CurrentLocationAlias string
}

func NewWorldState(currentLocation string) *WorldState {
	return &WorldState{
		CurrentLocationAlias: currentLocation,
	}
}
