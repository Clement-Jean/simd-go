package simdgo

//go:noescape
func SplatU8x16(a uint8) Uint8x16

//go:noescape
func SplatU8x32(a uint8) Uint8x32

//go:noescape
func SplatU8x64(a uint8) Uint8x64

//go:noescape
func SplatI8x16(a int8) Int8x16

//go:noescape
func SplatI8x32(a int8) Int8x32

//go:noescape
func SplatI8x64(a int8) Int8x64

func ZeroU8x16() Uint8x16 { return SplatU8x16(0) }

func ZeroU8x32() Uint8x32 { return SplatU8x32(0) }

func ZeroU8x64() Uint8x64 { return SplatU8x64(0) }

func ZeroI8x16() Int8x16 { return SplatI8x16(0) }

func ZeroI8x32() Int8x32 { return SplatI8x32(0) }

func ZeroI8x64() Int8x64 { return SplatI8x64(0) }
