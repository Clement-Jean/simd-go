package simdgo_test

import (
	"testing"

	simd "github.com/Clement-Jean/simd-go"
)

func testSplatU8[T simd.Uint8x16 | simd.Uint8x32 | simd.Uint8x64](t *testing.T, got T, expected uint8) {
	t.Helper()

	for i := 0; i < len(got); i++ {
		if got[i] != expected {
			t.Fatalf("expected %d at index %d, got %d", expected, i, got[i])
		}
	}
}

func testSplatI8[T simd.Int8x16 | simd.Int8x32 | simd.Int8x64](t *testing.T, got T, expected int8) {
	t.Helper()

	for i := 0; i < len(got); i++ {
		if got[i] != expected {
			t.Fatalf("expected %d at index %d, got %d", expected, i, got[i])
		}
	}
}

func TestSplatU8x16(t *testing.T) {
	c := simd.SplatU8x16(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x16(t *testing.T) {
	c := simd.ZeroU8x16()

	testSplatU8(t, c, 0)
}

func BenchmarkSplat16(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatU8x16(1)
	}
}

func TestSplatU8x32(t *testing.T) {
	c := simd.SplatU8x32(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x32(t *testing.T) {
	c := simd.ZeroU8x32()

	testSplatU8(t, c, 0)
}

func BenchmarkSplat32(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatU8x32(1)
	}
}

func TestSplatU8x64(t *testing.T) {
	c := simd.SplatU8x64(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x64(t *testing.T) {
	c := simd.ZeroU8x64()

	testSplatU8(t, c, 0)
}

func BenchmarkSplat64(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatU8x64(1)
	}
}

func TestSplatI8x16(t *testing.T) {
	c := simd.SplatI8x16(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x16(t *testing.T) {
	c := simd.ZeroI8x16()

	testSplatI8(t, c, 0)
}

func BenchmarkSplatI16(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatI8x16(1)
	}
}

func TestSplatI8x32(t *testing.T) {
	c := simd.SplatI8x32(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x32(t *testing.T) {
	c := simd.ZeroI8x32()

	testSplatI8(t, c, 0)
}

func BenchmarkSplatI32(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatI8x32(1)
	}
}

func TestSplatI8x64(t *testing.T) {
	c := simd.SplatI8x64(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x64(t *testing.T) {
	c := simd.ZeroI8x64()

	testSplatI8(t, c, 0)
}

func BenchmarkSplatI64(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simd.SplatI8x64(1)
	}
}
