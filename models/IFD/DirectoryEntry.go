package ifd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/vault-thirteen/TIFFer/models"
	bo "github.com/vault-thirteen/TIFFer/models/ByteOrder"
	tag "github.com/vault-thirteen/TIFFer/models/Tag"
	t "github.com/vault-thirteen/TIFFer/models/Type"
	bt "github.com/vault-thirteen/TIFFer/models/basic-types"
	"github.com/vault-thirteen/auxie/NTS"
	"github.com/vault-thirteen/auxie/reader"
)

// FastValueLimitSize is the maximum amount of data which can be stored in the
// 'ValueOrOffset' field. Since this field is a DWORD, then it is 4 Bytes.
const FastValueLimitSize = 4

const (
	ErrTypeCastFailure = "type casting has failed"
	ErrInFirstSubIFD   = "error in first SubIFD: %v"
	ErrInNthSubIFD     = "error in SubIFD #%v: %v"
)

// DirectoryEntry is the Directory Entry described in the TIFF 6.0
// Specification.
type DirectoryEntry struct {
	// Tag is the tag of a Directory Entry.
	Tag tag.Tag

	// Type is the type of data items of a Directory Entry.
	Type t.Type

	// Count is the number of data items of a Directory Entry.
	Count models.Count

	// ValueOrOffset stores either a value or an offset of a value.
	// The computed (real) values of value and offset are stored in Value and
	// Offset fields.
	ValueOrOffset models.ValueOrOffset

	// Offset is the computed size of the offset of the value.
	// When 'ValueOrOffset' field contains a value, Offset is zero. When
	// 'ValueOrOffset' field contains a value offset, Offset must be non-zero.
	Offset models.OffsetOfValue

	// Value is the computed value of the value.
	// Normally, value is an array of data items.
	Value any

	// Below are the fields for internal usage.

	// Data item size (in Bytes).
	dataItemSize byte

	// hasFastValue flag indicates that data items value fits in the
	// ValueOrOffset field. If the value is fast, it is stored in the
	// ValueOrOffset field. If the value is not fast, then its offset is stored
	// in the ValueOrOffset field.
	hasFastValue bool

	// TagName is a human-readable tag name.
	TagName string

	// hasSubIFD shows whether this DE has a sub-IFD or not.
	hasSubIFD bool

	// SubIFD is the sub-IFD for the tag.
	SubIFD *SubIFD

	// Chain of SubIFDs.
	SubIFDs []*SubIFD

	// isTagKnown flag is true for known tags, i.e. for those tags which have a
	// known name, i.e. textual alias.
	isTagKnown bool

	// isTypeRegistered flag is true for those tags which have a type rule.
	// Those tags who have no type rule are automatically considered as valid
	// as stated in the TIFF 6.0 Specification. To count those "shadow" tags,
	// we use this flag.
	isTypeRegistered bool
}

// NewDE constructs a first-pass model of a Directory Entry from the stream.
// First-pass model means that we collect tags, data item models, data item
// counts, data item value offsets, but we do not read actual values.
func NewDE(rs *reader.Reader, byteOrder bo.ByteOrder) (de *DirectoryEntry, err error) {
	switch byteOrder {
	case bo.BigEndian:
		return newDE_BE(rs)
	case bo.LittleEndian:
		return newDE_LE(rs)
	default:
		return nil, fmt.Errorf(bo.ErrUnsupportedBO, byteOrder)
	}
}

