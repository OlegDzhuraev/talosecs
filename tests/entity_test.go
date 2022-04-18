package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func TestEntityIsAlive(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	e.Add(&CompA{})

	if !talosecs.IsAlive(e) {
		t.Fatalf("Entity should be alive, but it is not")
	}
}

func TestEntityDiesWithoutComponents(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	e.Add(&CompA{})

	talosecs.DelComponent[*CompA](e)

	if talosecs.IsAlive(e) {
		t.Fatalf("Entity should be died, but it is not")
	}
}

func TestKillEntity(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	e.Add(&CompA{})

	talosecs.KillEntity(e)

	if talosecs.IsAlive(e) {
		t.Fatalf("Entity should be died, but it is not")
	}
}

func TestGetEntity(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	cA := &CompA{}
	e.Add(cA)

	entId := talosecs.GetEntity(cA)
	if entId != 1 {
		t.Fatalf("Entity id should be 1, but it is %v", entId)
	}
}

func TestSameEntity(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	cA := &CompA{}
	cB := &CompB{}
	e.Add(cA)
	e.Add(cB)

	if !talosecs.SameEntity(cA, cB) {
		t.Fatalf("Entity should be same for both components, but it is not")
	}
}

func TestEntityAddNonPointerComponent(t *testing.T) {
	talosecs.Reset()
	e := talosecs.NewEntity()
	cA := CompA{}

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Add should fall when adding non-pointer component")
		}
	}()

	e.Add(cA)
}
