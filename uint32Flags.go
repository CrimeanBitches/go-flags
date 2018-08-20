package flags

type Uint32 uint32

func (f Uint32) Add(b ...Uint32) Uint32 {
	for i := 0; i < len(b); i++ {
		f = f | b[i]
	}
	return f
}

func (f Uint32) Remove(b ...Uint32) Uint32 {
	for i := 0; i < len(b); i++ {
		f = f ^ b[i]
	}
	return f
}

func (f Uint32) Intersect(b ...Uint32) Uint32 {
	s := Uint32(0).Add(b...)
	return f & s
}

func (f Uint32) IsAny(b ...Uint32) bool {
	return f.Intersect(b...) > 0
}

func (f Uint32) IsAll(b ...Uint32) bool {
	return f.Intersect(b...) == f
}
