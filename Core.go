package talosecs

var systems []System
var entsAlive = map[Entity]bool{} // todo is map needed there? I think simple slice is better
var entsComponents = map[Entity][]any{}
var currentEntityId Entity
var componentsEnts = map[any]Entity{}
var signals []any

// NewEntity creates new Entity in the game world.
func NewEntity() Entity {
	currentEntityId++
	id := currentEntityId
	entsAlive[id] = true
	return id
}

func GetEntity(comp any) Entity  { return componentsEnts[comp] }
func IsAlive(entity Entity) bool { return entsAlive[entity] }
func SameEntity(a, b any) bool   { return GetEntity(a) == GetEntity(b) }

func KillEntity(entity Entity) {
	if IsAlive(entity) {
		delete(entsAlive, entity)

		for _, component := range entsComponents[entity] {
			delete(componentsEnts, component)
		}

		delete(entsComponents, entity)
	}
}

// AddComponent adds any Component to the specified Entity. Component is a simple data struct.
func AddComponent(entity Entity, comp any) {
	componentsEnts[comp] = entity
	entsComponents[entity] = append(entsComponents[entity], comp)
}

// DelComponent removes component of type T of specified entity. It will be not catched in next systems and next GetComponent (even in the same frame).
func DelComponent[T any](entity Entity) {
	for component, mappedEnt := range componentsEnts {
		if mappedEnt == entity {
			if typedC, ok := component.(T); ok {
				delete(componentsEnts, typedC)
				break
			}
		}
	}

	entityComponents := entsComponents[entity]
	for componentIndex, component := range entityComponents {
		if _, ok := component.(T); ok {
			entsComponents[entity] = fastRemove(entityComponents, componentIndex)
			break
		}
	}
}

func fastRemove(slice []any, i int) []any {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
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

// AddSystem adds system to the game flow. Add them before you call Init and Update. Order is important!
func AddSystem(system System) { systems = append(systems, system) }

// Init initializes all of your ECS systems. Call it once on game world start before calling Update
func Init() {
	for _, sys := range systems {
		if init, ok := sys.(SystemInit); ok {
			init.Init()
		}
	}
}

// Update calls Update of each System every frame. Used to handle most of game logic.
func Update() {
	for _, sys := range systems {
		sys.Update()
	}

	signals = nil
}

func TryAddSignal[T any](signal T) bool {
	for _, innerSignal := range signals {
		if _, ok := innerSignal.(T); ok {
			return false
		}
	}

	signals = append(signals, signal)
	return true
}

func GetSignal[T any]() (T, bool) {
	for _, signal := range signals {
		if typedSignal, ok := signal.(T); ok {
			return typedSignal, true
		}
	}

	var defaultT T
	return defaultT, false
}

func SuspendSignal[T any]() {
	for i, signal := range signals {
		if _, ok := signal.(T); ok {
			signals = append(signals[:i], signals[i+1:]...)
			break
		}
	}
}
