package flags_test

import "testing"
import "grisp/protocol/flags"

func TestByteFlags(t *testing.T) {
	v := make([]flags.Byte, 8)
	for i := uint(0); i < 8; i++ {
		v[i] = flags.Byte(1 << i)
	}

	flag := flags.Byte(0)

	for i := 0; i < len(v); i++ {
		flag = flag.Add(v[i])
	}

	if byte(flag) != 255 {
		t.Errorf("add error. expected: %v, got: %v", 255, byte(flag))
	}

	for i := 0; i < len(v); i++ {
		if !flag.Is(v[i]) {
			t.Error("is error")
		}
	}

	for i := 0; i < len(v); i++ {
		flag = flag.Remove(v[i])
	}

	if byte(flag) != 0 {
		t.Error("remove error")
	}

	for i := 0; i < len(v); i++ {
		if flag.Is(v[i]) {
			t.Error("not is error")
		}
	}
}

func TestUint16Flags(t *testing.T) {
	v := make([]flags.Uint16, 16)
	for i := uint(0); i < 16; i++ {
		v[i] = flags.Uint16(1 << i)
	}

	flag := flags.Uint16(0)

	for i := 0; i < len(v); i++ {
		flag = flag.Add(v[i])
	}

	if uint16(flag) != 65535 {
		t.Error("add error")
	}

	for i := 0; i < len(v); i++ {
		if !flag.Is(v[i]) {
			t.Error("is error")
		}
	}

	for i := 0; i < len(v); i++ {
		flag = flag.Remove(v[i])
	}

	if uint16(flag) != 0 {
		t.Error("remove error")
	}

	for i := 0; i < len(v); i++ {
		if flag.Is(v[i]) {
			t.Error("not is error")
		}
	}
}

func TestUint32Flags(t *testing.T) {
	v := make([]flags.Uint32, 32)
	for i := uint(0); i < 32; i++ {
		v[i] = flags.Uint32(1 << i)
	}

	flag := flags.Uint32(0)

	for i := 0; i < len(v); i++ {
		flag = flag.Add(v[i])
	}

	if uint32(flag) != 4294967295 {
		t.Error("add error")
	}

	for i := 0; i < len(v); i++ {
		if !flag.Is(v[i]) {
			t.Error("is error")
		}
	}

	for i := 0; i < len(v); i++ {
		flag = flag.Remove(v[i])
	}

	if uint32(flag) != 0 {
		t.Error("remove error")
	}

	for i := 0; i < len(v); i++ {
		if flag.Is(v[i]) {
			t.Error("not is error")
		}
	}
}

func TestUint64Flags(t *testing.T) {
	v := make([]flags.Uint64, 64)
	for i := uint(0); i < 64; i++ {
		v[i] = flags.Uint64(1 << i)
	}

	flag := flags.Uint64(0)

	for i := 0; i < len(v); i++ {
		flag = flag.Add(v[i])
	}

	if uint64(flag) != 18446744073709551615 {
		t.Error("add error")
	}

	for i := 0; i < len(v); i++ {
		if !flag.Is(v[i]) {
			t.Error("is error")
		}
	}

	for i := 0; i < len(v); i++ {
		flag = flag.Remove(v[i])
	}

	if uint64(flag) != 0 {
		t.Error("remove error")
	}

	for i := 0; i < len(v); i++ {
		if flag.Is(v[i]) {
			t.Error("not is error")
		}
	}
}
