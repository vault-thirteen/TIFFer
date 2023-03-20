package helper

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	bt "github.com/vault-thirteen/TIFFer/models/basic-types"
)

const ErrUnexpectedDataSize = "unexpected data size: %v vs %v"

// ReadBytes is the helper function to read an exact number of bytes from the
// stream.
func ReadBytes(r io.Reader, bytesCount int) (bytes []byte, err error) {
	bytes = make([]byte, bytesCount)
	var n int
	n, err = io.ReadFull(r, bytes)
	if n != bytesCount {
		return bytes, fmt.Errorf(ErrUnexpectedDataSize, bytesCount, n)
	}
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}

// ReadByte reads one byte and returns it.
func ReadByte(r io.Reader) (b byte, err error) {
	var bytes []byte
	bytes, err = ReadBytes(r, 1)
	if err != nil {
		return b, err
	}

	return bytes[0], nil
}

// Read2Bytes reads two bytes and returns them.
func Read2Bytes(r io.Reader) (bytes []byte, err error) {
	return ReadBytes(r, 2)
}

// Read4Bytes reads four bytes and returns them.
func Read4Bytes(r io.Reader) (bytes []byte, err error) {
	return ReadBytes(r, 4)
}

// Read8Bytes reads eight bytes and returns them.
func Read8Bytes(r io.Reader) (bytes []byte, err error) {
	return ReadBytes(r, 8)
}

// ReadWord_BE reads a word using the big endian technique and returns it.
func ReadWord_BE(r io.Reader) (w bt.Word, err error) {
	var bytes []byte
	bytes, err = Read2Bytes(r)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(bytes), nil
}

// ReadWord_LE reads a word using the little endian technique and returns it.
func ReadWord_LE(r io.Reader) (w bt.Word, err error) {
	var bytes []byte
	bytes, err = Read2Bytes(r)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(bytes), nil
}

// ReadDWord_BE reads a double word using the big endian technique and returns
// it.
func ReadDWord_BE(r io.Reader) (dw bt.DWord, err error) {
	var bytes []byte
	bytes, err = Read4Bytes(r)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(bytes), nil
}

// ReadDWord_LE reads a double word using the little endian technique and
// returns it.
func ReadDWord_LE(r io.Reader) (dw bt.DWord, err error) {
	var bytes []byte
	bytes, err = Read4Bytes(r)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(bytes), nil
}

// ReadASCII reads an ASCII byte and returns it.
func ReadASCII(r io.Reader) (b byte, err error) {
	return ReadByte(r)
}

// ReadShort_BE reads a Short using the big endian technique and returns it.
func ReadShort_BE(r io.Reader) (w bt.Word, err error) {
	return ReadWord_BE(r)
}

// ReadShort_LE reads a Short using the little endian technique and returns it.
func ReadShort_LE(r io.Reader) (w bt.Word, err error) {
	return ReadWord_LE(r)
}

// ReadLong_BE reads a Long using the big endian technique and returns it.
func ReadLong_BE(r io.Reader) (dw bt.DWord, err error) {
	return ReadDWord_BE(r)
}

// ReadLong_LE reads a Long using the little endian technique and returns it.
func ReadLong_LE(r io.Reader) (dw bt.DWord, err error) {
	return ReadDWord_LE(r)
}

// ReadRational_BE reads a Rational using the big endian technique and returns
// it.
// This method is not fully compatible with TIFF 6.0 Specification.
// More information can be found in the comments to the 'models.Type' type.
func ReadRational_BE(r io.Reader) (rat *big.Rat, err error) {
	return ReadSRational_BE(r)
}

// ReadRational_LE reads a Rational using the little endian technique and
// returns it.
// This method is not fully compatible with TIFF 6.0 Specification.
// More information can be found in the comments to the 'models.Type' type.
func ReadRational_LE(r io.Reader) (rat *big.Rat, err error) {
	return ReadSRational_LE(r)
}

// ReadSByte reads one signed byte and returns it.
func ReadSByte(r io.Reader) (sb int8, err error) {
	var bytes []byte
	bytes, err = ReadBytes(r, 1)
	if err != nil {
		return sb, err
	}

	return int8(bytes[0]), nil
}

// ReadUndefined reads Undefined type item and returns it.
func ReadUndefined(r io.Reader) (b byte, err error) {
	return ReadByte(r)
}

// ReadSShort_BE reads an SShort using the big endian technique and returns it.
func ReadSShort_BE(r io.Reader) (ss int16, err error) {
	var w bt.Word
	w, err = ReadWord_BE(r)
	if err != nil {
		return ss, err
	}

	return int16(w), nil
}

// ReadSShort_LE reads an SShort using the little endian technique and returns
// it.
func ReadSShort_LE(r io.Reader) (ss int16, err error) {
	var w bt.Word
	w, err = ReadWord_LE(r)
	if err != nil {
		return ss, err
	}

	return int16(w), nil
}

// ReadSLong_BE reads an SLong using the big endian technique and returns it.
func ReadSLong_BE(r io.Reader) (sl int32, err error) {
	var dw bt.DWord
	dw, err = ReadDWord_BE(r)
	if err != nil {
		return sl, err
	}

	return int32(dw), nil
}

// ReadSLong_LE reads an SLong using the little endian technique and returns it.
func ReadSLong_LE(r io.Reader) (sl int32, err error) {
	var dw bt.DWord
	dw, err = ReadDWord_LE(r)
	if err != nil {
		return sl, err
	}

	return int32(dw), nil
}

// ReadSRational_BE reads an SRational using the big endian technique and returns
// it.
// This method is not fully compatible with TIFF 6.0 Specification.
// More information can be found in the comments to the 'models.Type' type.
func ReadSRational_BE(r io.Reader) (rat *big.Rat, err error) {
	var numerator bt.DWord
	numerator, err = ReadDWord_BE(r)
	if err != nil {
		return rat, err
	}

	var denominator bt.DWord
	denominator, err = ReadDWord_BE(r)
	if err != nil {
		return rat, err
	}

	return big.NewRat(int64(numerator), int64(denominator)), nil
}

// ReadSRational_LE reads an SRational using the little endian technique and
// returns it.
// This method is not fully compatible with TIFF 6.0 Specification.
// More information can be found in the comments to the 'models.Type' type.
func ReadSRational_LE(r io.Reader) (rat *big.Rat, err error) {
	var numerator bt.DWord
	numerator, err = ReadDWord_LE(r)
	if err != nil {
		return rat, err
	}

	var denominator bt.DWord
	denominator, err = ReadDWord_LE(r)
	if err != nil {
		return rat, err
	}

	return big.NewRat(int64(numerator), int64(denominator)), nil
}

// ReadFloat_BE reads a Float using the big endian technique and returns it.
func ReadFloat_BE(r io.Reader) (f float32, err error) {
	err = binary.Read(r, binary.BigEndian, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// ReadFloat_LE reads a Float using the little endian technique and returns it.
func ReadFloat_LE(r io.Reader) (f float32, err error) {
	err = binary.Read(r, binary.LittleEndian, &f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// ReadDouble_BE reads a Double using the big endian technique and returns it.
func ReadDouble_BE(r io.Reader) (d float64, err error) {
	err = binary.Read(r, binary.BigEndian, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

// ReadDouble_LE reads a Double using the little endian technique and returns it.
func ReadDouble_LE(r io.Reader) (d float64, err error) {
	err = binary.Read(r, binary.LittleEndian, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}
