package ifd

import (
	"github.com/vault-thirteen/TIFFermodels/Tag"
)

// requiredTagsForBilevelImages is a list of Tags required for Bilevel Images.
var requiredTagsForBilevelImages = []tag.Tag{
	tag.ImageWidth,
	tag.ImageLength,
	tag.Compression,
	tag.PhotometricInterpretation,
	tag.StripOffsets,
	tag.RowsPerStrip,
	tag.StripByteCounts,
	tag.XResolution,
	tag.YResolution,
	tag.ResolutionUnit,
}

// RequiredTagsForBilevelImages returns a list of tags required for bilevel
// images.
func RequiredTagsForBilevelImages() []tag.Tag {
	return requiredTagsForBilevelImages
}

// requiredTagsForGrayscaleImages is a list of Tags required for Grayscale
// Images.
var requiredTagsForGrayscaleImages = []tag.Tag{
	tag.ImageWidth,
	tag.ImageLength,
	tag.BitsPerSample, // *
	tag.Compression,
	tag.PhotometricInterpretation,
	tag.StripOffsets,
	tag.RowsPerStrip,
	tag.StripByteCounts,
	tag.XResolution,
	tag.YResolution,
	tag.ResolutionUnit,
}

// RequiredTagsForGrayscaleImages returns a list of tags required for grayscale
// images.
func RequiredTagsForGrayscaleImages() []tag.Tag {
	return requiredTagsForGrayscaleImages
}

// requiredTagsForRGBPaletteColorImages is a list of Tags required for RGB
// Palette Color Images.
var requiredTagsForRGBPaletteColorImages = []tag.Tag{
	tag.ImageWidth,
	tag.ImageLength,
	tag.BitsPerSample, // *
	tag.Compression,
	tag.PhotometricInterpretation,
	tag.StripOffsets,
	tag.RowsPerStrip,
	tag.StripByteCounts,
	tag.XResolution,
	tag.YResolution,
	tag.ResolutionUnit,
	tag.ColorMap, // **
}

// RequiredTagsForRGBPaletteColorImages returns a list of tags required for RGB
// palette colour images.
func RequiredTagsForRGBPaletteColorImages() []tag.Tag {
	return requiredTagsForRGBPaletteColorImages
}

// requiredTagsForRGBImages is a list of Tags required for RGB Images.
var requiredTagsForRGBImages = []tag.Tag{
	tag.ImageWidth,
	tag.ImageLength,
	tag.BitsPerSample, // *
	tag.Compression,
	tag.PhotometricInterpretation,
	tag.StripOffsets,
	tag.SamplesPerPixel, // ***
	tag.RowsPerStrip,
	tag.StripByteCounts,
	tag.XResolution,
	tag.YResolution,
	tag.ResolutionUnit,
}

// RequiredTagsForRGBImages returns a list of tags required for RGB images.
func RequiredTagsForRGBImages() []tag.Tag {
	return requiredTagsForRGBImages
}
