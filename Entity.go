package talosecs

type Entity int64

// Add is similar to the AddComponent(entity, component), just shorter way to do this
func (e Entity) Add(c any) { AddComponent(e, c) }
