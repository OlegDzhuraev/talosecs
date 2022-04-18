package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func TestFilterWithAmount(t *testing.T) {
	talosecs.Reset()
	makeEntA()

	filter := talosecs.FilterWith[*CompA]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterWith2Amount(t *testing.T) {
	talosecs.Reset()
	makeEntA()
	makeEntAb()

	filter, _ := talosecs.FilterWith2[*CompA, *CompB]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterWithCorrectComponent(t *testing.T) {
	talosecs.Reset()
	makeEntA()

	filter := talosecs.FilterWith[*CompA]()

	var result *CompA
	for _, c := range filter {
		result = c
		break
	}

	if result == nil {
		t.Fatalf("No component found")
	}
}

func TestFilterWith3Amount(t *testing.T) {
	talosecs.Reset()
	makeEntA()
	makeEntAbc()

	filter, _, _ := talosecs.FilterWith3[*CompA, *CompC, *CompB]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW1Excl1(t *testing.T) {
	talosecs.Reset()
	makeEntA()
	makeEntAb()

	filter := talosecs.FilterW1Excl1[*CompA, *CompB]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW2Excl1Amount(t *testing.T) {
	talosecs.Reset()
	makeEntA()
	makeEntAb()
	makeEntAc()

	filter, _ := talosecs.FilterW2Excl1[*CompA, *CompC, *CompB]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW1Excl2Amount(t *testing.T) {
	talosecs.Reset()

	makeEntA()
	makeEntAb()
	makeEntAc()
	makeEntAbc()

	filter := talosecs.FilterW1Excl2[*CompA, *CompB, *CompC]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW2Excl2Amount(t *testing.T) {
	talosecs.Reset()

	makeEntA()
	makeEntAb()
	makeEntAbc()
	makeEntAbcd()

	filter, _ := talosecs.FilterW2Excl2[*CompA, *CompB, *CompC, *CompD]()

	amount := 0
	for range filter {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func makeEntA() {
	e := talosecs.NewEntity()
	e.Add(&CompA{})
}

func makeEntAb() {
	e := talosecs.NewEntity()
	e.Add(&CompA{})
	e.Add(&CompB{})
}

func makeEntAc() {
	e := talosecs.NewEntity()
	e.Add(&CompA{})
	e.Add(&CompC{})
}

func makeEntAbc() {
	e := talosecs.NewEntity()
	e.Add(&CompA{})
	e.Add(&CompB{})
	e.Add(&CompC{})
}

func makeEntAbcd() {
	e := talosecs.NewEntity()
	e.Add(&CompA{})
	e.Add(&CompB{})
	e.Add(&CompC{})
	e.Add(&CompD{})
}
