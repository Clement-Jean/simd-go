#include <arm_neon.h>

void VdupqNU8(uint8x16_t* v1, uint8_t* v0) { *v1 = vdupq_n_u8(*v0); }
void VandqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vandq_u8(*v0, *v1); }
void VorrqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vorrq_u8(*v0, *v1); }
void VeorqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = veorq_u8(*v0, *v1); }
