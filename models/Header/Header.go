package hdr

import (
	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/rs"

	"github.com/vault-thirteen/TIFFer/models"
	"github.com/vault-thirteen/TIFFer/models/ByteOrder"
	ifd "github.com/vault-thirteen/TIFFer/models/IFD"
	"github.com/vault-thirteen/TIFFer/models/MagicNumber"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

// Header is the Image File Header described in the TIFF 6.0 Specification.
type Header struct {
	// ByteOrder is the byte order, used for encoding the TIFF.
	ByteOrder bo.ByteOrder

	// MagicNumber is always 42 for TIFF 6.0.
	MagicNumber mn.MagicNumber

	// OffsetOfFirstIFD is an offset of the first IFD.
	OffsetOfFirstIFD models.OffsetOfIFD

	// FirstIFD is a pointer to the first IFD.
	FirstIFD *ifd.IFD
}

// New constructs the Header from the reader.
func New(rs *rs.ReaderSeeker) (h *Header, err error) {
	h = &Header{}

	// Byte order.
	h.ByteOrder, err = bo.New(rs)
	if err != nil {
		return nil, err
	}

	// Magic number.
	h.MagicNumber, err = mn.New(rs, h.ByteOrder)
	if err != nil {
		return nil, err
	}

	// OffsetOfValue of the first IFD.
	h.OffsetOfFirstIFD, err = h.readIFDOffset(rs, h.ByteOrder)
	if err != nil {
		return nil, err
	}

	return h, nil
}

// readIFDOffset reads the IFD offset and returns it.
func (h *Header) readIFDOffset(rs *rs.ReaderSeeker, byteOrder bo.ByteOrder) (ifdOffset bt.DWord, err error) {
	switch byteOrder {
	case bo.BigEndian:
		return rs.ReadDWord_BE()
	case bo.LittleEndian:
		return rs.ReadDWord_LE()
	default:
		return 0, errors.ErrorFromFormattedText(bo.Errf_UnsupportedBO, byteOrder)
	}
}
