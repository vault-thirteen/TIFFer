package helper

import (
	"math/big"

	bt "github.com/vault-thirteen/TIFFer/models/basic-types"
	"github.com/vault-thirteen/auxie/reader"
)

// ReadASCII reads an ASCII byte.
func ReadASCII(rs *reader.Reader) (b byte, err error) {
	return rs.ReadByte()
}

// ReadUndefined reads Undefined type item.
func ReadUndefined(rs *reader.Reader) (b byte, err error) {
	return rs.ReadByte()
}

// ReadRational_BE reads a Rational using the big endian technique.
// This method is not fully compatible with TIFF 6.0 Specification.
func ReadRational_BE(rs *reader.Reader) (rat *big.Rat, err error) {
	return ReadSRational_BE(rs)
}

// ReadRational_LE reads a Rational using the little endian technique.
// This method is not fully compatible with TIFF 6.0 Specification.
func ReadRational_LE(rs *reader.Reader) (rat *big.Rat, err error) {
	return ReadSRational_LE(rs)
}

// ReadSRational_BE reads an SRational using the big endian technique.
// This method is not fully compatible with TIFF 6.0 Specification.
func ReadSRational_BE(rs *reader.Reader) (rat *big.Rat, err error) {
	var numerator bt.DWord
	numerator, err = rs.ReadDWord_BE()
	if err != nil {
		return rat, err
	}

	var denominator bt.DWord
	denominator, err = rs.ReadDWord_BE()
	if err != nil {
		return rat, err
	}

	return big.NewRat(int64(numerator), int64(denominator)), nil
}

// ReadSRational_LE reads an SRational using the little endian technique.
// This method is not fully compatible with TIFF 6.0 Specification.
func ReadSRational_LE(rs *reader.Reader) (rat *big.Rat, err error) {
	var numerator bt.DWord
	numerator, err = rs.ReadDWord_LE()
	if err != nil {
		return rat, err
	}

	var denominator bt.DWord
	denominator, err = rs.ReadDWord_LE()
	if err != nil {
		return rat, err
	}

	return big.NewRat(int64(numerator), int64(denominator)), nil
}
