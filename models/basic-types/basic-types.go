package bt

import "math/big"

// This file is placed here purely for the reason of documenting the library.
// Please, do read all the notes below. It may be important.

// Byte is a byte type.
type Byte = byte

// Short type.
// Well ... Normally, Short is the C language type holding a signed 16-bit
// integer number, but... TIFF format does everything opposite to the rest of
// the world. In TIFF, Short means an unsigned 16-bit integer number. Ha-ha.
// https://en.wikipedia.org/wiki/C_data_types We do not care who and how
// creates these weird names, we prefer Microsoft's names, written below.
type Short = uint16

// Long type.
// The same as with above ... Normally, Long is the C language type holding a
// signed 32-bit integer number, but... TIFF format does everything opposite to
// the rest of the world. In TIFF, Long means an unsigned 32-bit integer number.
// https://en.wikipedia.org/wiki/C_data_types We do not care who and how
// creates these weird names, we prefer Microsoft's names, written below.
type Long = uint32

// SByte means a Signed Byte type.
type SByte = int8

// SShort means a Signed Short type.
// The reason is stated above.
type SShort = int16

// SLong means a Signed Long type.
// The reason is stated above.
type SLong = int32

// Float is a 32-bit floating point number.
type Float = float32

// Double is a 64-bit floating point number.
type Double = float64

// Rational is an unsigned rational type.
// At least, it should be unsigned by the TIFF 6.0 Specification.
// Unfortunately Golang does not support unsigned rationals out-of-the-box.
type Rational = *big.Rat

// SRational is a signed rational type.
type SRational = *big.Rat

// Notes on the rational types.
//
// Golang provides very limited support for rational numbers:
//
//	1.	Rational numbers are always signed, i.e. only signed integer numerators
//		and denominators are supported.
//
//	2.	Numerators and denominators use 64-bit integer numbers, i.e. 32-bit
//		versions of rational numbers are not supported.

// Word is the word type.
// It is a 16-bit unsigned integer.
// WORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f8573df3-a44a-4a50-b070-ac4c3aa78e3c
type Word = uint16

// DWord is the double word type.
// It is a 32-bit unsigned integer.
// DWORD type in the open specification by Microsoft:
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/262627d8-3418-4627-9218-4ffe110850b2
type DWord = uint32

// War. War never changes.
// The war of formats, names and everything else is always the same.
// Adobe wanted to rule the world, but it turned out to be impossible.
