package talosecs

type System interface {
	Update()
}

type SystemInit interface {
	Init()
}

// todo render/ui system?
