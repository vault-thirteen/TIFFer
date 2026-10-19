package main

import (
	"math/big"
	"os"

	"github.com/vault-thirteen/auxie/cio"
	ae "github.com/vault-thirteen/auxie/errors"

	tiff "github.com/vault-thirteen/TIFFer/models/TIFF"
	tag "github.com/vault-thirteen/TIFFer/models/Tag"
	bt "github.com/vault-thirteen/TIFFer/models/basic-types"
)

const UsageHint = `Usage:
	example.exe file.tiff`

const (
	ExitCode_ArgumentsError = 1
)

func main() {
	if len(os.Args) < 2 {
		cio.PrintMessage(UsageHint)
		os.Exit(ExitCode_ArgumentsError)
	}
	filePath := os.Args[1]

	err := work(filePath)
	if err != nil {
		panic(err)
	}
}

func work(filePath string) (err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return err
	}

	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	var t *tiff.TIFF
	t, err = tiff.New(f)
	if err != nil {
		return err
	}

	err = playInSandbox(t)
	if err != nil {
		return err
	}

	return nil
}

func playInSandbox(t *tiff.TIFF) (err error) {
	var bb []byte
	bb, err = t.IFDs()[0].DirectoryEntriesByTagName["ImageDescription"].ValueAsArrayOfByte()
	if err != nil {
		return err
	}
	cio.PrintMessage("ImageDescription:", bb)

	var ss []string
	ss, err = t.IFDs()[0].DirectoryEntriesByTagName["ImageDescription"].ValueAsArrayOfString()
	if err != nil {
		return err
	}
	cio.PrintMessage("ImageDescription:", ss)

	var ww []bt.Word
	ww, err = t.IFDs()[0].DirectoryEntriesByTagName["BitsPerSample"].ValueAsArrayOfShort()
	if err != nil {
		return err
	}
	cio.PrintMessage("BitsPerSample:", ww)

	var dwdw []bt.DWord
	dwdw, err = t.IFDs()[0].DirectoryEntriesByTagName["StripOffsets"].ValueAsArrayOfLong()
	if err != nil {
		return err
	}
	cio.PrintMessage("StripOffsets:", dwdw)

	var rr []*big.Rat
	rr, err = t.IFDs()[0].DirectoryEntriesByTagName["XResolution"].ValueAsArrayOfRational()
	if err != nil {
		return err
	}
	cio.PrintMessage("XResolution:", rr)

	var uu []byte
	uu, err = t.IFDs()[0].DirectoryEntriesByTagName["ICCProfile"].ValueAsArrayOfUndefined()
	if err != nil {
		return err
	}
	cio.PrintMessage("ICCProfile:", uu)

	// Play with some EXIF data.
	var fNumber []*big.Rat
	fNumber, err = t.IFDs()[0].DirectoryEntriesByTagName["ExifIFD"].SubIFDs[0].
		DirectoryEntriesByTagNumber[tag.FNumber].ValueAsArrayOfRational()
	if err != nil {
		return err
	}
	cio.PrintMessage("F-Number (EXIF Tag):", fNumber)

	// Play with some GPS data.
	var gpsAltitude []*big.Rat
	gpsAltitude, err = t.IFDs()[0].DirectoryEntriesByTagName["GPSIFD"].SubIFD.
		DirectoryEntriesByTagName["GPSAltitude"].ValueAsArrayOfRational()
	if err != nil {
		return err
	}
	cio.PrintMessage("GPS Altitude (GPS Tag):", gpsAltitude)

	return nil
}
