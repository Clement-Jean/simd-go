package internal

/*
#include <arm_neon.h>
*/
import "C"

// typedef __attribute__((neon_vector_type(16))) uint8_t uint8x16_t;
type Uint8x16 = C.uint8x16_t

// typedef uchar uint8_t;
type Uint8 = C.uint8_t

//go:linkname VdupqNU8 VdupqNU8
//go:noescape
func VdupqNU8(r *Uint8x16, v0 *Uint8)

//go:linkname VandqU8 VandqU8
//go:noescape
func VandqU8(r *Uint8x16, v0 *Uint8x16, v1 *Uint8x16)

//go:linkname VorrqU8 VorrqU8
//go:noescape
func VorrqU8(r *Uint8x16, v0 *Uint8x16, v1 *Uint8x16)

//go:linkname VeorqU8 VeorqU8
//go:noescape
func VeorqU8(r *Uint8x16, v0 *Uint8x16, v1 *Uint8x16)
