package flags

type Uint64 uint64

func (f Uint64) Is(b Uint64) bool {
	return (f & b) > 0
}

func (f Uint64) Add(b Uint64) Uint64 {
	return f | b
}

func (f Uint64) Remove(b Uint64) Uint64 {
	return f ^ b
}
