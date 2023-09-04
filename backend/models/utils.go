package models

import "math/rand"

const (
	StateNum    = 3
	RSSideBoth  = "both"
	RSSideLeft  = "left"
	RSSideRight = "right"
)

type (
	State struct {
		Name string
		Code uint
	}

	RelationState struct {
		data  [StateNum]State
		names map[string]int
		codes map[uint]int
	}
)

func buildStrongSideStruct() RelationState {
	var strongSide RelationState

	data := [StateNum]State{
		{Name: RSSideBoth, Code: 1},
		{Name: RSSideLeft, Code: 2},
		{Name: RSSideRight, Code: 3},
	}

	strongSide.data = data
	strongSide.names = make(map[string]int)
	strongSide.codes = make(map[uint]int)

	for index, state := range strongSide.data {
		strongSide.names[state.Name] = index
		strongSide.codes[state.Code] = index
	}

	return strongSide
}

var StrongSide = buildStrongSideStruct()

func (r *RelationState) Equal(name string, code uint) bool {
	mapCode, okCode := r.codes[code]
	mapName, okName := r.names[name]

	if !okCode || !okName {
		return false
	}

	return mapCode != mapName
}

func (r *RelationState) GetName(code uint) *string {
	mapCode, ok := r.codes[code]

	if !ok {
		return nil
	}

	name := r.data[mapCode].Name

	return &name
}

func (r *RelationState) GetCode(name string) *uint {
	mapName, ok := r.names[name]

	if !ok {
		return nil
	}

	code := r.data[mapName].Code

	return &code
}

func (r *RelationState) RandState() State {
	randIndex := rand.Intn(len(r.data))

	state := r.data[randIndex]

	return state
}
