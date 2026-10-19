package models

// SubfileType.
const (
	SubfileTypeFullResolutionImageData    = 1
	SubfileTypeReducedResolutionImageData = 2
	SubfileTypeSinglePageOfMultiPageImage = 3
)

// BitsPerSample.
const (
	BitsPerSample4 = 4 // 16 shades of grey.
	BitsPerSample8 = 8 // 256 shades of grey.
)

// Compression.
const (
	// CompressionNone
	/* No compression, but pack data into bytes as tightly as possible, leaving
	no unused bits (except at the end of a row). The component values are
	stored as an array of type BYTE. Each scan line (row) is padded to the next
	BYTE boundary. */
	CompressionNone = 1

	// CompressionCCITTGroup3
	/* CCITT Group 3 1-Dimensional Modified Huffman run length encoding. See
	Section 10 for a description of Modified Huffman Compression. */
	CompressionCCITTGroup3 = 2

	// CompressionT4
	/* T4-encoding: CCITT T.4 bi-level encoding as specified in section 4,
	Coding, of CCITT Recommendation T.4: “Standardization of Group 3 Facsimile
	apparatus for document transmission.” International Telephone and Telegraph
	Consultative Committee (CCITT, Geneva: 1988). */
	CompressionT4 = 3

	// CompressionT6
	/* T6-encoding: CCITT T.6 bi-level encoding as specified in section 2 of
	CCITT Recommendation T.6: “Facsimile coding schemes and coding control
	functions for Group 4 facsimile apparatus.” International Telephone and
	Telegraph Consultative Committee (CCITT, Geneva: 1988). */
	CompressionT6 = 4

	CompressionJPEG = 6

	// CompressionPackBits
	/* PackBits compression, a simple byte-oriented run length scheme. See the
	PackBits section for details. Data compression applies only to raster image
	data. All other TIFF fields are unaffected. */
	CompressionPackBits = 32773
)

// PhotometricInterpretation.
const (
	// PhotometricInterpretationWhiteIsZero
	/* For bilevel and grayscale images: 0 is imaged as white. The maximum
	value is imaged as black. This is the normal value for Compression=2. */
	PhotometricInterpretationWhiteIsZero = 0

	// PhotometricInterpretationBlackIsZero
	/* For bilevel and grayscale images: 0 is imaged as black. The maximum
	value is imaged as white. If this value is specified for Compression=2,
	 the image should display and print reversed. */
	PhotometricInterpretationBlackIsZero = 1

	// PhotometricInterpretationRGB
	/* RGB is the value of PhotometricInterpretation which indicates RGB Full
	Color Images. */
	PhotometricInterpretationRGB = 2

	// PhotometricInterpretationRGBPaletteColor
	/* In this model, a color is described with a single component. The value
	of the component is used as an index into the red, green and blue curves in
	the ColorMap field to retrieve an RGB triplet that defines the color. When
	PhotometricInterpretation=3 is used, ColorMap must be present and
	SamplesPerPixel must be 1. */
	PhotometricInterpretationRGBPaletteColor = 3

	// PhotometricInterpretationTransparencyMask
	/* This means that the image is used to define an irregularly shaped region
	of another image in the same TIFF file. SamplesPerPixel and BitsPerSample
	must be 1. PackBits compression is recommended. The 1-bits define the
	interior of the region; the 0-bits define the exterior of the region. */
	PhotometricInterpretationTransparencyMask = 4

	PhotometricInterpretationCMYK = 5

	PhotometricInterpretatioYCbCr = 6

	// PhotometricInterpretatioCIELAB
	// 1976 CIE L*a*b*.
	PhotometricInterpretatioCIELAB = 8
)

// Threshholding.
const (
	// ThreshholdingNoDitheringOrHalftoning
	/* No dithering or halftoning has been applied to the image data. */
	ThreshholdingNoDitheringOrHalftoning = 1

	// ThreshholdingOrderedDitherOrHalftone
	/* An ordered dither or halftone technique has been applied to the image
	data. */
	ThreshholdingOrderedDitherOrHalftone = 2

	// ThreshholdingRandomized
	/* A randomized process such as error diffusion has been applied to the
	image data. */
	ThreshholdingRandomized = 3
)

