package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func (system *SystemA) Init() {
	wasSystemInited = true
}

func (system *SystemA) Update() {
	wasSystemUpdated = true
}

var wasSystemInited bool
var wasSystemUpdated bool

func TestAddAndInitSystem(t *testing.T) {
	preSetupSystemsTest()
	talosecs.Init()

	if !wasSystemInited {
		t.Fatalf("UpdateSystem init was not runned, system was not added")
	}
}

func TestUpdate(t *testing.T) {
	preSetupSystemsTest()

	talosecs.Init()
	talosecs.Update()

	if !wasSystemUpdated {
		t.Fatalf("UpdateSystem Update was not runned")
	}
}

func preSetupSystemsTest() {
	wasSystemUpdated = false
	talosecs.Reset()
	layer := talosecs.NewLayer()
	talosecs.AddLayer(layer)
	talosecs.AddSystem(layer, &SystemA{})
}
