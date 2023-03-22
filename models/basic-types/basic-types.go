package bt

import (
	bt "github.com/vault-thirteen/auxie/BasicTypes"
)

// Mapping of common sense types to the types used by TIFF 6.0 Specification.
// Some types are used for convenience.

// This file is placed here purely for the reason of documenting the library.
// Please, do read all the notes below. It may be important.

// Byte types.
type Byte = bt.Byte
type SByte = bt.SByte
type ASCII = bt.Byte
type Undefined = bt.Byte

// Short types.
// Well ... Normally, Short is the C language type holding a signed 16-bit
// integer number, but... TIFF format does everything opposite to the rest of
// the world. In TIFF, Short means an unsigned 16-bit integer number. Ha-ha.
// https://en.wikipedia.org/wiki/C_data_types We do not care who and how
// creates these weird names, we prefer Microsoft's names, written below.
type Short = bt.UShort
type SShort = bt.SShort
type Word = bt.Word

// Long types.
// The same as with above ... Normally, Long is the C language type holding a
// signed 32-bit integer number, but... TIFF format does everything opposite to
// the rest of the world. In TIFF, Long means an unsigned 32-bit integer number.
// https://en.wikipedia.org/wiki/C_data_types We do not care who and how
// creates these weird names, we prefer Microsoft's names, written below.
type Long = bt.ULong
type SLong = bt.SLong
type DWord = bt.DWord

// Rational types.
// At least, it should be unsigned by the TIFF 6.0 Specification.
// Unfortunately Golang does not support unsigned rationals out-of-the-box.
// Rational numbers in TIFF consist of two long numbers, where the first one is
// a long numerator and the second one is a long denominator. Rational type
// should use unsigned longs and signed rational type should use signed longs.
type Rational = bt.Rational
type SRational = bt.Rational

// Notes on the rational types.
//
// Golang provides very limited support for rational numbers:
//
//	1.	Rational numbers are always signed, i.e. only signed integer numerators
//		and denominators are supported.
//
//	2.	Numerators and denominators use 64-bit integer numbers, i.e. 32-bit
//		versions of rational numbers are not supported.

// Float or floating point types.
type Float = bt.Float
type Double = bt.Double

// War. War never changes.
// The war of formats, names and everything else is always the same.
// Adobe wanted to rule the world, but it turned out to be impossible.
