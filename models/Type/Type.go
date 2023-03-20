package t

// The word 'type' is reserved by Go language.
// Well, you reap what you sow. Ha-ha.

import bt "github.com/vault-thirteen/TIFFer/models/basic-types"

const (
	Byte      = 1  // Byte.
	ASCII     = 2  // Byte.
	Short     = 3  // WORD, uint16, 2 Bytes.
	Long      = 4  // DWORD, uint32, 4 Bytes.
	Rational  = 5  // Two DWORD, two uint32, 2x4 Bytes.
	SByte     = 6  // Byte.
	Undefined = 7  // Byte.
	SShort    = 8  // int16, 2 Bytes.
	SLong     = 9  // int32, 4 Bytes.
	SRational = 10 // Two DWORD, two int32, 2x4 Bytes.
	Float     = 11 // float32, 4 Bytes.
	Double    = 12 // float64, 8 Bytes.
)

const ErrUnknownType = "unknown data item type: %v"

// Type is the type of data items of a Directory Entry.
type Type = bt.Word
