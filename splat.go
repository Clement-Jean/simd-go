package simdgo

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatU8x16(a uint8) Uint8x16 {
	return Uint8x16{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a}
}

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatU8x32(a uint8) Uint8x32 {
	return Uint8x32{
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
	}
}

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatU8x64(a uint8) Uint8x64 {
	return Uint8x64{
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
	}
}

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatI8x16(a int8) Int8x16 {
	return Int8x16{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a}
}

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatI8x32(a int8) Int8x32 {
	return Int8x32{
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
	}
}

// Constructs a new SIMD vector with all elements set to the given value (a).
func SplatI8x64(a int8) Int8x64 {
	return Int8x64{
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
		a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a,
	}
}

// Constructs a new SIMD vector with all elements set to 0.
func ZeroU8x16() Uint8x16 { return SplatU8x16(0) }

// Constructs a new SIMD vector with all elements set to 0.
func ZeroU8x32() Uint8x32 { return SplatU8x32(0) }

// Constructs a new SIMD vector with all elements set to 0.
func ZeroU8x64() Uint8x64 { return SplatU8x64(0) }

// Constructs a new SIMD vector with all elements set to 0.
func ZeroI8x16() Int8x16 { return SplatI8x16(0) }

// Constructs a new SIMD vector with all elements set to 0.
func ZeroI8x32() Int8x32 { return SplatI8x32(0) }

// Constructs a new SIMD vector with all elements set to 0.
func ZeroI8x64() Int8x64 { return SplatI8x64(0) }
