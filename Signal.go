package talosecs

var signals []any

// TryAddSignal adds a new signal to the game flow. If signal of same type was already added, it will be cancelled and return false.
func TryAddSignal[T any](signal T) bool {
	if _, isExist := GetSignal[T](); isExist {
		return false
	}

	signals = append(signals, signal)
	return true
}

// GetSignal returns registered signal of type T. If there is no signal now, returns false.
func GetSignal[T any]() (T, bool) {
	for _, signal := range signals {
		if typedSignal, ok := signal.(T); ok {
			return typedSignal, true
		}
	}

	var defaultT T
	return defaultT, false
}

// SuspendSignal of type T. It means that signal will not be passed to a next updateSystems.
func SuspendSignal[T any]() {
	for i, signal := range signals {
		if _, ok := signal.(T); ok {
			signals = append(signals[:i], signals[i+1:]...)
			break
		}
	}
}
