package simdgo_test

import (
	"testing"

	simd "github.com/Clement-Jean/simd-go"
)

func TestAddU8x16(t *testing.T) {
	a := simd.SplatU8x16(2)
	b := simd.SplatU8x16(255)
	c := simd.AddU8x16(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}

func BenchmarkVectorAdd16(b *testing.B) {
	x := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.AddU8x16(x, y)
	}
}

func TestAddU8x32(t *testing.T) {
	a := simd.SplatU8x32(2)
	b := simd.SplatU8x32(255)
	c := simd.AddU8x32(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}

func BenchmarkVectorAdd32(b *testing.B) {
	x := simd.Uint8x32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.AddU8x32(x, y)
	}
}

func TestAddU8x64(t *testing.T) {
	a := simd.SplatU8x64(2)
	b := simd.SplatU8x64(255)
	c := simd.AddU8x64(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}

func BenchmarkVectorAdd64(b *testing.B) {
	x := simd.Uint8x64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.AddU8x64(x, y)
	}
}

func TestSubU8x16(t *testing.T) {
	a := simd.SplatU8x16(255)
	b := simd.SplatU8x16(2)
	c := simd.SubU8x16(a, b)

	for _, d := range c {
		if d != 253 {
			t.Fatalf("expected 253, got %d", d)
		}
	}
}

func BenchmarkVectorSub16(b *testing.B) {
	x := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SubU8x16(x, y)
	}
}

func TestSubU8x32(t *testing.T) {
	a := simd.SplatU8x32(255)
	b := simd.SplatU8x32(2)
	c := simd.SubU8x32(a, b)

	for _, d := range c {
		if d != 253 {
			t.Fatalf("expected 253, got %d", d)
		}
	}
}

func BenchmarkVectorSub32(b *testing.B) {
	x := simd.Uint8x32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SubU8x32(x, y)
	}
}

func TestSubU8x64(t *testing.T) {
	a := simd.SplatU8x64(255)
	b := simd.SplatU8x64(2)
	c := simd.SubU8x64(a, b)

	for _, d := range c {
		if d != 253 {
			t.Fatalf("expected 253, got %d", d)
		}
	}
}

func BenchmarkVectorSub64(b *testing.B) {
	x := simd.Uint8x64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SubU8x64(x, y)
	}
}

func TestEqualU8x16(t *testing.T) {
	a := simd.SplatU8x16(2)
	b := simd.SplatU8x16(2)
	c := simd.EqualU8x16(a, b)

	for _, d := range c {
		if d != 255 {
			t.Fatalf("expected 255, got %d", d)
		}
	}
}

func BenchmarkVectorEqual16(b *testing.B) {
	x := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	y := simd.Uint8x16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.EqualU8x16(x, y)
	}
}

func TestMaxU8x16(t *testing.T) {
	a := simd.SplatU8x16(1)
	b := simd.SplatU8x16(2)
	c := simd.MaxU8x16(a, b)

	for _, d := range c {
		if d != 2 {
			t.Fatalf("expected 2, got %d", d)
		}
	}
}

func TestMinU8x16(t *testing.T) {
	a := simd.SplatU8x16(1)
	b := simd.SplatU8x16(2)
	c := simd.MinU8x16(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}
