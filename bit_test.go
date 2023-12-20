package simdgo_test

import (
	"testing"

	simd "github.com/Clement-Jean/simd-go"
)

func TestBitAndU8x16(t *testing.T) {
	a := simd.SplatU8x16(1)
	b := simd.SplatU8x16(0)
	c := simd.BitAndU8x16(a, b)

	for _, d := range c {
		if d != 0 {
			t.Fatalf("expected 0, got %d", d)
		}
	}
}

func TestBitAndU8x32(t *testing.T) {
	a := simd.SplatU8x32(1)
	b := simd.SplatU8x32(0)
	c := simd.BitAndU8x32(a, b)

	for _, d := range c {
		if d != 0 {
			t.Fatalf("expected 0, got %d", d)
		}
	}
}

func TestBitAndU8x64(t *testing.T) {
	a := simd.SplatU8x64(1)
	b := simd.SplatU8x64(0)
	c := simd.BitAndU8x64(a, b)

	for _, d := range c {
		if d != 0 {
			t.Fatalf("expected 0, got %d", d)
		}
	}
}

func TestBitOrU8x16(t *testing.T) {
	a := simd.SplatU8x16(1)
	b := simd.SplatU8x16(0)
	c := simd.BitOrU8x16(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}

func TestBitXorU8x16(t *testing.T) {
	a := simd.SplatU8x16(1)
	b := simd.SplatU8x16(0)
	c := simd.BitXorU8x16(a, b)

	for _, d := range c {
		if d != 1 {
			t.Fatalf("expected 1, got %d", d)
		}
	}
}
