package ifd

import (
	"fmt"
	"io"

	"github.com/vault-thirteen/TIFFer/models"
	"github.com/vault-thirteen/TIFFer/models/ByteOrder"
	"github.com/vault-thirteen/TIFFer/models/Tag"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
	"github.com/vault-thirteen/auxie/rs"
)

const (
	ErrSubIFDOffsetMustBeLong          = "sub-IFD offset must be long"
	ErrSubIFDCanNotHaveMultipleOffsets = "sub-IFD can not have multiple offsets"
)

// SubIFD is the Sub Image File Directory.
//
// It is not described in the TIFF 6.0 Specification, and it is someone's hack
// of the format. It is so ugly, that I do not have normal words to comment it.
// I hope that Adobe Inc., who supports all this mess, will bear the
// responsibility.
type SubIFD struct {
	// NumberOfDirectoryEntries is the number of directory entries.
	NumberOfDirectoryEntries models.NumberOfDirectoryEntries

	// DirectoryEntries is an array of directory entries.
	DirectoryEntries []*DirectoryEntry

	// OffsetOfNextSubIFD is an offset of the next SubIFD.
	OffsetOfNextSubIFD models.OffsetOfIFD

	// NextSubIFD is a pointer to the next SubIFD.
	NextSubIFD *SubIFD

	// Directory entries by tag number.
	DirectoryEntriesByTagNumber map[tag.Tag]*DirectoryEntry

	// Directory entries by tag name.
	// Except those who have empty or unknown tag names.
	DirectoryEntriesByTagName map[string]*DirectoryEntry

	// Statistics holds various statistical data about this SubIFD.
	Statistics *Statistics
}

// NewSubIFD constructs a first-pass model of a SubIFD from the stream.
// First-pass model means that we collect tags, data item models, data item
// counts, data item value offsets, but we do not read actual values.
func NewSubIFD(rs *rs.ReaderSeeker, byteOrder bo.ByteOrder, ifdOffset models.OffsetOfIFD) (si *SubIFD, err error) {
	_, err = rs.Seek(int64(ifdOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}

	switch byteOrder {
	case bo.BigEndian:
		return newSubIFD_BE(rs)
	case bo.LittleEndian:
		return newSubIFD_LE(rs)
	default:
		return nil, fmt.Errorf(bo.ErrUnsupportedBO, byteOrder)
	}
}

// newSubIFD_BE is a SubIFD first-pass constructor using big endian byte order.
func newSubIFD_BE(rs *rs.ReaderSeeker) (si *SubIFD, err error) {
	si = &SubIFD{
		Statistics: new(Statistics),
	}

	// Number of Directory Entries.
	si.NumberOfDirectoryEntries, err = rs.ReadWord_BE()
	if err != nil {
		return nil, err
	}

	// Directory Entries.
	si.DirectoryEntries = make([]*DirectoryEntry, 0)
	var e *DirectoryEntry
	for j := bt.Word(0); j < si.NumberOfDirectoryEntries; j++ {
		e, err = NewDE(rs, bo.BigEndian)
		if err != nil {
			return nil, err
		}

		si.DirectoryEntries = append(si.DirectoryEntries, e)
	}

	// OffsetOfValue of next SubIFD.
	si.OffsetOfNextSubIFD, err = rs.ReadDWord_BE()
	if err != nil {
		return nil, err
	}

	return si, nil
}

// newSubIFD_LE is a SubIFD first-pass constructor using little endian byte order.
func newSubIFD_LE(rs *rs.ReaderSeeker) (si *SubIFD, err error) {
	si = &SubIFD{
		Statistics: new(Statistics),
	}

	// Number of Directory Entries.
	si.NumberOfDirectoryEntries, err = rs.ReadWord_LE()
	if err != nil {
		return nil, err
	}

	// Directory Entries.
	si.DirectoryEntries = make([]*DirectoryEntry, 0)
	var e *DirectoryEntry
	for j := bt.Word(0); j < si.NumberOfDirectoryEntries; j++ {
		e, err = NewDE(rs, bo.LittleEndian)
		if err != nil {
			return nil, err
		}

		si.DirectoryEntries = append(si.DirectoryEntries, e)
	}

	// OffsetOfValue of next SubIFD.
	si.OffsetOfNextSubIFD, err = rs.ReadDWord_LE()
	if err != nil {
		return nil, err
	}

	return si, nil
}

// IsLast tells whether this SubIFD is last in the sequence or not.
func (si *SubIFD) IsLast() bool {
	return si.OffsetOfNextSubIFD == LastIFDOffsetOfNextIFD
}

// ProcessValues processes values of the SubIFD.
// Here we read values and try to decode (parse) them.
func (si *SubIFD) ProcessValues(rs *rs.ReaderSeeker, byteOrder bo.ByteOrder) (err error) {
	for _, curDE := range si.DirectoryEntries {
		err = curDE.ProcessValues(rs, byteOrder)
		if err != nil {
			return fmt.Errorf(ErrDE, curDE.Tag, curDE.TagName, err.Error())
		}
	}

	err = si.processDEMaps()
	if err != nil {
		return err
	}

	return nil
}

// processDEMaps fills the fast-access maps of Directory Entries.
func (si *SubIFD) processDEMaps() (err error) {
	si.DirectoryEntriesByTagNumber = make(map[tag.Tag]*DirectoryEntry)
	si.DirectoryEntriesByTagName = make(map[string]*DirectoryEntry)

	var isDuplicate bool
	for _, e := range si.DirectoryEntries {
		// Check for duplicates.
		_, isDuplicate = si.DirectoryEntriesByTagNumber[e.Tag]
		if isDuplicate {
			return fmt.Errorf(ErrDuplicateTagNumber, e.Tag)
		}

		_, isDuplicate = si.DirectoryEntriesByTagName[e.TagName]
		if isDuplicate {
			return fmt.Errorf(ErrDuplicateTagName, e.TagName)
		}

		// Save the ED into maps.
		si.DirectoryEntriesByTagNumber[e.Tag] = e

		if (len(e.TagName) > 0) && (e.TagName != tag.NameUnknown) {
			si.DirectoryEntriesByTagName[e.TagName] = e
		}
	}

	return nil
}

func (si *SubIFD) FillStatistics() {
	si.Statistics.KnownTagsCount = 0
	si.Statistics.UnKnownTagsCount = 0
	si.Statistics.CountOfTagsWithRegisteredType = 0
	si.Statistics.CountOfTagsWithUnRegisteredType = 0

	for _, e := range si.DirectoryEntries {
		if e.isTagKnown {
			si.Statistics.KnownTagsCount++
		} else {
			si.Statistics.UnKnownTagsCount++
		}

		if e.isTypeRegistered {
			si.Statistics.CountOfTagsWithRegisteredType++
		} else {
			si.Statistics.CountOfTagsWithUnRegisteredType++
		}
	}
}