// FillOrder.
const (
	// FillOrder1
	/* Pixels are arranged within a byte such that pixels with lower column
	values are stored in the higher-order bits of the byte. */
	FillOrder1 = 1

	// FillOrder2
	/* Pixels are arranged within a byte such that pixels with lower column
	values are stored in the lower-order bits of the byte. */
	FillOrder2 = 2
)

// Orientation.
const (
	// Orientation1
	/* The 0th row represents the visual top of the image, and the 0th column
	represents the visual left-hand side. */
	Orientation1 = 1

	// Orientation2
	/* The 0th row represents the visual top of the image, and the 0th column
	represents the visual right-hand side. */
	Orientation2 = 2

	// Orientation3
	/* The 0th row represents the visual bottom of the image, and the 0th
	column represents the visual right-hand side. */
	Orientation3 = 3

	// Orientation4
	/* The 0th row represents the visual bottom of the image, and the 0th
	column represents the visual left-hand side. */
	Orientation4 = 4

	// Orientation5
	/* The 0th row represents the visual left-hand side of the image, and the
	0th column represents the visual top. */
	Orientation5 = 5

	// Orientation6
	/* The 0th row represents the visual right-hand side of the image, and the
	0th column represents the visual top. */
	Orientation6 = 6

	// Orientation7
	/* The 0th row represents the visual right-hand side of the image, and the
	0th column represents the visual bottom. */
	Orientation7 = 7

	// Orientation8
	/* The 0th row represents the visual left-hand side of the image, and the
	0th column represents the visual bottom. */
	Orientation8 = 8
)

// SamplesPerPixel.
const (
	// TagSamplesPerPixel3
	/* The number of components per pixel. This number is 3 for RGB images,
	unless extra samples are present. See the ExtraSamples field for further
	information. */
	TagSamplesPerPixel3 = 3

	TagSamplesPerPixel4 = 4
)

// PlanarConfiguration.
const (
	// PlanarConfigurationChunky
	/* Chunky format. The component values for each pixel are stored
	contiguously. The order of the components within the pixel is specified by
	PhotometricInterpretation. For example, for RGB data, the data is stored as
	RGBRGBRGB... */
	PlanarConfigurationChunky = 1

	// PlanarConfigurationPlanar
	/* The components are stored in separate “component planes.” The values in
	StripOffsets and StripByteCounts are then arranged as a 2-dimensional
	array, with SamplesPerPixel rows and StripsPerImage columns. (All of the
	columns for row 0 are stored first, followed by the columns of row 1, and
	so on.) PhotometricInterpretation describes the type of data stored in each
	component plane. For example, RGB data is stored with the Red components in
	one component plane, the Green in another, and the Blue in another. */
	PlanarConfigurationPlanar = 2
)

// ResolutionUnit.
const (
	// ResolutionUnitNone
	/* No absolute unit of measurement. Used for images that may have a
	non-square aspect ratio but no meaningful absolute dimensions. */
	ResolutionUnitNone = 1

	// ResolutionUnitInch
	/* Inch. */
	ResolutionUnitInch = 2

	// ResolutionUnitCentimeter
	/* Centimeter. */
	ResolutionUnitCentimeter = 3
)

// Predictor.
const (
	PredictorNone                   = 1
	PredictorHorizontalDifferencing = 2
)

// InkSet.
const (
	// InkSetCMYK
	/* The order of the components is cyan, magenta, yellow, black. Usually, a
	value of 0 represents 0% ink coverage and a value of 255 represents 100%
	ink coverage for that component, but see DotRange below. The InkNames field
	should not exist when InkSet=1. */
	InkSetCMYK = 1

	// InkSetNotCMYK
	/* See the InkNames field for a description of the inks to be used. */
	InkSetNotCMYK = 2
)

// ExtraSamples.
const (
	ExtraSamplesUnspecifiedData = 0

	// ExtraSampleAlphaDataPreMultipliedColor
	/* Associated alpha data is opacity information; it is fully described in
	Section 21 */
	ExtraSampleAlphaDataPreMultipliedColor = 1

	// ExtraSamplesUnassociatedAlphaData
	/* Unassociated alpha data is transparency information that logically
	exists independent of an image; it is commonly called a soft matte. */
	ExtraSamplesUnassociatedAlphaData = 2
)

// JPEGProc.
const (
	JPEGProcBaselineSequentialProcess    = 1
	JPEGProcLosslessProcessHuffmanCoding = 14
)
