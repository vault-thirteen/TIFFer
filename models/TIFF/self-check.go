package tiff

import (
	"fmt"
	"unsafe"

	"github.com/vault-thirteen/TIFFer/models"
	"github.com/vault-thirteen/TIFFer/models/Tag"
	"github.com/vault-thirteen/TIFFer/models/Type"
	"github.com/vault-thirteen/TIFFer/models/basic-types"
)

const ErrSizeErrorInType = "size error in type: %v"

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
		return fmt.Errorf(ErrSizeErrorInType, "WORD")
	}
	if uintptrToInt(unsafe.Sizeof(bt.DWord(0))) != DWordSize {
		return fmt.Errorf(ErrSizeErrorInType, "DWORD")
	}

	if uintptrToInt(unsafe.Sizeof(models.OffsetOfIFD(0))) != OffsetOfIFDSize {
		return fmt.Errorf(ErrSizeErrorInType, "OffsetOfIFD")
	}
	if uintptrToInt(unsafe.Sizeof(models.NumberOfDirectoryEntries(0))) != NumberOfDirectoryEntriesSize {
		return fmt.Errorf(ErrSizeErrorInType, "NumberOfDirectoryEntries")
	}
	if uintptrToInt(unsafe.Sizeof(tag.Tag(0))) != TagSize {
		return fmt.Errorf(ErrSizeErrorInType, "Tag")
	}
	if uintptrToInt(unsafe.Sizeof(t.Type(0))) != TypeSize {
		return fmt.Errorf(ErrSizeErrorInType, "Type")
	}
	if uintptrToInt(unsafe.Sizeof(models.Count(0))) != CountSize {
		return fmt.Errorf(ErrSizeErrorInType, "Count")
	}
	if uintptrToInt(unsafe.Sizeof(models.ValueOrOffset(0))) != ValueOrOffsetSize {
		return fmt.Errorf(ErrSizeErrorInType, "ValueOrOffset")
	}
	if uintptrToInt(unsafe.Sizeof(models.OffsetOfValue(0))) != OffsetOfValueSize {
		return fmt.Errorf(ErrSizeErrorInType, "OffsetOfValue")
	}

	return nil
}

// uintptrToInt converts uintptr into int.
func uintptrToInt(u uintptr) int {
	return *(*int)(unsafe.Pointer(&u))
}
