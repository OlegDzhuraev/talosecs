package talosecs

func FilterWith[T any]() []T {
	var result []T
	for c := range componentsEnts {
		if typedC, ok := c.(T); ok {
			result = append(result, typedC)
		}
	}

	return result
}

func FilterWith2[A any, B any]() ([]A, []B) {
	var resultA []A
	var resultB []B

	for c, entId := range componentsEnts {
		if typeA, ok := c.(A); ok {
			if typeB, ok2 := GetComponent[B](entId); ok2 {
				resultA = append(resultA, typeA)
				resultB = append(resultB, typeB)
			}
		}
	}

	return resultA, resultB
}

func FilterWith3[A any, B any, C any]() ([]A, []B, []C) {
	var resultA []A
	var resultB []B
	var resultC []C

	for c, entId := range componentsEnts {
		if typeA, ok := c.(A); ok {
			if typeB, ok2 := GetComponent[B](entId); ok2 {
				if typeC, ok3 := GetComponent[C](entId); ok3 {
					resultA = append(resultA, typeA)
					resultB = append(resultB, typeB)
					resultC = append(resultC, typeC)
				}
			}
		}
	}

	return resultA, resultB, resultC
}

func FilterW1Excl1[With any, Without any]() []With {
	var result []With
	for c := range componentsEnts {
		if typedC, ok := c.(With); ok {
			if _, isExclude := c.(Without); isExclude == false {
				result = append(result, typedC)
			}
		}
	}

	return result
}

func FilterW2Excl1[WithA any, WithB any, Without any]() ([]WithA, []WithB) {
	var resultA []WithA
	var resultB []WithB

	for c, entId := range componentsEnts {
		if typedA, ok := c.(WithA); ok {
			if typedB, ok2 := GetComponent[WithB](entId); ok2 {
				if _, isExclude := c.(Without); isExclude == false {
					resultA = append(resultA, typedA)
					resultB = append(resultB, typedB)
				}
			}
		}
	}

	return resultA, resultB
}

func FilterW1Excl2[WithA any, WithoutA any, WithoutB any]() []WithA {
	var result []WithA

	for c := range componentsEnts {
		if typedC, ok := c.(WithA); ok {
			if _, isExcludeA := c.(WithoutA); isExcludeA == false {
				if _, isExcludeB := c.(WithoutB); isExcludeB == false {
					result = append(result, typedC)
				}
			}
		}
	}

	return result
}

func FilterW2Excl2[WithA any, WithB any, WithoutA any, WithoutB any]() ([]WithA, []WithB) {
	var resultA []WithA
	var resultB []WithB

	for c, entId := range componentsEnts {
		if typedA, ok := c.(WithA); ok {
			if typedB, ok2 := GetComponent[WithB](entId); ok2 {
				if _, isExcludeA := c.(WithoutA); isExcludeA == false {
					if _, isExcludeB := c.(WithoutB); isExcludeB == false {
						resultA = append(resultA, typedA)
						resultB = append(resultB, typedB)
					}
				}
			}
		}
	}

	return resultA, resultB
}
