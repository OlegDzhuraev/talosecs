package talosecs

import "reflect"

func fastRemove[T any](slice []T, i int) []T {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func indexOf[T comparable](slice []T, obj T) (int, bool) {
	for i, sliceObj := range slice {
		if sliceObj == obj {
			return i, true
		}
	}
	return 0, false
}

func removeIfContains[T comparable](slice []T, obj T) []T {
	if i, ok := indexOf(slice, obj); ok {
		return fastRemove(slice, i)
	}

	return slice
}

func appendIfNotContains[T comparable](slice []T, obj T) []T {
	if _, ok := indexOf(slice, obj); ok {
		return slice
	}

	return append(slice, obj)
}

func isPointer(obj any) bool { return reflect.ValueOf(obj).Kind() == reflect.Ptr }
