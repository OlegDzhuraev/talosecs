package talosecs

var oneFrames []any

// AddOneFrame is same to AddComponent, but components added with this function live only one frame.
func AddOneFrame(entity Entity, comp any) {
	oneFrames = append(oneFrames, comp)
	AddComponent(entity, comp)
}

func clearOneFrames() {
	for _, of := range oneFrames {
		DelSpecificComponent(of, GetEntity(of))
	}

	oneFrames = nil
}
