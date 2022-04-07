package talosecs

import "reflect"

func fastRemove(slice []any, i int) []any {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func isPointer(obj any) bool { return reflect.ValueOf(obj).Kind() == reflect.Ptr }
