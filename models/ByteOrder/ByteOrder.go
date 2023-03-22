package bo

import (
	"fmt"

	"github.com/vault-thirteen/auxie/reader"
)

const (
	Unknown      = ByteOrder(0)
	BigEndian    = ByteOrder(1)
	LittleEndian = ByteOrder(2)
)

const ByteOrderMarkSize = 2

const (
	ByteI = 'I'
	ByteM = 'M'
)

const (
	ErrUnsupportedBOM = "unsupported byte order mark: %v"
	ErrUnsupportedBO  = "unsupported byte order: %v"
)

// ByteOrder is the byte order.
// It can be either big endian or little endian.
type ByteOrder byte

// New reads the byte order from the stream and returns it.
func New(rs *reader.Reader) (bo ByteOrder, err error) {
	var ba []byte
	ba, err = rs.ReadBytes(ByteOrderMarkSize)
	if err != nil {
		return Unknown, err
	}

	if (ba[0] == ByteI) && (ba[1] == ByteI) {
		return LittleEndian, nil
	} else if (ba[0] == ByteM) && (ba[1] == ByteM) {
		return BigEndian, nil
	}

	return Unknown, fmt.Errorf(ErrUnsupportedBOM, ba)
}
