package talosecs

import "reflect"

var currentEntityId Entity
var entsComponents = map[Entity][]any{}       // for filters
var componentsEnts = map[any]Entity{}         // for GetEntity
var componentsPool = map[reflect.Type][]any{} // for GetComponent + filters too
var oneFrames []any
var systems []System
var signals []any
var isInitialized bool

// NewEntity creates new Entity in the game world.
func NewEntity() Entity {
	currentEntityId++
	id := currentEntityId
	entsComponents[id] = []any{}

	return id
}

func GetEntity(comp any) Entity { return componentsEnts[comp] }
func IsAlive(entity Entity) bool {
	_, exist := entsComponents[entity]
	return exist
}

func SameEntity(a, b any) bool { return GetEntity(a) == GetEntity(b) }

func KillEntity(entity Entity) {
	if IsAlive(entity) {
		for _, component := range entsComponents[entity] {
			delete(componentsEnts, component)
		}

		delete(entsComponents, entity)
	}
}

// AddComponent adds any Component to the specified Entity. Component is a simple data struct.
func AddComponent(entity Entity, comp any) {
	typeOf := reflect.TypeOf(comp)
	var componentsSlice []any // todo use getComponentsSlice()

	if foundSlice, ok := componentsPool[typeOf]; ok {
		componentsSlice = foundSlice
	}

	componentsSlice = append(componentsSlice, comp)

	componentsPool[typeOf] = componentsSlice
	componentsEnts[comp] = entity
	entsComponents[entity] = append(entsComponents[entity], comp)
}

// AddOneFrame is same to AddComponent, but components added with this function live only one frame.
func AddOneFrame(entity Entity, comp any) {
	oneFrames = append(oneFrames, comp)
	AddComponent(entity, comp)
}

// DelComponent removes component of type T of specified entity. It will be not caught in next systems and next GetComponent (even in the same frame).
func DelComponent[T any](entity Entity) {
	comp, isExist := GetComponent[T](entity)

	if isExist {
		DelConcreteComponent(comp, entity)
	}
}

func DelConcreteComponent(comp any, entity Entity) {
	delete(componentsEnts, comp)

	typeOf := reflect.TypeOf(comp)
	if foundSlice, ok := componentsPool[typeOf]; ok { // todo use getComponentsSlice()
		for i, c := range foundSlice {
			if c == comp {
				componentsPool[typeOf] = append(foundSlice[:i], foundSlice[i+1:]...)
				break
			}
		}
	}

	entityComponents := entsComponents[entity]
	for index, iteratingComp := range entityComponents {
		if comp == iteratingComp {
			entsComponents[entity] = fastRemove(entityComponents, index)
			break
		}
	}

	if len(entsComponents[entity]) == 0 {
		KillEntity(entity)
	}
}

// GetComponent returns component of type T, attached to the entity. If there is component, returns false in 2nd result
func GetComponent[T any](entity Entity) (T, bool) {
	for _, c := range entsComponents[entity] {
		if c2, ok := c.(T); ok {
			return c2, true
		}
	}
	var defaultT T
	return defaultT, false
}

func HasComponent[T any](entity Entity) bool {
	_, has := GetComponent[T](entity)
	return has
}

func getComponentsSlice[T any]() ([]any, int) {
	var defaultT T // workaround :/
	typeT := reflect.TypeOf(defaultT)
	var slice []any

	if foundSlice, ok := componentsPool[typeT]; ok {
		slice = foundSlice
	} else {
		componentsPool[typeT] = slice
	}

	return slice, len(slice)
}

// AddSystem adds system to the game flow. Add them before you call Init and Update. Order is important!
func AddSystem(system System) {
	if isInitialized {
		panic("Can't add Systems after init!")
	}

	systems = append(systems, system)
}

// Init initializes all of your ECS systems. Call it once on game world start before calling Update
func Init() {
	initSystems()
	isInitialized = true
}

func initSystems() {
	for _, sys := range systems {
		if init, ok := sys.(SystemInit); ok {
			init.Init()
		}
	}
}

// Update calls Update of each System every frame. Used to handle most of game logic.
func Update() {
	updateSystems()
	clearOneFrames()
	signals = nil
}

func updateSystems() {
	for _, sys := range systems {
		sys.Update()
	}
}

func clearOneFrames() {
	for _, comp := range oneFrames {
		DelConcreteComponent(comp, GetEntity(comp))
	}

	oneFrames = nil
}

// TryAddSignal adds a new signal to the game flow. If signal of same type was already added, it will be cancelled and return false.
func TryAddSignal[T any](signal T) bool {
	for _, innerSignal := range signals {
		if _, ok := innerSignal.(T); ok {
			return false
		}
	}

	signals = append(signals, signal)
	return true
}

// GetSignal returns registered signal of type T. If there is no signal now, returns false.
func GetSignal[T any]() (T, bool) {
	for _, signal := range signals {
		if typedSignal, ok := signal.(T); ok {
			return typedSignal, true
		}
	}

	var defaultT T
	return defaultT, false
}

// SuspendSignal of type T. It means that signal will not be passed to a next systems.
func SuspendSignal[T any]() {
	for i, signal := range signals {
		if _, ok := signal.(T); ok {
			signals = append(signals[:i], signals[i+1:]...)
			break
		}
	}
}
