package talosecs

type Entity uint32

// Add is similar to the AddComponent(entity, component), just shorter way to do this
func (e Entity) Add(c any) { AddComponent(e, c) }
