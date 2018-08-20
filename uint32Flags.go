package flags

type Uint32 uint32

func (f Uint32) Is(b Uint32) bool {
	return (f & b) > 0
}

func (f Uint32) Add(b Uint32) Uint32 {
	return f | b
}

func (f Uint32) Remove(b Uint32) Uint32 {
	return f ^ b
}
