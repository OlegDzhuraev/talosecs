package talosecs

import "reflect"

var entsComponents = map[Entity][]any{}       // for filters
var componentsEnts = map[any]Entity{}         // for GetEntity
var componentsPool = map[reflect.Type][]any{} // for GetComponent + filters too

// AddComponent adds any Component to the specified Entity. Component is a simple data struct.
func AddComponent(entity Entity, comp any) {
	if !isPointer(comp) {
		panic("Only pointers to components allowed!")
	}

	componentsSlice, typeOf := getComponentsSpecificOutType(comp)
	componentsSlice = append(componentsSlice, comp)
	componentsPool[typeOf] = componentsSlice

	componentsEnts[comp] = entity
	entsComponents[entity] = append(entsComponents[entity], comp)
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

	componentsSlice, typeOf := getComponentsSpecificOutType(comp)
	for i, c := range componentsSlice {
		if c == comp {
			componentsPool[typeOf] = fastRemove(componentsSlice, i)
			break
		}
	}

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

func getComponentsGeneric[T any]() ([]any, int) {
	var defaultT T // workaround :/
	return getComponentsSameTo(defaultT)
}

func getComponentsSameTo(comp any) ([]any, int) {
	return getComponentsOfType(reflect.TypeOf(comp))
}

func getComponentsSpecificOutType(comp any) ([]any, reflect.Type) {
	rType := reflect.TypeOf(comp)
	slice, _ := getComponentsOfType(rType)
	return slice, rType
}

func getComponentsOfType(rType reflect.Type) ([]any, int) {
	var slice []any

	if foundSlice, ok := componentsPool[rType]; ok {
		slice = foundSlice
	} else {
		componentsPool[rType] = slice
	}

	return slice, len(slice)
}
