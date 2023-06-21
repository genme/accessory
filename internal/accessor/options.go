package accessor

type Option func(*generator)

// Type sets type name to genarator.
func Type(typeName string) Option {
	return func(g *generator) {
		g.typ = typeName
	}
}

// ForceGetters sets force generation for getters
func ForceGetters(getters bool) Option {
	return func(g *generator) {
		g.getters = getters
	}
}

// ForceSetters sets force generation for setters
func ForceSetters(setters bool) Option {
	return func(g *generator) {
		g.setters = setters
	}
}

// Output sets output file path to genarator.
func Output(output string) Option {
	return func(g *generator) {
		g.output = output
	}
}

// Receiver sets receiver name to genarator.
func Receiver(receiver string) Option {
	return func(g *generator) {
		g.receiver = receiver
	}
}

// Lock sets lock field name to genarator.
func Lock(lock string) Option {
	return func(g *generator) {
		g.lock = lock
	}
}
