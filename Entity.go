package talosecs

type Entity uint32

var currentEntityId Entity

// Add is similar to the AddComponent(entity, component), just shorter way to do this
func (e Entity) Add(c any) { AddComponent(e, c) }

// OneFrame is similar to the AddOneFrame(entity, component), just shorter way to do this
func (e Entity) OneFrame(c any) { AddOneFrame(e, c) }

// NewEntity creates new Entity in the game world.
func NewEntity() Entity {
	currentEntityId++
	id := currentEntityId
	entsComponents[id] = []any{}

	return id
}

func GetEntity(comp any) Entity        { return componentsEnts[comp] }
func SameEntity(compA, compB any) bool { return GetEntity(compA) == GetEntity(compB) }

func IsAlive(entity Entity) bool {
	_, exist := entsComponents[entity]
	return exist
}

func KillEntity(entity Entity) {
	if IsAlive(entity) {
		for _, component := range entsComponents[entity] {
			DelSpecificComponent(component, entity)
		}
	}
}

func fullRemoveEntity(entity Entity) {
	if IsAlive(entity) {
		delete(entsComponents, entity)
	}
}
