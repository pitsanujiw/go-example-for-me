package pointer

type PointerExample struct {
	name string
	value string
}

func (p *PointerExample) Load() {
	p.name = "Hello world"
	p.value = "value"
}
