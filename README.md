# simd-go

:warning: This is a proof of concept! :warning:

## Types

The following types are defined:

- Uint8x16
- Uint8x32
- Uint8x64
- Int8x16
- Int8x32
- Int8x64

## Functions

### Bitwise operations

- `func BitAndU8x16(a, b Uint8x16) Uint8x16`
- `func BitAndU8x32(a, b Uint8x32) Uint8x32`
- `func BitAndU8x64(a, b Uint8x64) Uint8x64`

- `func BitOrU8x16(a, b Uint8x16) Uint8x16`
- `func BitOrU8x32(a, b Uint8x32) Uint8x32`
- `func BitOrU8x64(a, b Uint8x64) Uint8x64`

- `func BitXorU8x16(a, b Uint8x16) Uint8x16`
- `func BitXorU8x32(a, b Uint8x32) Uint8x32`
- `func BitXorU8x64(a, b Uint8x64) Uint8x64`

### Arithmetics

- `func AddU8x16(a, b Uint8x16) Uint8x16`
- `func AddU8x32(a, b Uint8x32) Uint8x32`
- `func AddU8x64(a, b Uint8x64) Uint8x64`

- `func SubU8x16(a, b Uint8x16) Uint8x16`
- `func SubU8x32(a, b Uint8x32) Uint8x32`
- `func SubU8x64(a, b Uint8x64) Uint8x64`

- `func EqualU8x16(a, b Uint8x16) Uint8x16`
- `func EqualU8x32(a, b Uint8x32) Uint8x32`
- `func EqualU8x64(a, b Uint8x64) Uint8x64`

- `func MaxU8x16(a, b Uint8x16) Uint8x16`
- `func MaxU8x32(a, b Uint8x32) Uint8x32`
- `func MaxU8x64(a, b Uint8x64) Uint8x64`

- `func MinU8x16(a, b Uint8x16) Uint8x16`
- `func MinU8x32(a, b Uint8x32) Uint8x32`
- `func MinU8x64(a, b Uint8x64) Uint8x64`