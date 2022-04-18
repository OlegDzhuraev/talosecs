package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

var wasLayerUpdated bool

func (system *LayerTestSystem) Update() {
	wasLayerUpdated = true
}

func TestSingleLayerUpdate(t *testing.T) {
	layer := preSetupLayerTest()

	layer.Update()

	if !wasLayerUpdated {
		t.Fatalf("Layer Update was not ran")
	}
}

func TestAddNonSystem(t *testing.T) {
	talosecs.Reset()
	layer := talosecs.NewLayer()
	var notSystem CompA

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Layer Add should fall when adding non-system")
		}
	}()

	layer.Add(notSystem)
}

func TestInactiveLayer(t *testing.T) {
	layer := preSetupLayerTest()

	layer.Active = false
	layer.Update()

	if wasLayerUpdated {
		t.Fatalf("Layer Update was ran, but it should not")
	}
}

func TestRestrictDoubleLayer(t *testing.T) {
	layer := preSetupLayerTest()

	if talosecs.AddLayer(layer) {
		t.Fatalf("Layer should not be added twice times")
	}
}

func TestAddSystemAfterInit(t *testing.T) {
	talosecs.Reset()
	layer := talosecs.NewLayer()
	talosecs.Init()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Add system try after ECS was inited should fail")
		}
	}()

	layer.Add(&SystemA{})
}

func preSetupLayerTest() *talosecs.Layer {
	wasLayerUpdated = false
	talosecs.Reset()
	layer := talosecs.NewLayer()
	talosecs.AddLayer(layer)
	layer.Add(&LayerTestSystem{})
	talosecs.Init()

	return layer
}
