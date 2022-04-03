package talosecs

func fastRemove(slice []any, i int) []any {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
