package hdr

import (
	"fmt"
	"io"

	"github.com/vault-thirteen/TIFFer/helper"
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
func New(r io.Reader) (h *Header, err error) {
	h = &Header{}

	// Byte order.
	h.ByteOrder, err = bo.New(r)
	if err != nil {
		return nil, err
	}

	// Magic number.
	h.MagicNumber, err = mn.New(r, h.ByteOrder)
	if err != nil {
		return nil, err
	}

	// OffsetOfValue of the first IFD.
	h.OffsetOfFirstIFD, err = h.readIFDOffset(r, h.ByteOrder)
	if err != nil {
		return nil, err
	}

	return h, nil
}

// readIFDOffset reads the IFD offset and returns it.
func (h *Header) readIFDOffset(r io.Reader, byteOrder bo.ByteOrder) (ifdOffset bt.DWord, err error) {
	switch byteOrder {
	case bo.BigEndian:
		return helper.ReadDWord_BE(r)
	case bo.LittleEndian:
		return helper.ReadDWord_LE(r)
	default:
		return 0, fmt.Errorf(bo.ErrUnsupportedBO, byteOrder)
	}
}
