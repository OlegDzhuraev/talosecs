package talosEcsTests

import (
	"github.com/OlegDzhuraev/talosecs"
	"testing"
)

func TestFilterWithAmount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterWith[*CompA]{}
	talosecs.AddFilter(filter)

	makeEntA()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterWith2Amount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW2[*CompA, *CompB]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAb()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterWithCorrectComponent(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterWith[*CompA]{}
	talosecs.AddFilter(filter)

	makeEntA()

	var result *CompA
	for _, c := range filter.A {
		result = c
		break
	}

	if result == nil {
		t.Fatalf("No component found")
	}
}

func TestFilterWith3Amount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW3[*CompA, *CompC, *CompB]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAbc()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW1Excl1(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW1Exc1[*CompA, *CompB]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAb()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW2Excl1Amount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW2Exc1[*CompA, *CompC, *CompB]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAb()
	makeEntAc()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW1Excl2Amount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW1Exc2[*CompA, *CompB, *CompC]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAb()
	makeEntAc()
	makeEntAbc()

	amount := 0
	for range filter.A {
		amount++
	}

	if amount != 1 {
		t.Fatalf("Should be found only 1 component, found %v", amount)
	}
}

func TestFilterW2Excl2Amount(t *testing.T) {
	talosecs.Reset()

	filter := &talosecs.FilterW2Exc2[*CompA, *CompB, *CompC, *CompD]{}
	talosecs.AddFilter(filter)

	makeEntA()
	makeEntAb()
	makeEntAbc()
	makeEntAbcd()

	amount := 0
	for range filter.A {
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
