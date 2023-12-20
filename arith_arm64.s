#include "textflag.h"

// func AddU8x16(a, b Uint8x16) Uint8x16
TEXT ·AddU8x16(SB),NOSPLIT,$0-48
	MOVD $a+0(FP), R0
	VLD1.P 32(R0), [V0.D2, V1.D2]
	VADD V1.B16, V0.B16, V0.B16
	MOVD $ret+32(FP), R0
	VST1.P [V0.D2], (R0)
	RET

// func AddU8x32(a, b Uint8x32) Uint8x32
TEXT ·AddU8x32(SB),NOSPLIT,$0-96
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VADD V2.B16, V0.B16, V0.B16
	VADD V3.B16, V1.B16, V1.B16
	MOVD $ret+64(FP), R0
	VST1.P [V0.D2, V1.D2], (R0)
	RET

	// func AddU8x164a, b Uint8x64) Uint8x64
TEXT ·AddU8x64(SB),NOSPLIT,$0-192
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VLD1.P 64(R0), [V4.D2, V5.D2, V6.D2, V7.D2]
	VADD V4.B16, V0.B16, V0.B16
	VADD V5.B16, V1.B16, V1.B16
	VADD V6.B16, V2.B16, V2.B16
	VADD V7.B16, V3.B16, V3.B16
	MOVD $ret+128(FP), R0
	VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R0)
	RET

// func SubU8x16(a, b Uint8x16) Uint8x16
TEXT ·SubU8x16(SB),NOSPLIT,$0-48
	MOVD $a+0(FP), R0
	VLD1.P 32(R0), [V0.D2, V1.D2]
	VSUB V1.B16, V0.B16, V0.B16
	MOVD $ret+32(FP), R0
	VST1.P [V0.D2], (R0)
	RET

// func SubU8x32(a, b Uint8x32) Uint8x32
TEXT ·SubU8x32(SB),NOSPLIT,$0-96
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VSUB V2.B16, V0.B16, V0.B16
	VSUB V3.B16, V1.B16, V1.B16
	MOVD $ret+64(FP), R0
	VST1.P [V0.D2, V1.D2], (R0)
	RET

	// func SubU8x164a, b Uint8x64) Uint8x64
TEXT ·SubU8x64(SB),NOSPLIT,$0-192
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VLD1.P 64(R0), [V4.D2, V5.D2, V6.D2, V7.D2]
	VSUB V4.B16, V0.B16, V0.B16
	VSUB V5.B16, V1.B16, V1.B16
	VSUB V6.B16, V2.B16, V2.B16
	VSUB V7.B16, V3.B16, V3.B16
	MOVD $ret+128(FP), R0
	VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R0)
	RET

// func EqualU8x16(a, b Uint8x16) Uint8x16
TEXT ·EqualU8x16(SB),NOSPLIT,$0-48
	MOVD $a+0(FP), R0
	VLD1.P 32(R0), [V0.D2, V1.D2]
	VCMEQ V1.B16, V0.B16, V0.B16
	MOVD $ret+32(FP), R0
	VST1.P [V0.D2], (R0)
	RET

// func EqualU8x32(a, b Uint8x32) Uint8x32
TEXT ·EqualU8x32(SB),NOSPLIT,$0-96
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VCMEQ V2.B16, V0.B16, V0.B16
	VCMEQ V3.B16, V1.B16, V1.B16
	MOVD $ret+64(FP), R0
	VST1.P [V0.D2, V1.D2], (R0)
	RET

	// func EqualU8x164a, b Uint8x64) Uint8x64
TEXT ·EqualU8x64(SB),NOSPLIT,$0-192
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VLD1.P 64(R0), [V4.D2, V5.D2, V6.D2, V7.D2]
	VCMEQ V4.B16, V0.B16, V0.B16
	VCMEQ V5.B16, V1.B16, V1.B16
	VCMEQ V6.B16, V2.B16, V2.B16
	VCMEQ V7.B16, V3.B16, V3.B16
	MOVD $ret+128(FP), R0
	VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R0)
	RET

// func MaxU8x16(a, b Uint8x16) Uint8x16
TEXT ·MaxU8x16(SB),NOSPLIT,$0-48
	MOVD $a+0(FP), R0
	VLD1.P 32(R0), [V0.D2, V1.D2]
	VUMAX V1.B16, V0.B16, V0.B16
	MOVD $ret+32(FP), R0
	VST1.P [V0.D2], (R0)
	RET

// func MaxU8x32(a, b Uint8x32) Uint8x32
TEXT ·MaxU8x32(SB),NOSPLIT,$0-96
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VUMAX V2.B16, V0.B16, V0.B16
	VUMAX V3.B16, V1.B16, V1.B16
	MOVD $ret+64(FP), R0
	VST1.P [V0.D2, V1.D2], (R0)
	RET

	// func MaxU8x64(a, b Uint8x64) Uint8x64
TEXT ·MaxU8x64(SB),NOSPLIT,$0-192
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VLD1.P 64(R0), [V4.D2, V5.D2, V6.D2, V7.D2]
	VUMAX V4.B16, V0.B16, V0.B16
	VUMAX V5.B16, V1.B16, V1.B16
	VUMAX V6.B16, V2.B16, V2.B16
	VUMAX V7.B16, V3.B16, V3.B16
	MOVD $ret+128(FP), R0
	VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R0)
	RET

// func MinU8x16(a, b Uint8x16) Uint8x16
TEXT ·MinU8x16(SB),NOSPLIT,$0-48
	MOVD $a+0(FP), R0
	VLD1.P 32(R0), [V0.D2, V1.D2]
	VUMIN V1.B16, V0.B16, V0.B16
	MOVD $ret+32(FP), R0
	VST1.P [V0.D2], (R0)
	RET

// func MinU8x32(a, b Uint8x32) Uint8x32
TEXT ·MinU8x32(SB),NOSPLIT,$0-96
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VUMIN V2.B16, V0.B16, V0.B16
	VUMIN V3.B16, V1.B16, V1.B16
	MOVD $ret+64(FP), R0
	VST1.P [V0.D2, V1.D2], (R0)
	RET

	// func MinU8x64(a, b Uint8x64) Uint8x64
TEXT ·MinU8x64(SB),NOSPLIT,$0-192
	MOVD $a+0(FP), R0
	VLD1.P 64(R0), [V0.D2, V1.D2, V2.D2, V3.D2]
	VLD1.P 64(R0), [V4.D2, V5.D2, V6.D2, V7.D2]
	VUMIN V4.B16, V0.B16, V0.B16
	VUMIN V5.B16, V1.B16, V1.B16
	VUMIN V6.B16, V2.B16, V2.B16
	VUMIN V7.B16, V3.B16, V3.B16
	MOVD $ret+128(FP), R0
	VST1.P [V0.D2, V1.D2, V2.D2, V3.D2], (R0)
	RET
