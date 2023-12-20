package tests

import (
	"testing"
	"unsafe"

	simd "github.com/Clement-Jean/simd-go"
	"github.com/Clement-Jean/simd-go/tests/internal"
)

func ReferenceAnd(a, b simd.Uint8x16) simd.Uint8x16 {
	var result simd.Uint8x16

	internal.VandqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result[0])),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)

	return result
}

func ReferenceOr(a, b simd.Uint8x16) simd.Uint8x16 {
	var result simd.Uint8x16

	internal.VorrqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result[0])),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)

	return result
}

func ReferenceEor(a, b simd.Uint8x16) simd.Uint8x16 {
	var result simd.Uint8x16

	internal.VeorqU8(
		(*internal.Uint8x16)(unsafe.Pointer(&result[0])),
		(*internal.Uint8x16)(unsafe.Pointer(&a)),
		(*internal.Uint8x16)(unsafe.Pointer(&b)),
	)

	return result
}

func FuzzAnd(f *testing.F) {
	f.Add(uint8(1), uint8(0))
	f.Add(uint8(0), uint8(1))
	f.Add(uint8(1), uint8(1))

	f.Fuzz(func(t *testing.T, a, b uint8) {
		aItems := simd.SplatU8x16(a)
		bItems := simd.SplatU8x16(b)
		got := ReferenceAnd(aItems, bItems)
		expected := simd.BitAndU8x16(aItems, bItems)

		if len(expected) != len(got) {
			t.Fatalf("expected len %d, got %d", len(expected), len(got))
		}

		t.Logf("%v\n", expected)
		t.Logf("%v\n", got)
		for i := 0; i < len(got); i++ {
			if expected[i] != uint8(got[i]) {
				t.Fail()
			}
		}
	})
}

func FuzzOr(f *testing.F) {
	f.Add(uint8(1), uint8(0))
	f.Add(uint8(0), uint8(1))
	f.Add(uint8(1), uint8(1))

	f.Fuzz(func(t *testing.T, a, b uint8) {
		aItems := simd.SplatU8x16(a)
		bItems := simd.SplatU8x16(b)
		got := ReferenceOr(aItems, bItems)
		expected := simd.BitOrU8x16(aItems, bItems)

		if len(expected) != len(got) {
			t.Fatalf("expected len %d, got %d", len(expected), len(got))
		}

		t.Logf("%v\n", expected)
		t.Logf("%v\n", got)
		for i := 0; i < len(got); i++ {
			if expected[i] != uint8(got[i]) {
				t.Fail()
			}
		}
	})
}

func FuzzXor(f *testing.F) {
	f.Add(uint8(1), uint8(0))
	f.Add(uint8(0), uint8(1))
	f.Add(uint8(1), uint8(1))

	f.Fuzz(func(t *testing.T, a, b uint8) {
		aItems := simd.SplatU8x16(a)
		bItems := simd.SplatU8x16(b)
		got := ReferenceEor(aItems, bItems)
		expected := simd.BitXorU8x16(aItems, bItems)

		if len(expected) != len(got) {
			t.Fatalf("expected len %d, got %d", len(expected), len(got))
		}

		t.Logf("%v\n", expected)
		t.Logf("%v\n", got)
		for i := 0; i < len(got); i++ {
			if expected[i] != uint8(got[i]) {
				t.Fail()
			}
		}
	})
}
