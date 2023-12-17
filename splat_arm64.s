#include "textflag.h"

// func SplatU8x16(a uint8) Uint8x16
TEXT ·SplatU8x16(SB),NOSPLIT,$0-24
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	RET

// func SplatU8x32(a uint8) Uint8x32
TEXT ·SplatU8x32(SB),NOSPLIT,$0-40
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	MOVD R0, ret+24(FP)
	MOVD R0, ret+32(FP)
	RET

// func SplatU8x64(a uint8) Uint8x64
TEXT ·SplatU8x64(SB),NOSPLIT,$0-72
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	MOVD R0, ret+24(FP)
	MOVD R0, ret+32(FP)
	MOVD R0, ret+40(FP)
	MOVD R0, ret+48(FP)
	MOVD R0, ret+56(FP)
	MOVD R0, ret+64(FP)
	RET

// func SplatI8x16(a int8) Int8x16
TEXT ·SplatI8x16(SB),NOSPLIT,$0-24
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	RET

// func SplatI8x32(a int8) Int8x32
TEXT ·SplatI8x32(SB),NOSPLIT,$0-40
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	MOVD R0, ret+24(FP)
	MOVD R0, ret+32(FP)
	RET

// func SplatI8x64(a int8) Int8x64
TEXT ·SplatI8x64(SB),NOSPLIT,$0-72
	MOVB a+0(FP), R0
	VDUP R0, V0.B16
	VMOV V0.D[0], R0
	MOVD R0, ret+8(FP)
	MOVD R0, ret+16(FP)
	MOVD R0, ret+24(FP)
	MOVD R0, ret+32(FP)
	MOVD R0, ret+40(FP)
	MOVD R0, ret+48(FP)
	MOVD R0, ret+56(FP)
	MOVD R0, ret+64(FP)
	RET
