package flags

type Uint16 uint16

func (f Uint16) Is(b Uint16) bool {
	return (f & b) > 0
}

func (f Uint16) Add(b Uint16) Uint16 {
	return f | b
}

func (f Uint16) Remove(b Uint16) Uint16 {
	return f ^ b
}
