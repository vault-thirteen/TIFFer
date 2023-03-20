package ifd

import (
	"fmt"
	"io"

	"github.com/vault-thirteen/TIFFer/helper"
	"github.com/vault-thirteen/TIFFer/models"
	"github.com/vault-thirteen/TIFFer/models/ByteOrder"
	"github.com/vault-thirteen/TIFFer/models/Tag"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

const (
	ErrUnexpectedSequenceEnd = "unexpected sequence end in IFD[%v]"
	ErrDuplicateTagNumber    = "duplicate tag number: %v"
	ErrDuplicateTagName      = "duplicate tag name: %v"
	ErrDE                    = `error in DE (Tag=%v,TagName="%v"): %v`
)

// LastIFDOffsetOfNextIFD is the value of 'OffsetOfIFD' field of the last
// IFD in the sequence.
const LastIFDOffsetOfNextIFD = 0

// IFD is the Image File Directory described in the TIFF 6.0 Specification.
type IFD struct {
	// NumberOfDirectoryEntries is the number of directory entries.
	NumberOfDirectoryEntries models.NumberOfDirectoryEntries

	// DirectoryEntries is an array of directory entries.
	DirectoryEntries []*DirectoryEntry

	// OffsetOfNextIFD is an offset of the next IFD.
	OffsetOfNextIFD models.OffsetOfIFD

	// NextIFD is a pointer to the next IFD.
	NextIFD *IFD

	// Directory entries by tag number.
	DirectoryEntriesByTagNumber map[tag.Tag]*DirectoryEntry

	// Directory entries by tag name.
	// Except those who have empty or unknown tag names.
	DirectoryEntriesByTagName map[string]*DirectoryEntry

	// Statistics holds various statistical data about this IFD.
	Statistics *Statistics
}

// NewIFD constructs a first-pass model of an IFD from the stream.
// First-pass model means that we collect tags, data item models, data item
// counts, data item value offsets, but we do not read actual values.
func NewIFD(
	rs models.ReaderSeeker,
	byteOrder bo.ByteOrder,
	ifdOffset models.OffsetOfIFD,
) (i *IFD, err error) {
	_, err = rs.Seek(int64(ifdOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}

	switch byteOrder {
	case bo.BigEndian:
		return newIFD_BE(rs)
	case bo.LittleEndian:
		return newIFD_LE(rs)
	default:
		return nil, fmt.Errorf(bo.ErrUnsupportedBO, byteOrder)
	}
}

// newIFD_BE is an IFD first-pass constructor using big endian byte order.
func newIFD_BE(r io.Reader) (i *IFD, err error) {
	i = &IFD{
		Statistics: new(Statistics),
	}

	// Number of Directory Entries.
	i.NumberOfDirectoryEntries, err = helper.ReadWord_BE(r)
	if err != nil {
		return nil, err
	}

	// Directory Entries.
	i.DirectoryEntries = make([]*DirectoryEntry, 0)
	var e *DirectoryEntry
	for j := bt.Word(0); j < i.NumberOfDirectoryEntries; j++ {
		e, err = NewDE(r, bo.BigEndian)
		if err != nil {
			return nil, err
		}

		i.DirectoryEntries = append(i.DirectoryEntries, e)
	}

	// OffsetOfValue of next IFD.
	i.OffsetOfNextIFD, err = helper.ReadDWord_BE(r)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// newIFD_LE is an IFD first-pass constructor using little endian byte order.
func newIFD_LE(r io.Reader) (i *IFD, err error) {
	i = &IFD{
		Statistics: new(Statistics),
	}

	// Number of Directory Entries.
	i.NumberOfDirectoryEntries, err = helper.ReadWord_LE(r)
	if err != nil {
		return nil, err
	}

	// Directory Entries.
	i.DirectoryEntries = make([]*DirectoryEntry, 0)
	var e *DirectoryEntry
	for j := bt.Word(0); j < i.NumberOfDirectoryEntries; j++ {
		e, err = NewDE(r, bo.LittleEndian)
		if err != nil {
			return nil, err
		}

		i.DirectoryEntries = append(i.DirectoryEntries, e)
	}

	// OffsetOfValue of next IFD.
	i.OffsetOfNextIFD, err = helper.ReadDWord_LE(r)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// IsLast tells whether this IFD is last in the sequence or not.
func (i *IFD) IsLast() bool {
	return i.OffsetOfNextIFD == LastIFDOffsetOfNextIFD
}

// ProcessValues processes values of the IFD.
// Here we read values and try to decode (parse) them.
func (i *IFD) ProcessValues(rs models.ReaderSeeker, byteOrder bo.ByteOrder) (err error) {
	for _, curDE := range i.DirectoryEntries {
		err = curDE.ProcessValues(rs, byteOrder)
		if err != nil {
			return fmt.Errorf(ErrDE, curDE.Tag, curDE.TagName, err.Error())
		}
	}

	err = i.processDEMaps()
	if err != nil {
		return err
	}

	return nil
}

// processDEMaps fills the fast-access maps of Directory Entries.
func (i *IFD) processDEMaps() (err error) {
	i.DirectoryEntriesByTagNumber = make(map[tag.Tag]*DirectoryEntry)
	i.DirectoryEntriesByTagName = make(map[string]*DirectoryEntry)

	var isDuplicate bool
	for _, e := range i.DirectoryEntries {
		// Check for duplicates.
		_, isDuplicate = i.DirectoryEntriesByTagNumber[e.Tag]
		if isDuplicate {
			return fmt.Errorf(ErrDuplicateTagNumber, e.Tag)
		}

		_, isDuplicate = i.DirectoryEntriesByTagName[e.TagName]
		if isDuplicate {
			return fmt.Errorf(ErrDuplicateTagName, e.TagName)
		}

		// Save the ED into maps.
		i.DirectoryEntriesByTagNumber[e.Tag] = e

		if (len(e.TagName) > 0) && (e.TagName != tag.NameUnknown) {
			i.DirectoryEntriesByTagName[e.TagName] = e
		}
	}

	return nil
}

// ProcessSubIFDs processes sub-IFDs of the IFD.
// Here we read sub-IFDs of all tags who have them.
func (i *IFD) ProcessSubIFDs(rs models.ReaderSeeker, byteOrder bo.ByteOrder) (err error) {
	for _, curDE := range i.DirectoryEntries {
		err = curDE.ProcessSubIFDs(rs, byteOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IFD) FillStatistics() {
	i.Statistics.KnownTagsCount = 0
	i.Statistics.UnKnownTagsCount = 0
	i.Statistics.CountOfTagsWithRegisteredType = 0
	i.Statistics.CountOfTagsWithUnRegisteredType = 0

	for _, e := range i.DirectoryEntries {
		if e.isTagKnown {
			i.Statistics.KnownTagsCount++
		} else {
			i.Statistics.UnKnownTagsCount++
		}

		if e.isTypeRegistered {
			i.Statistics.CountOfTagsWithRegisteredType++
		} else {
			i.Statistics.CountOfTagsWithUnRegisteredType++
		}
	}
}
