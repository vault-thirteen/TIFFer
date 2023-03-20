package mn

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/vault-thirteen/TIFFerhelper"
	"github.com/vault-thirteen/TIFFermodels/ByteOrder"
	"github.com/vault-thirteen/TIFFermodels/basic-types"
)

const (
	Unknown  = MagicNumber(0)
	TIFF_6_0 = MagicNumber(42)
)

const MagicNumberSize = 2

const (
	ErrUnsupportedMagicNumber = "unsupported magic number: %v"
)

// MagicNumber is the TIFF 6.0 magic number.
type MagicNumber bt.Word

// New reads the magic number from the stream and returns it.
func New(r io.Reader, byteOrder bo.ByteOrder) (mn MagicNumber, err error) {
	var ba []byte
	ba, err = helper.ReadBytes(r, MagicNumberSize)
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

	return Unknown, fmt.Errorf(ErrUnsupportedMagicNumber, mn)
}
