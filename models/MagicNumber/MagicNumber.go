package mn

import (
	"encoding/binary"

	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/rs"

	"github.com/vault-thirteen/TIFFer/models/ByteOrder"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

const (
	Unknown  = MagicNumber(0)
	TIFF_6_0 = MagicNumber(42)
)

const MagicNumberSize = 2

const (
	Errf_UnsupportedMagicNumber = "unsupported magic number: %v"
)

// MagicNumber is the TIFF 6.0 magic number.
type MagicNumber bt.Word

// New reads the magic number from the stream and returns it.
func New(rs *rs.ReaderSeeker, byteOrder bo.ByteOrder) (mn MagicNumber, err error) {
	var ba []byte
	ba, err = rs.ReadBytes(MagicNumberSize)
	if err != nil {
		return Unknown, err
	}

	switch byteOrder {
	case bo.BigEndian:
		mn = MagicNumber(binary.BigEndian.Uint16(ba))

	case bo.LittleEndian:
		mn = MagicNumber(binary.LittleEndian.Uint16(ba))
	}

	switch mn {
	case TIFF_6_0:
		return mn, nil
	}

	return Unknown, errors.ErrorFromFormattedText(Errf_UnsupportedMagicNumber, mn)
}
