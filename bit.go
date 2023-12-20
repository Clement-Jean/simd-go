package simdgo

//go:noescape
func BitAndU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) BitAnd(b Uint8x16) Uint8x16 {
	return BitAndU8x16(a, b)
}

//go:noescape
func BitAndU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) BitAnd(b Uint8x32) Uint8x32 {
	return BitAndU8x32(a, b)
}

//go:noescape
func BitAndU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) BitAnd(b Uint8x64) Uint8x64 {
	return BitAndU8x64(a, b)
}

//go:noescape
func BitOrU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) BitOr(b Uint8x16) Uint8x16 {
	return BitOrU8x16(a, b)
}

//go:noescape
func BitOrU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) BitOr(b Uint8x32) Uint8x32 {
	return BitOrU8x32(a, b)
}

//go:noescape
func BitOrU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) BitOr(b Uint8x64) Uint8x64 {
	return BitOrU8x64(a, b)
}

//go:noescape
func BitXorU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) BitXor(b Uint8x16) Uint8x16 {
	return BitXorU8x16(a, b)
}

//go:noescape
func BitXorU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) BitXor(b Uint8x32) Uint8x32 {
	return BitXorU8x32(a, b)
}

//go:noescape
func BitXorU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) BitXor(b Uint8x64) Uint8x64 {
	return BitXorU8x64(a, b)
}

func BitNotU8x16(a, b Uint8x16) Uint8x16 {
	return BitXorU8x16(a, SplatU8x16(0xFF))
}

func (a Uint8x16) BitNot(b Uint8x16) Uint8x16 {
	return BitNotU8x16(a, b)
}
