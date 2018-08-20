package flags

type Byte byte

func (f Byte) Is(b Byte) bool {
	return (f & b) > 0
}

func (f Byte) Add(b Byte) Byte {
	return f | b
}

func (f Byte) Remove(b Byte) Byte {
	return f ^ b
}
