package flags

type Uint16 uint16

func (f Uint16) Add(b ...Uint16) Uint16 {
	for i := 0; i < len(b); i++ {
		f = f | b[i]
	}
	return f
}

func (f Uint16) Remove(b ...Uint16) Uint16 {
	for i := 0; i < len(b); i++ {
		f = f ^ b[i]
	}
	return f
}

func (f Uint16) Intersect(b ...Uint16) Uint16 {
	s := Uint16(0).Add(b...)
	return f & s
}

func (f Uint16) IsAny(b ...Uint16) bool {
	return f.Intersect(b...) > 0
}

func (f Uint16) IsAll(b ...Uint16) bool {
	return f.Intersect(b...) == f
}
