package talosecs

type Entity int64
type EntitySet map[Entity]bool

// Add is similar to the AddComponent(entity, component), just shorter way to do this
func (e Entity) Add(c any) { AddComponent(e, c) }
