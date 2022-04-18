package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func TestTryAddSignalFunc(t *testing.T) {
	talosecs.Reset()
	if !talosecs.TryAddSignal(&CompA{}) {
		t.Fatalf(`No Signal was added by initial TryAddSignal`)
	}
}

func TestSignalAdd(t *testing.T) {
	talosecs.Reset()
	talosecs.TryAddSignal(&CompA{})
	if _, ok := talosecs.GetSignal[*CompA](); !ok {
		t.Fatalf(`No Signal was added`)
	}
}

func TestSignalDoubleAdd(t *testing.T) {
	talosecs.Reset()
	talosecs.TryAddSignal(&CompA{})

	if talosecs.TryAddSignal(&CompA{}) {
		t.Fatalf(`Double signals is not allowed`)
	}
}

func TestSignalRemove(t *testing.T) {
	talosecs.Reset()
	talosecs.TryAddSignal(&CompA{})
	talosecs.Update()

	if _, ok := talosecs.GetSignal[*CompA](); ok {
		t.Fatalf(`Signal was not removed after one frame was passed`)
	}
}

func TestSuspendSignal(t *testing.T) {
	talosecs.Reset()
	talosecs.TryAddSignal(&CompA{})
	talosecs.SuspendSignal[*CompA]()

	if _, ok := talosecs.GetSignal[*CompA](); ok {
		t.Fatalf(`Signal was not removed after suspend call`)
	}
}