// newDE_BE is a Directory Entry first-pass constructor using big endian byte
// order.
func newDE_BE(rs *reader.Reader) (e *DirectoryEntry, err error) {
	e = &DirectoryEntry{}

	// Tag.
	e.Tag, err = rs.ReadWord_BE()
	if err != nil {
		return nil, err
	}

	// Type.
	e.Type, err = rs.ReadWord_BE()
	if err != nil {
		return nil, err
	}

	// Count.
	e.Count, err = rs.ReadDWord_BE()
	if err != nil {
		return nil, err
	}

	// Value or OffsetOfValue.
	e.ValueOrOffset, err = rs.ReadDWord_BE()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// newDE_LE is a Directory Entry first-pass constructor using little endian byte
// order.
func newDE_LE(rs *reader.Reader) (e *DirectoryEntry, err error) {
	e = &DirectoryEntry{}

	// Tag.
	e.Tag, err = rs.ReadWord_LE()
	if err != nil {
		return nil, err
	}

	// Type.
	e.Type, err = rs.ReadWord_LE()
	if err != nil {
		return nil, err
	}

	// Count.
	e.Count, err = rs.ReadDWord_LE()
	if err != nil {
		return nil, err
	}

	// Value or OffsetOfValue.
	e.ValueOrOffset, err = rs.ReadDWord_LE()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// ProcessValues processes the directory entry data.
// Here we read values and try to decode (parse) them.
func (de *DirectoryEntry) ProcessValues(rs *reader.Reader, byteOrder bo.ByteOrder) (err error) {
	err = de.processDataItemSize()
	if err != nil {
		return err
	}

	de.processHasFastValue()
	de.processTagName()

	err = de.processType()
	if err != nil {
		return err
	}

	err = de.processValue(rs, byteOrder)
	if err != nil {
		return err
	}

	return nil
}

// ValueAsArrayOfByte tries to return the value as array of bytes.
func (de *DirectoryEntry) ValueAsArrayOfByte() (v []bt.Byte, err error) {
	var ok bool
	v, ok = de.Value.([]byte)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfString tries to return the value as array of strings.
func (de *DirectoryEntry) ValueAsArrayOfString() (v []string, err error) {
	var ok bool
	v, ok = de.Value.([]string)
	if ok {
		return v, nil
	}

	var buf []byte
	buf, ok = de.Value.([]byte)
	if ok {
		return nts.ByteArrayToStrings(buf)
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfShort tries to return the value as array of shorts.
func (de *DirectoryEntry) ValueAsArrayOfShort() (v []bt.Word, err error) {
	var ok bool
	v, ok = de.Value.([]bt.Word)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfLong tries to return the value as array of longs.
func (de *DirectoryEntry) ValueAsArrayOfLong() (v []bt.DWord, err error) {
	var ok bool
	v, ok = de.Value.([]bt.DWord)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfRational tries to return the value as array of rationals.
func (de *DirectoryEntry) ValueAsArrayOfRational() (v []bt.Rational, err error) {
	var ok bool
	v, ok = de.Value.([]*big.Rat)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfSByte tries to return the value as array of signed bytes.
func (de *DirectoryEntry) ValueAsArrayOfSByte() (v []bt.SByte, err error) {
	var ok bool
	v, ok = de.Value.([]int8)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfUndefined tries to return the value as array of unknowns.
// TIFF 6.0 Specification states that Unknown type is the Byte type.
func (de *DirectoryEntry) ValueAsArrayOfUndefined() (v []bt.Byte, err error) {
	var ok bool
	v, ok = de.Value.([]byte)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfSShort tries to return the value as array of signed shorts.
func (de *DirectoryEntry) ValueAsArrayOfSShort() (v []bt.SShort, err error) {
	var ok bool
	v, ok = de.Value.([]int16)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfSLong tries to return the value as array of signed longs.
func (de *DirectoryEntry) ValueAsArrayOfSLong() (v []bt.SLong, err error) {
	var ok bool
	v, ok = de.Value.([]int32)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfSRational tries to return the value as array of signed
// rationals.
func (de *DirectoryEntry) ValueAsArrayOfSRational() (v []bt.SRational, err error) {
	var ok bool
	v, ok = de.Value.([]*big.Rat)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfFloat tries to return the value as array of floats.
func (de *DirectoryEntry) ValueAsArrayOfFloat() (v []bt.Float, err error) {
	var ok bool
	v, ok = de.Value.([]float32)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// ValueAsArrayOfDouble tries to return the value as array of doubles.
func (de *DirectoryEntry) ValueAsArrayOfDouble() (v []bt.Double, err error) {
	var ok bool
	v, ok = de.Value.([]float64)
	if ok {
		return v, nil
	}

	return nil, errors.New(ErrTypeCastFailure)
}

// DataItemSize returns the size (in Bytes) of a data item of a Directory Entry.
func (de *DirectoryEntry) DataItemSize() int {
	return int(de.dataItemSize)
}

// HasFastValue tells whether the value of the Directory Entry is fast or not.
func (de *DirectoryEntry) HasFastValue() bool {
	return de.hasFastValue
}

// HasSubIFD shows whether this DE has a sub-IFD or not.
func (de *DirectoryEntry) HasSubIFD() bool {
	return de.hasSubIFD
}

// ProcessSubIFDs processes the directory entry sub-IFDs.
// Here we read sub-IFDs of all tags who have them.
func (de *DirectoryEntry) ProcessSubIFDs(rs *reader.Reader, byteOrder bo.ByteOrder) (err error) {
	de.processHasSubIFD()

	if !de.hasSubIFD {
		return nil
	}

	// OK. This tag has a SubIFD.
	de.SubIFDs = make([]*SubIFD, 0)

	// The 'Sub-IFD' is not described in the TIFF 6.0 Specification and
	// documentation for it is very poor, so we better make some fool checks.
	if de.Type != t.Long {
		return errors.New(ErrSubIFDOffsetMustBeLong)
	}
	if de.Count != 1 {
		return errors.New(ErrSubIFDCanNotHaveMultipleOffsets)
	}

	// Pass I.
	err = de.readSubIFDPassOne(rs, byteOrder)
	if err != nil {
		return err
	}

	// Pass II.
	err = de.readSubIFDPassTwo(rs, byteOrder)
	if err != nil {
		return err
	}

	return nil
}

// readSubIFDPassOne performs a first read pass of the SubIFD.
// In this pass we briefly read structures.
func (de *DirectoryEntry) readSubIFDPassOne(rs *reader.Reader, byteOrder bo.ByteOrder) (err error) {
	var si *SubIFD

	// First SubIFD.
	si, err = NewSubIFD(rs, byteOrder, de.ValueOrOffset)
	if err != nil {
		return fmt.Errorf(ErrInFirstSubIFD, err.Error())
	}
	de.SubIFDs = append(de.SubIFDs, si)
	de.SubIFD = de.SubIFDs[0]

	// Last read SubIFD.
	var lrSubIFD = de.lastReadSubIFD()

	// We know that normal IFDs are chained. Does it work for SubIFDs ?
	// We can not say for sure, but we can try to read them !
	// Try to read the rest chain of SubIFDs for any case.
	// We do not know what else those TIFF-format-hackers prepared for us.
	n := 2
	if !lrSubIFD.IsLast() {
		si, err = NewSubIFD(rs, byteOrder, lrSubIFD.OffsetOfNextSubIFD)
		if err != nil {
			return fmt.Errorf(ErrInNthSubIFD, n, err.Error())
		}
		de.SubIFDs = append(de.SubIFDs, si)
		lrSubIFD = de.lastReadSubIFD()
		n++
	}

	// SubIFD links.
	var idxMax = len(de.SubIFDs) - 1
	for idx, curSubIFD := range de.SubIFDs {
		if !curSubIFD.IsLast() {
			if idx+1 > idxMax {
				return fmt.Errorf(ErrUnexpectedSequenceEnd, idx)
			}
			curSubIFD.NextSubIFD = de.SubIFDs[idx+1]
		}
	}

	return nil
}

// readSubIFDPassTwo performs a second-pass read of the SubIFD.
// In this pass we read values and try to decode them.
func (de *DirectoryEntry) readSubIFDPassTwo(rs *reader.Reader, byteOrder bo.ByteOrder) (err error) {
	for _, curIFD := range de.SubIFDs {
		err = curIFD.ProcessValues(rs, byteOrder)
		if err != nil {
			return err
		}

		curIFD.FillStatistics()
	}

	return nil
}

// lastReadSubIFD returns the last read SubIFD.
func (de *DirectoryEntry) lastReadSubIFD() *SubIFD {
	l := len(de.SubIFDs)

	if l == 0 {
		return nil
	}

	return de.SubIFDs[l-1]
}
