package talosecs

import "reflect"

var isInitialized bool

// Init initializes all of your ECS updateSystems. Call it once on game world start before calling Update
func Init() {
	if isInitialized {
		panic("ECS already initialized!")
	}

	doInitSystems()
	isInitialized = true
}

// Update calls Update of each UpdateSystem every frame. Used to handle most of game logic.
func Update() {
	doUpdateSystems()
	clearOneFrames()
	signals = nil
}

func Reset() {
	componentsPool = map[reflect.Type][]any{}
	entsComponents = map[Entity][]any{}
	componentsEnts = map[any]Entity{}
	signals = nil
	oneFrames = nil
	updateSystems = nil
	initSystems = nil

	currentEntityId = 0
	isInitialized = false
}
