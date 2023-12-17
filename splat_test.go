package simdgo_test

import (
	"testing"

	simdgo "github.com/Clement-Jean/simd-go"
)

func testSplatU8[T simdgo.Uint8x16 | simdgo.Uint8x32 | simdgo.Uint8x64](t *testing.T, got T, expected uint8) {
	t.Helper()

	for i := 0; i < len(got); i++ {
		if got[i] != expected {
			t.Fatalf("expected %d at index %d, got %d", expected, i, got[i])
		}
	}
}

func testSplatI8[T simdgo.Int8x16 | simdgo.Int8x32 | simdgo.Int8x64](t *testing.T, got T, expected int8) {
	t.Helper()

	for i := 0; i < len(got); i++ {
		if got[i] != expected {
			t.Fatalf("expected %d at index %d, got %d", expected, i, got[i])
		}
	}
}

func TestSplatU8x16(t *testing.T) {
	c := simdgo.SplatU8x16(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x16(t *testing.T) {
	c := simdgo.ZeroU8x16()

	testSplatU8(t, c, 0)
}

func TestSplatU8x32(t *testing.T) {
	c := simdgo.SplatU8x32(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x32(t *testing.T) {
	c := simdgo.ZeroU8x32()

	testSplatU8(t, c, 0)
}

func TestSplatU8x64(t *testing.T) {
	c := simdgo.SplatU8x64(1)

	testSplatU8(t, c, 1)
}

func TestZeroU8x64(t *testing.T) {
	c := simdgo.ZeroU8x64()

	testSplatU8(t, c, 0)
}

func TestSplatI8x16(t *testing.T) {
	c := simdgo.SplatI8x16(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x16(t *testing.T) {
	c := simdgo.ZeroI8x16()

	testSplatI8(t, c, 0)
}

func TestSplatI8x32(t *testing.T) {
	c := simdgo.SplatI8x32(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x32(t *testing.T) {
	c := simdgo.ZeroI8x32()

	testSplatI8(t, c, 0)
}

func TestSplatI8x64(t *testing.T) {
	c := simdgo.SplatI8x64(-1)

	testSplatI8(t, c, -1)
}

func TestZeroI8x64(t *testing.T) {
	c := simdgo.ZeroI8x64()

	testSplatI8(t, c, 0)
}
