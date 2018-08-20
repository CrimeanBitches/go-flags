package flags

type Byte byte

func (f Byte) Add(b ...Byte) Byte {
	for i := 0; i < len(b); i++ {
		f = f | b[i]
	}
	return f
}

func (f Byte) Remove(b ...Byte) Byte {
	for i := 0; i < len(b); i++ {
		f = f ^ b[i]
	}
	return f
}

func (f Byte) Intersect(b ...Byte) Byte {
	s := Byte(0).Add(b...)
	return f & s
}

func (f Byte) IsAny(b ...Byte) bool {
	return f.Intersect(b...) > 0
}

func (f Byte) IsAll(b ...Byte) bool {
	return f.Intersect(b...) == f
}
