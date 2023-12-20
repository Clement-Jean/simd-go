package simdgo

//go:noescape
func AddU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) Add(b Uint8x16) Uint8x16 {
	return AddU8x16(a, b)
}

//go:noescape
func AddU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) Add(b Uint8x32) Uint8x32 {
	return AddU8x32(a, b)
}

//go:noescape
func AddU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) Add(b Uint8x64) Uint8x64 {
	return AddU8x64(a, b)
}

//go:noescape
func SubU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) Sub(b Uint8x16) Uint8x16 {
	return SubU8x16(a, b)
}

//go:noescape
func SubU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) Sub(b Uint8x32) Uint8x32 {
	return SubU8x32(a, b)
}

//go:noescape
func SubU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) Sub(b Uint8x64) Uint8x64 {
	return SubU8x64(a, b)
}

//go:noescape
func EqualU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) Equal(b Uint8x16) Uint8x16 {
	return EqualU8x16(a, b)
}

//go:noescape
func EqualU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) Equal(b Uint8x32) Uint8x32 {
	return EqualU8x32(a, b)
}

//go:noescape
func EqualU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) Equal(b Uint8x64) Uint8x64 {
	return EqualU8x64(a, b)
}

//go:noescape
func MaxU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) Max(b Uint8x16) Uint8x16 {
	return MaxU8x16(a, b)
}

//go:noescape
func MaxU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) Max(b Uint8x32) Uint8x32 {
	return MaxU8x32(a, b)
}

//go:noescape
func MaxU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) Max(b Uint8x64) Uint8x64 {
	return MaxU8x64(a, b)
}

//go:noescape
func MinU8x16(a, b Uint8x16) Uint8x16

func (a Uint8x16) Min(b Uint8x16) Uint8x16 {
	return MinU8x16(a, b)
}

//go:noescape
func MinU8x32(a, b Uint8x32) Uint8x32

func (a Uint8x32) Min(b Uint8x32) Uint8x32 {
	return MinU8x32(a, b)
}

//go:noescape
func MinU8x64(a, b Uint8x64) Uint8x64

func (a Uint8x64) Min(b Uint8x64) Uint8x64 {
	return MinU8x64(a, b)
}
