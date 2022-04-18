package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func TestDoubleInitNotAllowed(t *testing.T) {
	talosecs.Reset()

	talosecs.Init()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Double init should fail")
		}
	}()

	talosecs.Init()
}
