package flags

type Uint64 uint64

func (f Uint64) Add(b ...Uint64) Uint64 {
	for i := 0; i < len(b); i++ {
		f = f | b[i]
	}
	return f
}

func (f Uint64) Remove(b ...Uint64) Uint64 {
	for i := 0; i < len(b); i++ {
		f = f ^ b[i]
	}
	return f
}

func (f Uint64) Intersect(b ...Uint64) Uint64 {
	s := Uint64(0).Add(b...)
	return f & s
}

func (f Uint64) IsAny(b ...Uint64) bool {
	return f.Intersect(b...) > 0
}

func (f Uint64) IsAll(b ...Uint64) bool {
	return f.Intersect(b...) == f
}
