package talosecs

type Layer struct {
	Active        bool
	initSystems   []InitSystem
	updateSystems []UpdateSystem
}

func NewLayer() *Layer { return &Layer{Active: true} }

func AddSystem(layer *Layer, system any) {
	if isInitialized {
		panic("Can't add Systems after init!")
	}

	isSystem := false
	if initSystem, ok := system.(InitSystem); ok {
		layer.initSystems = append(layer.initSystems, initSystem)
		isSystem = true
	}
	if updateSystem, ok := system.(UpdateSystem); ok {
		layer.updateSystems = append(layer.updateSystems, updateSystem)
		isSystem = true
	}

	if !isSystem {
		panic("Not a system!")
	}
}

func (layer *Layer) Add(system any) { AddSystem(layer, system) }

func (layer *Layer) Init() {
	for _, initSystem := range layer.initSystems {
		initSystem.Init()
	}
}

func (layer *Layer) Update() {
	if layer.Active {
		for _, updateSystem := range layer.updateSystems {
			updateSystem.Update()
		}
	}
}

var layers []*Layer

func AddLayer(layer *Layer) {
	for _, registeredLayer := range layers {
		if registeredLayer == layer {
			return
		}
	}

	layers = append(layers, layer)
}

func initLayers() {
	for _, layer := range layers {
		layer.Init()
	}
}

func updateLayers() {
	for _, layer := range layers {
		layer.Update()
	}
}
