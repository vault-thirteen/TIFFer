package tiff

import (
	"unsafe"

	"github.com/vault-thirteen/auxie/errors"

	"github.com/vault-thirteen/TIFFer/models"
	"github.com/vault-thirteen/TIFFer/models/Tag"
	"github.com/vault-thirteen/TIFFer/models/Type"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

const Errf_SizeErrorInType = "size error in type: %v"

// doSelfCheck performs various self checks of the library.
func doSelfCheck() (err error) {
	err = checkTypeSizes()
	if err != nil {
		return err
	}

	return nil
}

// checkTypeSizes checks size of various models.
// Some algorithms use dynamic size of variables, but some other parts of code
// use hard-coded methods, because otherwise the code would be slow. To ensure
// that everything work as intended, we must perform type size checks.
//
// Example. If for some reason someone decides to support bigger sizes and
// changes WORD type to DWORD, the hard-coded 'PutUint32' method will become
// useless for new type sizes.
func checkTypeSizes() (err error) {
	// These lists of constants are intentionally written inside the function !
	const (
		WordSize  = 2
		DWordSize = WordSize * 2
	)
	const (
		OffsetOfIFDSize              = DWordSize
		NumberOfDirectoryEntriesSize = WordSize
		TagSize                      = WordSize
		TypeSize                     = WordSize
		CountSize                    = DWordSize
		ValueOrOffsetSize            = DWordSize
		OffsetOfValueSize            = DWordSize
	)

	if uintptrToInt(unsafe.Sizeof(bt.Word(0))) != WordSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "WORD")
	}
	if uintptrToInt(unsafe.Sizeof(bt.DWord(0))) != DWordSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "DWORD")
	}

	if uintptrToInt(unsafe.Sizeof(models.OffsetOfIFD(0))) != OffsetOfIFDSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "OffsetOfIFD")
	}
	if uintptrToInt(unsafe.Sizeof(models.NumberOfDirectoryEntries(0))) != NumberOfDirectoryEntriesSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "NumberOfDirectoryEntries")
	}
	if uintptrToInt(unsafe.Sizeof(tag.Tag(0))) != TagSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "Tag")
	}
	if uintptrToInt(unsafe.Sizeof(t.Type(0))) != TypeSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "Type")
	}
	if uintptrToInt(unsafe.Sizeof(models.Count(0))) != CountSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "Count")
	}
	if uintptrToInt(unsafe.Sizeof(models.ValueOrOffset(0))) != ValueOrOffsetSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "ValueOrOffset")
	}
	if uintptrToInt(unsafe.Sizeof(models.OffsetOfValue(0))) != OffsetOfValueSize {
		return errors.ErrorFromFormattedText(Errf_SizeErrorInType, "OffsetOfValue")
	}

	return nil
}

// uintptrToInt converts uintptr into int.
func uintptrToInt(u uintptr) int {
	return *(*int)(unsafe.Pointer(&u))
}
