package talosecs

var entsComponents = map[Entity][]any{} // for fast KillEntity
var componentsEnts = map[any]Entity{}   // for GetEntity

// AddComponent adds any Component to the specified Entity. Component is a simple data struct.
func AddComponent(entity Entity, comp any) {
	if !isPointer(comp) {
		panic("Only pointers to components allowed!")
	}

	// todo disallow multiple component of the same type

	componentsEnts[comp] = entity
	entsComponents[entity] = append(entsComponents[entity], comp)

	updateFilters(entity)
}

// DelComponent removes component of type T of specified entity. It will be not caught in next updateSystems and next GetComponent (even in the same frame).
func DelComponent[T any](entity Entity) {
	comp, isExist := GetComponent[T](entity)

	if isExist {
		DelSpecificComponent(comp, entity)
	}
}

func DelSpecificComponent(comp any, entity Entity) {
	delete(componentsEnts, comp)

	entityComponents := entsComponents[entity]
	for i, iteratingC := range entityComponents {
		if comp == iteratingC {
			entsComponents[entity] = fastRemove(entityComponents, i)
			break
		}
	}

	if len(entsComponents[entity]) == 0 {
		fullRemoveEntity(entity)
	}

	updateFilters(entity)
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
