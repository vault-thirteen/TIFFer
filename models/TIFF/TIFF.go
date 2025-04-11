package tiff

import (
	"fmt"

	hdr "github.com/vault-thirteen/TIFFer/models/Header"
	ifd "github.com/vault-thirteen/TIFFer/models/IFD"
	iors "github.com/vault-thirteen/auxie/ReaderSeeker"
	"github.com/vault-thirteen/auxie/rs"
)

const (
	ErrInFirstIFD = "error in first IFD: %v"
	ErrInNthIFD   = "error in IFD #%v: %v"
)

// TIFF is an object storing information about TIFF file conforming to the TIFF
// 6.0 Specification.
//
// TIFF format was developed by	Aldus Corporation. At this moment, in the year
// 2023, the owner of this technology is Adobe Inc., who bought Aldus
// Corporation in September 1994. More information about TIFF format can be
// found in Wikipedia: https://en.wikipedia.org/wiki/TIFF
//
// TIFF specification (Revision 6.0 Final — June 3, 1992):
// https://developer.adobe.com/content/dam/udp/en/open/standards/tiff/TIFF6.pdf
//
// Current implementation of this format is experimental.
// The main purpose of this library is reading meta-data from TIFF files.
type TIFF struct {
	// header is the TIFF header.
	header *hdr.Header

	// ifds is a list of IFDs.
	ifds []*ifd.IFD
}

// New constructs the TIFF object from the byte reader.
//
// Due to the nature of TIFF format, which is highly complex and controversial,
// the object is constructed in several passes, similar to some video encoders
// that require two passes. On the first pass we collect brief data about IFDs
// and their directory entries. On the second pass we collect data item values
// for those tags which we are interested in. On the third pass we collect
// information about so-called Sub-IFDs, which are not a part of the
// TIFF 6.0 Specification, but they are used by some tools.
func New(stream iors.ReaderSeeker) (t *TIFF, err error) {
	t = &TIFF{
		ifds: make([]*ifd.IFD, 0),
	}

	err = doSelfCheck()
	if err != nil {
		return nil, err
	}

	var readerSeeker *rs.ReaderSeeker
	readerSeeker, err = rs.New(stream)
	if err != nil {
		return nil, err
	}

	// Header.
	t.header, err = hdr.New(readerSeeker)
	if err != nil {
		return nil, err
	}

	// Pass I.
	err = t.readPassOne(readerSeeker)
	if err != nil {
		return nil, err
	}

	// Pass II.
	err = t.readPassTwo(readerSeeker)
	if err != nil {
		return nil, err
	}

	// Pass III.
	err = t.readPassThree(readerSeeker)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// readPassOne performs a first read pass of the TIFF object.
// In this pass we briefly read structures.
func (t *TIFF) readPassOne(rs *rs.ReaderSeeker) (err error) {
	var i *ifd.IFD

	// First IFD.
	i, err = ifd.NewIFD(rs, t.header.ByteOrder, t.header.OffsetOfFirstIFD)
	if err != nil {
		return fmt.Errorf(ErrInFirstIFD, err.Error())
	}
	t.ifds = append(t.ifds, i)
	t.header.FirstIFD = t.ifds[0]

	// Last read IFD.
	var lrIFD = t.lastReadIFD()

	// Rest IFDs.
	n := 2
	for !lrIFD.IsLast() {
		i, err = ifd.NewIFD(rs, t.header.ByteOrder, lrIFD.OffsetOfNextIFD)
		if err != nil {
			return fmt.Errorf(ErrInNthIFD, n, err.Error())
		}
		t.ifds = append(t.ifds, i)
		lrIFD = t.lastReadIFD()
		n++
	}

	// IFD links.
	var idxMax = len(t.ifds) - 1
	for idx, curIFD := range t.ifds {
		if !curIFD.IsLast() {
			if idx+1 > idxMax {
				return fmt.Errorf(ifd.ErrUnexpectedSequenceEnd, idx)
			}
			curIFD.NextIFD = t.ifds[idx+1]
		}
	}

	return nil
}

// readPassTwo performs a second-pass read of the TIFF object.
// In this pass we read values and try to decode them.
func (t *TIFF) readPassTwo(rs *rs.ReaderSeeker) (err error) {
	for _, curIFD := range t.ifds {
		err = curIFD.ProcessValues(rs, t.header.ByteOrder)
		if err != nil {
			return err
		}

		curIFD.FillStatistics()
	}

	return nil
}

// readPassThree performs a third-pass read of the TIFF object.
// In this pass we read sub-IFDs of tags.
func (t *TIFF) readPassThree(rs *rs.ReaderSeeker) (err error) {
	for _, curIFD := range t.ifds {
		err = curIFD.ProcessSubIFDs(rs, t.header.ByteOrder)
		if err != nil {
			return err
		}
	}

	return nil
}

// lastReadIFD returns the last read IFD.
func (t *TIFF) lastReadIFD() *ifd.IFD {
	l := len(t.ifds)

	if l == 0 {
		return nil
	}

	return t.ifds[l-1]
}

// Header returns TIFF's header.
func (t *TIFF) Header() (header *hdr.Header) {
	return t.header
}

// IFDs returns an array of TIFF's IFDs.
func (t *TIFF) IFDs() (ifds []*ifd.IFD) {
	return t.ifds
}
