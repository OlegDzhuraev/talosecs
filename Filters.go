package talosecs

func FilterWith[A any]() []A {
	slice, length := getComponentsGeneric[A]()
	var result = make([]A, length)

	for i, c := range slice { // I can't just return slice because it is of type []any, and there is no fast way of conversion to the []A
		if compA, ok := c.(A); ok {
			result[i] = compA
		}
	}

	return result
}

func FilterWith2[A any, B any]() ([]A, []B) {
	sliceA, maxLength := getComponentsGeneric[A]()
	var resultA = make([]A, maxLength)
	var resultB = make([]B, maxLength)

	i := 0
	for _, c := range sliceA {
		if compA, ok := c.(A); ok {
			e := GetEntity(compA)
			if compB, ok2 := GetComponent[B](e); ok2 {
				resultA[i] = compA
				resultB[i] = compB
				i++
			}
		}
	}

	return resultA[:i], resultB[:i]
}

func FilterWith3[A any, B any, C any]() ([]A, []B, []C) {
	sliceA, maxLength := getComponentsGeneric[A]()
	var resultA = make([]A, maxLength)
	var resultB = make([]B, maxLength)
	var resultC = make([]C, maxLength)

	i := 0
	for _, c := range sliceA {
		if compA, ok := c.(A); ok {
			e := GetEntity(compA)
			if compB, ok2 := GetComponent[B](e); ok2 {
				if compC, ok3 := GetComponent[C](e); ok3 {
					resultA[i] = compA
					resultB[i] = compB
					resultC[i] = compC
					i++
				}
			}
		}
	}

	return resultA[:i], resultB[:i], resultC[:i]
}

func FilterW1Excl1[A any, Excl any]() []A {
	slice, length := getComponentsGeneric[A]()
	var result = make([]A, length)

	i := 0
	for _, c := range slice {
		e := GetEntity(c)
		if compA, ok := c.(A); ok && !HasComponent[Excl](e) {
			result[i] = compA
			i++
		}
	}

	return result[:i]
}

func FilterW2Excl1[A any, B any, Excl any]() ([]A, []B) {
	sliceA, maxLength := getComponentsGeneric[A]()
	var resultA = make([]A, maxLength)
	var resultB = make([]B, maxLength)

	i := 0
	for _, c := range sliceA {
		if compA, ok := c.(A); ok {
			e := GetEntity(compA)
			if compB, ok2 := GetComponent[B](e); ok2 && !HasComponent[Excl](e) {
				resultA[i] = compA
				resultB[i] = compB
				i++
			}
		}
	}

	return resultA[:i], resultB[:i]
}

func FilterW1Excl2[A any, ExclA any, ExclB any]() []A {
	slice, length := getComponentsGeneric[A]()
	var result = make([]A, length)

	i := 0
	for _, c := range slice {
		e := GetEntity(c)
		if compA, ok := c.(A); ok && !HasComponent[ExclA](e) && !HasComponent[ExclB](e) {
			result[i] = compA
			i++
		}
	}

	return result[:i]
}

func FilterW2Excl2[A any, B any, ExclA any, ExclB any]() ([]A, []B) {
	sliceA, maxLength := getComponentsGeneric[A]()
	var resultA = make([]A, maxLength)
	var resultB = make([]B, maxLength)

	i := 0
	for _, c := range sliceA {
		if compA, ok := c.(A); ok {
			e := GetEntity(compA)
			if compB, ok2 := GetComponent[B](e); ok2 && !HasComponent[ExclA](e) && !HasComponent[ExclB](e) {
				resultA[i] = compA
				resultB[i] = compB
				i++
			}
		}
	}

	return resultA[:i], resultB[:i]
}
