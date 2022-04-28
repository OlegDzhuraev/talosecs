package talosecs

var filters []IFilter

type IFilter interface {
	addIfMatch(ent Entity)
}

type IEntitiesBox interface {
	resetEntities()
}

type EntitiesBox struct {
	Entities map[Entity]int // entity mapped to the filter components arrays index
}

func (eb *EntitiesBox) resetEntities() {
	eb.Entities = map[Entity]int{}
}

func AddFilter(filter any) {
	if f, ok := filter.(IFilter); ok {
		filters = append(filters, f)

		if box, ok2 := filter.(IEntitiesBox); ok2 {
			box.resetEntities()
		}
	}
}

func updateFilters(entity Entity) {
	for _, f := range filters {
		f.addIfMatch(entity)
	}
}

type FilterWith[A comparable] struct {
	EntitiesBox // todo make sparse set?
	A           []A
}

type FilterW2[A comparable, B comparable] struct {
	EntitiesBox
	A []A
	B []B
}

type FilterW3[A comparable, B comparable, C comparable] struct {
	EntitiesBox
	A []A
	B []B
	C []C
}

type FilterW1Exc1[A comparable, E1 comparable] struct {
	EntitiesBox
	A []A
}

type FilterW2Exc1[A comparable, B comparable, E1 comparable] struct {
	EntitiesBox
	A []A
	B []B
}

type FilterW1Exc2[A comparable, E1 comparable, E2 comparable] struct {
	EntitiesBox
	A []A
}

type FilterW2Exc2[A comparable, B comparable, E1 comparable, E2 comparable] struct {
	EntitiesBox
	A []A
	B []B
}

func (f *FilterWith[A]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if a, ok := match[A](entity); ok {
		if !exist {
			f.A = append(f.A, a)
			f.Entities[entity] = len(f.A) - 1
		}
	} else if exist {
		f.A = fastRemove(f.A, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW2[A, B]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if a, b, ok := match2[A, B](entity); ok {
		if !exist {
			f.A = append(f.A, a)
			f.B = append(f.B, b)
			f.Entities[entity] = len(f.A) - 1
		}
	} else if exist {
		f.A = fastRemove(f.A, i)
		f.B = fastRemove(f.B, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW3[A, B, C]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if a, b, c, ok := match3[A, B, C](entity); ok {
		if !exist {
			f.A = append(f.A, a)
			f.B = append(f.B, b)
			f.C = append(f.C, c)
			f.Entities[entity] = len(f.A) - 1
		}
	} else if exist {
		f.A = fastRemove(f.A, i)
		f.B = fastRemove(f.B, i)
		f.C = fastRemove(f.C, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW1Exc1[A, E1]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if _, hasExclude := match[E1](entity); !hasExclude {
		if a, ok := match[A](entity); ok {
			if !exist {
				f.A = append(f.A, a)
				f.Entities[entity] = len(f.A) - 1
			}
			return
		}
	}

	if exist {
		f.A = fastRemove(f.A, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW2Exc1[A, B, E1]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if _, hasExclude := match[E1](entity); !hasExclude {
		if a, b, ok := match2[A, B](entity); ok {
			if !exist {
				f.A = append(f.A, a)
				f.B = append(f.B, b)
				f.Entities[entity] = len(f.A) - 1
			}
			return
		}
	}

	if exist {
		f.A = fastRemove(f.A, i)
		f.B = fastRemove(f.B, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW1Exc2[A, E1, E2]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if _, _, hasExclude := match2any[E1, E2](entity); !hasExclude {
		if a, ok := match[A](entity); ok {
			if !exist {
				f.A = append(f.A, a)
				f.Entities[entity] = len(f.A) - 1
			}
			return
		}
	}

	if exist {
		f.A = fastRemove(f.A, i)
		delete(f.Entities, entity)
	}
}

func (f *FilterW2Exc2[A, B, E1, E2]) addIfMatch(entity Entity) {
	i, exist := f.Entities[entity]

	if _, _, hasExclude := match2any[E1, E2](entity); !hasExclude {
		if a, b, ok := match2[A, B](entity); ok {
			if !exist {
				f.A = append(f.A, a)
				f.B = append(f.B, b)
				f.Entities[entity] = len(f.A) - 1
			}
			return
		}
	}

	if exist {
		f.A = fastRemove(f.A, i)
		f.B = fastRemove(f.B, i)
		delete(f.Entities, entity)
	}
}

func match[A comparable](e Entity) (A, bool) {
	return GetComponent[A](e)
}

func match2[A comparable, B comparable](e Entity) (A, B, bool) {
	a, okA := GetComponent[A](e)
	b, okB := GetComponent[B](e)
	return a, b, okA && okB
}

func match2any[A comparable, B comparable](e Entity) (A, B, bool) {
	a, okA := GetComponent[A](e)
	b, okB := GetComponent[B](e)
	return a, b, okA || okB
}

func match3[A comparable, B comparable, C comparable](e Entity) (A, B, C, bool) {
	a, okA := GetComponent[A](e)
	b, okB := GetComponent[B](e)
	c, okC := GetComponent[C](e)
	return a, b, c, okA && okB && okC
}
