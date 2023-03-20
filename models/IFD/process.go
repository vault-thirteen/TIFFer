package ifd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/vault-thirteen/TIFFer/helper"
	"io"
	"math/big"
	"unsafe"

	"github.com/vault-thirteen/TIFFer/models"
	bo "github.com/vault-thirteen/TIFFer/models/ByteOrder"
	tag "github.com/vault-thirteen/TIFFer/models/Tag"
	"github.com/vault-thirteen/TIFFer/models/Type"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

func (de *DirectoryEntry) processDataItemSize() (err error) {
	switch de.Type {
	case t.Byte, // 8-bit.
		t.ASCII, // 7-bit.
		t.SByte: // 8-bit.
		de.dataItemSize = 1

	case t.Short, // 16-bit.
		t.SShort: // 16-bit.
		de.dataItemSize = 2

	case t.Long, // 32-bit.
		t.SLong: // 32-bit.
		de.dataItemSize = 4

	case t.Float: // 32-bit.
		de.dataItemSize = 4

	case t.Double, // 64-bit.
		t.Rational,  // 32-bit x2.
		t.SRational: // 32-bit x2.
		de.dataItemSize = 8

	case t.Undefined: // 8-bit.
		de.dataItemSize = 1

	default:
		return fmt.Errorf(t.ErrUnknownType, de.Type)
	}

	return nil
}

func (de *DirectoryEntry) processHasFastValue() {
	de.hasFastValue = uint(de.dataItemSize)*uint(de.Count) <= FastValueLimitSize
}

func (de *DirectoryEntry) processTagName() {
	var ok bool
	de.TagName, ok = tag.HumanReadableTagNames()[de.Tag]
	if ok {
		de.isTagKnown = true
	} else {
		de.isTagKnown = false
		de.TagName = tag.NameUnknown
	}
}

func (de *DirectoryEntry) processType() (err error) {
	if de.hasValidType() {
		return nil
	}

	return fmt.Errorf(ErrTypeIsNotValid, de.Type)
}

func (de *DirectoryEntry) processValue(
	rs models.ReaderSeeker,
	byteOrder bo.ByteOrder,
) (err error) {
	if de.hasFastValue {
		return de.readFastValue(byteOrder)
	}

	return de.readExternalValue(rs, byteOrder)
}

func (de *DirectoryEntry) readFastValue(byteOrder bo.ByteOrder) (err error) {
	de.Offset = 0

	var buf = make([]byte, unsafe.Sizeof(de.ValueOrOffset))
	switch byteOrder {
	case bo.BigEndian:
		binary.BigEndian.PutUint32(buf, de.ValueOrOffset)
	case bo.LittleEndian:
		binary.LittleEndian.PutUint32(buf, de.ValueOrOffset)
	default:
		return fmt.Errorf(bo.ErrUnsupportedBO, byteOrder)
	}

	de.Value, err = de.readValueFromStream(bytes.NewReader(buf), byteOrder)
	if err != nil {
		return err
	}

	return nil
}

func (de *DirectoryEntry) readExternalValue(
	rs models.ReaderSeeker,
	byteOrder bo.ByteOrder,
) (err error) {
	de.Offset = de.ValueOrOffset

	_, err = rs.Seek(int64(de.Offset), io.SeekStart)
	if err != nil {
		return err
	}

	de.Value, err = de.readValueFromStream(rs, byteOrder)
	if err != nil {
		return err
	}

	return nil
}

func (de *DirectoryEntry) readValueFromStream(
	r io.Reader,
	byteOrder bo.ByteOrder,
) (data any, err error) {
	switch de.Type {
	case t.Byte:
		return de.readArrayOfByte(r)
	case t.ASCII:
		return de.readArrayOfASCII(r)
	case t.Short:
		return de.readArrayOfShort(r, byteOrder)
	case t.Long:
		return de.readArrayOfLong(r, byteOrder)
	case t.Rational:
		return de.readArrayOfRational(r, byteOrder)
	case t.SByte:
		return de.readArrayOfSByte(r)
	case t.Undefined:
		return de.readArrayOfUndefined(r)
	case t.SShort:
		return de.readArrayOfSShort(r, byteOrder)
	case t.SLong:
		return de.readArrayOfSLong(r, byteOrder)
	case t.SRational:
		return de.readArrayOfSRational(r, byteOrder)
	case t.Float:
		return de.readArrayOfFloat(r, byteOrder)
	case t.Double:
		return de.readArrayOfDouble(r, byteOrder)
	default:
		return nil, fmt.Errorf(t.ErrUnknownType, de.Type)
	}
}

