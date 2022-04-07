package talosecs

type UpdateSystem interface {
	Update()
}

type InitSystem interface {
	Init()
}

// todo add render/ui system?

var updateSystems []UpdateSystem
var initSystems []InitSystem

// AddSystem adds system to the game flow. Add them before you call Init and Update. Order is important!
func AddSystem(system any) {
	if isInitialized {
		panic("Can't add Systems after init!")
	}

	isSystem := false
	if initSystem, ok := system.(InitSystem); ok {
		initSystems = append(initSystems, initSystem)
		isSystem = true
	}
	if updateSystem, ok := system.(UpdateSystem); ok {
		updateSystems = append(updateSystems, updateSystem)
		isSystem = true
	}

	if !isSystem {
		panic("Not a system!")
	}
}

func doInitSystems() {
	for _, is := range initSystems {
		is.Init()
	}
}

func doUpdateSystems() {
	for _, us := range updateSystems {
		us.Update()
	}
}