func (de *DirectoryEntry) readArrayOfByte(r io.Reader) (data []bt.Byte, err error) {
	data = make([]byte, 0, de.Count)
	var dataItem byte
	for i := models.Count(0); i < de.Count; i++ {
		dataItem, err = helper.ReadByte(r)
		if err != nil {
			return nil, err
		}
		data = append(data, dataItem)
	}
	return data, nil
}

func (de *DirectoryEntry) readArrayOfASCII(r io.Reader) (data []byte, err error) {
	data = make([]byte, 0, de.Count)
	var dataItem byte
	for i := models.Count(0); i < de.Count; i++ {
		dataItem, err = helper.ReadASCII(r)
		if err != nil {
			return nil, err
		}
		data = append(data, dataItem)
	}
	return data, nil
}

func (de *DirectoryEntry) readArrayOfShort(r io.Reader, byteOrder bo.ByteOrder) (data []bt.Word, err error) {
	data = make([]bt.Word, 0, de.Count)
	var dataItem bt.Word

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadShort_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadShort_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfLong(r io.Reader, byteOrder bo.ByteOrder) (data []bt.DWord, err error) {
	data = make([]bt.DWord, 0, de.Count)
	var dataItem bt.DWord

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadLong_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadLong_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfRational(r io.Reader, byteOrder bo.ByteOrder) (data []bt.Rational, err error) {
	data = make([]*big.Rat, 0, de.Count)
	var dataItem *big.Rat

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadRational_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadRational_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfSByte(r io.Reader) (data []bt.SByte, err error) {
	data = make([]int8, 0, de.Count)
	var dataItem int8
	for i := models.Count(0); i < de.Count; i++ {
		dataItem, err = helper.ReadSByte(r)
		if err != nil {
			return nil, err
		}
		data = append(data, dataItem)
	}
	return data, nil
}

func (de *DirectoryEntry) readArrayOfUndefined(r io.Reader) (data []bt.Byte, err error) {
	data = make([]byte, 0, de.Count)
	var dataItem byte
	for i := models.Count(0); i < de.Count; i++ {
		dataItem, err = helper.ReadUndefined(r)
		if err != nil {
			return nil, err
		}
		data = append(data, dataItem)
	}
	return data, nil
}

func (de *DirectoryEntry) readArrayOfSShort(r io.Reader, byteOrder bo.ByteOrder) (data []bt.SShort, err error) {
	data = make([]int16, 0, de.Count)
	var dataItem int16

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSShort_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSShort_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfSLong(r io.Reader, byteOrder bo.ByteOrder) (data []bt.SLong, err error) {
	data = make([]int32, 0, de.Count)
	var dataItem int32

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSLong_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSLong_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfSRational(r io.Reader, byteOrder bo.ByteOrder) (data []bt.SRational, err error) {
	data = make([]*big.Rat, 0, de.Count)
	var dataItem *big.Rat

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSRational_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadSRational_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfFloat(r io.Reader, byteOrder bo.ByteOrder) (data []bt.Float, err error) {
	data = make([]float32, 0, de.Count)
	var dataItem float32

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadFloat_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadFloat_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) readArrayOfDouble(r io.Reader, byteOrder bo.ByteOrder) (data []bt.Double, err error) {
	data = make([]float64, 0, de.Count)
	var dataItem float64

	switch byteOrder {
	case bo.BigEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadDouble_BE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	case bo.LittleEndian:
		for i := models.Count(0); i < de.Count; i++ {
			dataItem, err = helper.ReadDouble_LE(r)
			if err != nil {
				return nil, err
			}
			data = append(data, dataItem)
		}
	}

	return data, nil
}

func (de *DirectoryEntry) processHasSubIFD() {
	de.hasSubIFD = tag.IsSubIFDTag(de.Tag)
}
