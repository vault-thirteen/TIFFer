package tag

import bt "github.com/vault-thirteen/TIFFer/models/basic-types"

// Tags as per TIFF 6.0 Specification and some additional tags.
//
// TIFF Tag Reference, Private TIFF Tags:
// https://www.awaresystems.be/imaging/tiff/tiffs/private.html
const (
	NewSubfileType              = 254
	SubfileType                 = 255
	ImageWidth                  = 256
	ImageLength                 = 257 // Length is Height.
	BitsPerSample               = 258
	Compression                 = 259
	PhotometricInterpretation   = 262
	Threshholding               = 263
	CellWidth                   = 264
	CellLength                  = 265
	FillOrder                   = 266
	DocumentName                = 269
	ImageDescription            = 270
	Make                        = 271
	Model                       = 272
	StripOffsets                = 273
	Orientation                 = 274
	SamplesPerPixel             = 277
	RowsPerStrip                = 278
	StripByteCounts             = 279
	MinSampleValue              = 280
	MaxSampleValue              = 281
	XResolution                 = 282
	YResolution                 = 283
	PlanarConfiguration         = 284
	PageName                    = 285
	XPosition                   = 286
	YPosition                   = 287
	FreeOffsets                 = 288
	FreeByteCounts              = 289
	GrayResponseUnit            = 290
	GrayResponseCurve           = 291
	T4Options                   = 292
	T6Options                   = 293
	ResolutionUnit              = 296
	PageNumber                  = 297
	TransferFunction            = 301
	Software                    = 305
	DateTime                    = 306
	Artist                      = 315
	HostComputer                = 316
	Predictor                   = 317
	WhitePoint                  = 318
	PrimaryChromaticities       = 319
	ColorMap                    = 320
	HalftoneHints               = 321
	TileWidth                   = 322
	TileLength                  = 323
	TileOffsets                 = 324
	TileByteCounts              = 325
	BadFaxLines                 = 326
	CleanFaxData                = 327
	ConsecutiveBadFaxLines      = 328
	SubIFDs                     = 330
	InkSet                      = 332
	InkNames                    = 333
	NumberOfInks                = 334
	DotRange                    = 336
	TargetPrinter               = 337
	ExtraSamples                = 338
	SampleFormat                = 339
	SMinSampleValue             = 340
	SMaxSampleValue             = 341
	TransferRange               = 342
	ClipPath                    = 343
	XClipPathUnits              = 344
	YClipPathUnits              = 345
	Indexed                     = 346
	JPEGTables                  = 347
	OPIProxy                    = 351
	GlobalParametersIFD         = 400
	ProfileType                 = 401
	FaxProfile                  = 402
	CodingMethods               = 403
	VersionYear                 = 404
	ModeNumber                  = 405
	Decode                      = 433
	DefaultImageColor           = 434
	JPEGProc                    = 512
	JPEGInterchangeFormat       = 513
	JPEGInterchangeFormatLength = 514
	JPEGRestartInterval         = 515
	JPEGLosslessPredictors      = 517
	JPEGPointTransforms         = 518
	JPEGQTables                 = 519
	JPEGDCTables                = 520
	JPEGACTables                = 521
	YCbCrCoefficients           = 529
	YCbCrSubSampling            = 530
	YCbCrPositioning            = 531
	ReferenceBlackWhite         = 532
	StripRowCounts              = 559
	XMP                         = 700
)

// Private Tags.
const (
	ImageID                      = 32781
	WangAnnotation               = 32932
	Copyright                    = 33432
	MDFile                       = 33445
	MDScalePixel                 = 33446
	MDColorTable                 = 33447
	MDLabName                    = 33448
	MDSampleInfo                 = 33449
	MDPrepDate                   = 33450
	MDPrepTime                   = 33451
	MDFileUnits                  = 33452
	ModelPixelScale              = 33550
	IPTC                         = 33723
	INGRPacketData               = 33918
	INGRFlagRegisters            = 33919
	IrasBTransformationMatrix    = 33920
	ModelTiepoint                = 33922
	ModelTransformation          = 34264
	Photoshop                    = 34377
	ExifIFD                      = 34665
	ICCProfile                   = 34675
	ImageLayer                   = 34732
	GeoKeyDirectory              = 34735
	GeoDoubleParams              = 34736
	GeoAsciiParams               = 34737
	GPSIFD                       = 34853
	HylaFAXFaxRecvParams         = 34908
	HylaFAXFaxSubAddress         = 34909
	HylaFAXFaxRecvTime           = 34910
	ImageSourceData              = 37724
	InteroperabilityIFD          = 40965
	GDAL_METADATA                = 42112
	GDAL_NODATA                  = 42113
	OceScanjobDescription        = 50215
	OceApplicationSelector       = 50216
	OceIdentificationNumber      = 50217
	OceImageLogicCharacteristics = 50218
	DNGVersion                   = 50706
	DNGBackwardVersion           = 50707
	UniqueCameraModel            = 50708
	LocalizedCameraModel         = 50709
	CFAPlaneColor                = 50710
	CFALayout                    = 50711
	LinearizationTable           = 50712
	BlackLevelRepeatDim          = 50713
	BlackLevel                   = 50714
	BlackLevelDeltaH             = 50715
	BlackLevelDeltaV             = 50716
	WhiteLevel                   = 50717
	DefaultScale                 = 50718
	DefaultCropOrigin            = 50719
	DefaultCropSize              = 50720
	ColorMatrix1                 = 50721
	ColorMatrix2                 = 50722
	CameraCalibration1           = 50723
	CameraCalibration2           = 50724
	ReductionMatrix1             = 50725
	ReductionMatrix2             = 50726
	AnalogBalance                = 50727
	AsShotNeutral                = 50728
	AsShotWhiteXY                = 50729
	BaselineExposure             = 50730
	BaselineNoise                = 50731
	BaselineSharpness            = 50732
	BayerGreenSplit              = 50733
	LinearResponseLimit          = 50734
	CameraSerialNumber           = 50735
	LensInfo                     = 50736
	ChromaBlurRadius             = 50737
	AntiAliasStrength            = 50738
	ShadowScale                  = 50739
	DNGPrivateData               = 50740
	MakerNoteSafety              = 50741
	CalibrationIlluminant1       = 50778
	CalibrationIlluminant2       = 50779
	BestQualityScale             = 50780
	RawDataUniqueID              = 50781
	AliasLayerMetadata           = 50784
	OriginalRawFileName          = 50827
	OriginalRawFileData          = 50828
	ActiveArea                   = 50829
	MaskedAreas                  = 50830
	AsShotICCProfile             = 50831
	AsShotPreProfileMatrix       = 50832
	CurrentICCProfile            = 50833
	CurrentPreProfileMatrix      = 50834
	ColorimetricReference        = 50879
	TIFF_RSID                    = 50908
	GEO_METADATA                 = 50909
	CameraCalibrationSignature   = 50931
	ProfileCalibrationSignature  = 50932
	ExtraCameraProfiles          = 50933
	AsShotProfileName            = 50934
	NoiseReductionApplied        = 50935
	ProfileName                  = 50936
	ProfileHueSatMapDims         = 50937
	ProfileHueSatMapData1        = 50938
	ProfileHueSatMapData2        = 50939
	ProfileToneCurve             = 50940
	ProfileEmbedPolicy           = 50941
	ProfileCopyright             = 50942
	ForwardMatrix1               = 50964
	ForwardMatrix2               = 50965
	PreviewApplicationName       = 50966
	PreviewApplicationVersion    = 50967
	PreviewSettingsName          = 50968
	PreviewSettingsDigest        = 50969
	PreviewColorSpace            = 50970
	PreviewDateTime              = 50971
	RawImageDigest               = 50972
	OriginalRawFileDigest        = 50973
	SubTileBlockSize             = 50974
	RowInterleaveFactor          = 50975
	ProfileLookTableDims         = 50981
	ProfileLookTableData         = 50982
	OpcodeList1                  = 51008
	OpcodeList2                  = 51009
	OpcodeList3                  = 51022
	NoiseProfile                 = 51041
	OriginalDefaultFinalSize     = 51089
	OriginalBestQualityFinalSize = 51090
	OriginalDefaultCropSize      = 51091
	ProfileHueSatMapEncoding     = 51107
	ProfileLookTableEncoding     = 51108
	BaselineExposureOffset       = 51109
	DefaultBlackRender           = 51110
	NewRawImageDigest            = 51111
	RawToPreviewGain             = 51112
	DefaultUserCrop              = 51125
	DepthFormat                  = 51177
	DepthNear                    = 51178
	DepthFar                     = 51179
	DepthUnits                   = 51180
	DepthMeasureType             = 51181
	EnhanceParams                = 51182
)

// EXIF Tags.
const (
	ExposureTime             = 33434
	FNumber                  = 33437
	ExposureProgram          = 34850
	SpectralSensitivity      = 34852
	ISOSpeedRatings          = 34855
	OECF                     = 34856
	ExifVersion              = 36864
	DateTimeOriginal         = 36867
	DateTimeDigitized        = 36868
	ComponentsConfiguration  = 37121
	CompressedBitsPerPixel   = 37122
	ShutterSpeedValue        = 37377
	ApertureValue            = 37378
	BrightnessValue          = 37379
	ExposureBiasValue        = 37380
	MaxApertureValue         = 37381
	SubjectDistance          = 37382
	MeteringMode             = 37383
	LightSource              = 37384
	Flash                    = 37385
	FocalLength              = 37386
	SubjectArea              = 37396
	MakerNote                = 37500
	UserComment              = 37510
	SubsecTime               = 37520
	SubsecTimeOriginal       = 37521
	SubsecTimeDigitized      = 37522
	FlashpixVersion          = 40960
	ColorSpace               = 40961
	PixelXDimension          = 40962
	PixelYDimension          = 40963
	RelatedSoundFile         = 40964
	FlashEnergy              = 41483
	SpatialFrequencyResponse = 41484
	FocalPlaneXResolution    = 41486
	FocalPlaneYResolution    = 41487
	FocalPlaneResolutionUnit = 41488
	SubjectLocation          = 41492
	ExposureIndex            = 41493
	SensingMethod            = 41495
	FileSource               = 41728
	SceneType                = 41729
	CFAPattern               = 41730
	CustomRendered           = 41985
	ExposureMode             = 41986
	WhiteBalance             = 41987
	DigitalZoomRatio         = 41988
	FocalLengthIn35mmFilm    = 41989
	SceneCaptureType         = 41990
	GainControl              = 41991
	Contrast                 = 41992
	Saturation               = 41993
	Sharpness                = 41994
	DeviceSettingDescription = 41995
	SubjectDistanceRange     = 41996
	ImageUniqueID            = 42016
)

// GPS Tags.
const (
	GPSVersionID        = 0
	GPSLatitudeRef      = 1
	GPSLatitude         = 2
	GPSLongitudeRef     = 3
	GPSLongitude        = 4
	GPSAltitudeRef      = 5
	GPSAltitude         = 6
	GPSTimeStamp        = 7
	GPSSatellites       = 8
	GPSStatus           = 9
	GPSMeasureMode      = 10
	GPSDOP              = 11
	GPSSpeedRef         = 12
	GPSSpeed            = 13
	GPSTrackRef         = 14
	GPSTrack            = 15
	GPSImgDirectionRef  = 16
	GPSImgDirection     = 17
	GPSMapDatum         = 18
	GPSDestLatitudeRef  = 19
	GPSDestLatitude     = 20
	GPSDestLongitudeRef = 21
	GPSDestLongitude    = 22
	GPSDestBearingRef   = 23
	GPSDestBearing      = 24
	GPSDestDistanceRef  = 25
	GPSDestDistance     = 26
	GPSProcessingMethod = 27
	GPSAreaInformation  = 28
	GPSDateStamp        = 29
	GPSDifferential     = 30
)

// NameUnknown is the name used for unknown tags.
const NameUnknown = "Unknown"

// Tag is the tag of a Directory Entry.
type Tag = bt.Word

var humanReadableTagNames = map[Tag]string{
	NewSubfileType:              "NewSubfileType",
	SubfileType:                 "SubfileType",
	ImageWidth:                  "ImageWidth",
	ImageLength:                 "ImageLength",
	BitsPerSample:               "BitsPerSample",
	Compression:                 "Compression",
	PhotometricInterpretation:   "PhotometricInterpretation",
	Threshholding:               "Threshholding",
	CellWidth:                   "CellWidth",
	CellLength:                  "CellLength",
	FillOrder:                   "FillOrder",
	DocumentName:                "DocumentName",
	ImageDescription:            "ImageDescription",
	Make:                        "Make",
	Model:                       "Model",
	StripOffsets:                "StripOffsets",
	Orientation:                 "Orientation",
	SamplesPerPixel:             "SamplesPerPixel",
	RowsPerStrip:                "RowsPerStrip",
	StripByteCounts:             "StripByteCounts",
	MinSampleValue:              "MinSampleValue",
	MaxSampleValue:              "MaxSampleValue",
	XResolution:                 "XResolution",
	YResolution:                 "YResolution",
	PlanarConfiguration:         "PlanarConfiguration",
	PageName:                    "PageName",
	XPosition:                   "XPosition",
	YPosition:                   "YPosition",
	FreeOffsets:                 "FreeOffsets",
	FreeByteCounts:              "FreeByteCounts",
	GrayResponseUnit:            "GrayResponseUnit",
	GrayResponseCurve:           "GrayResponseCurve",
	T4Options:                   "T4Options",
	T6Options:                   "T6Options",
	ResolutionUnit:              "ResolutionUnit",
	PageNumber:                  "PageNumber",
	TransferFunction:            "TransferFunction",
	Software:                    "Software",
	DateTime:                    "DateTime",
	Artist:                      "Artist",
	HostComputer:                "HostComputer",
	Predictor:                   "Predictor",
	WhitePoint:                  "WhitePoint",
	PrimaryChromaticities:       "PrimaryChromaticities",
	ColorMap:                    "ColorMap",
	HalftoneHints:               "HalftoneHints",
	TileWidth:                   "TileWidth",
	TileLength:                  "TileLength",
	TileOffsets:                 "TileOffsets",
	TileByteCounts:              "TileByteCounts",
	BadFaxLines:                 "BadFaxLines",
	CleanFaxData:                "CleanFaxData",
	ConsecutiveBadFaxLines:      "ConsecutiveBadFaxLines",
	SubIFDs:                     "SubIFDs",
	InkSet:                      "InkSet",
	InkNames:                    "InkNames",
	NumberOfInks:                "NumberOfInks",
	DotRange:                    "DotRange",
	TargetPrinter:               "TargetPrinter",
	ExtraSamples:                "ExtraSamples",
	SampleFormat:                "SampleFormat",
	SMinSampleValue:             "SMinSampleValue",
	SMaxSampleValue:             "SMaxSampleValue",
	TransferRange:               "TransferRange",
	ClipPath:                    "ClipPath",
	XClipPathUnits:              "XClipPathUnits",
	YClipPathUnits:              "YClipPathUnits",
	Indexed:                     "Indexed",
	JPEGTables:                  "JPEGTables",
	OPIProxy:                    "OPIProxy",
	GlobalParametersIFD:         "GlobalParametersIFD",
	ProfileType:                 "ProfileType",
	FaxProfile:                  "FaxProfile",
	CodingMethods:               "CodingMethods",
	VersionYear:                 "VersionYear",
	ModeNumber:                  "ModeNumber",
	Decode:                      "Decode",
	DefaultImageColor:           "DefaultImageColor",
	JPEGProc:                    "JPEGProc",
	JPEGInterchangeFormat:       "JPEGInterchangeFormat",
	JPEGInterchangeFormatLength: "JPEGInterchangeFormatLength",
	JPEGRestartInterval:         "JPEGRestartInterval",
	JPEGLosslessPredictors:      "JPEGLosslessPredictors",
	JPEGPointTransforms:         "JPEGPointTransforms",
	JPEGQTables:                 "JPEGQTables",
	JPEGDCTables:                "JPEGDCTables",
	JPEGACTables:                "JPEGACTables",
	YCbCrCoefficients:           "YCbCrCoefficients",
	YCbCrSubSampling:            "YCbCrSubSampling",
	YCbCrPositioning:            "YCbCrPositioning",
	ReferenceBlackWhite:         "ReferenceBlackWhite",
	StripRowCounts:              "StripRowCounts",
	XMP:                         "XMP",

	// Private Tags.
	ImageID:                      "ImageID",
	WangAnnotation:               "WangAnnotation",
	Copyright:                    "Copyright",
	MDFile:                       "MDFile",
	MDScalePixel:                 "MDScalePixel",
	MDColorTable:                 "MDColorTable",
	MDLabName:                    "MDLabName",
	MDSampleInfo:                 "MDSampleInfo",
	MDPrepDate:                   "MDPrepDate",
	MDPrepTime:                   "MDPrepTime",
	MDFileUnits:                  "MDFileUnits",
	ModelPixelScale:              "ModelPixelScale",
	IPTC:                         "IPTC",
	INGRPacketData:               "INGRPacketData",
	INGRFlagRegisters:            "INGRFlagRegisters",
	IrasBTransformationMatrix:    "IrasBTransformationMatrix",
	ModelTiepoint:                "ModelTiepoint",
	ModelTransformation:          "ModelTransformation",
	Photoshop:                    "Photoshop",
	ExifIFD:                      "ExifIFD",
	ICCProfile:                   "ICCProfile",
	ImageLayer:                   "ImageLayer",
	GeoKeyDirectory:              "GeoKeyDirectory",
	GeoDoubleParams:              "GeoDoubleParams",
	GeoAsciiParams:               "GeoAsciiParams",
	GPSIFD:                       "GPSIFD",
	HylaFAXFaxRecvParams:         "HylaFAXFaxRecvParams",
	HylaFAXFaxSubAddress:         "HylaFAXFaxSubAddress",
	HylaFAXFaxRecvTime:           "HylaFAXFaxRecvTime",
	ImageSourceData:              "ImageSourceData",
	InteroperabilityIFD:          "InteroperabilityIFD",
	GDAL_METADATA:                "GDAL_METADATA",
	GDAL_NODATA:                  "GDAL_NODATA",
	OceScanjobDescription:        "OceScanjobDescription",
	OceApplicationSelector:       "OceApplicationSelector",
	OceIdentificationNumber:      "OceIdentificationNumber",
	OceImageLogicCharacteristics: "OceImageLogicCharacteristics",
	DNGVersion:                   "DNGVersion",
	DNGBackwardVersion:           "DNGBackwardVersion",
	UniqueCameraModel:            "UniqueCameraModel",
	LocalizedCameraModel:         "LocalizedCameraModel",
	CFAPlaneColor:                "CFAPlaneColor",
	CFALayout:                    "CFALayout",
	LinearizationTable:           "LinearizationTable",
	BlackLevelRepeatDim:          "BlackLevelRepeatDim",
	BlackLevel:                   "BlackLevel",
	BlackLevelDeltaH:             "BlackLevelDeltaH",
	BlackLevelDeltaV:             "BlackLevelDeltaV",
	WhiteLevel:                   "WhiteLevel",
	DefaultScale:                 "DefaultScale",
	DefaultCropOrigin:            "DefaultCropOrigin",
	DefaultCropSize:              "DefaultCropSize",
	ColorMatrix1:                 "ColorMatrix1",
	ColorMatrix2:                 "ColorMatrix2",
	CameraCalibration1:           "CameraCalibration1",
	CameraCalibration2:           "CameraCalibration2",
	ReductionMatrix1:             "ReductionMatrix1",
	ReductionMatrix2:             "ReductionMatrix2",
	AnalogBalance:                "AnalogBalance",
	AsShotNeutral:                "AsShotNeutral",
	AsShotWhiteXY:                "AsShotWhiteXY",
	BaselineExposure:             "BaselineExposure",
	BaselineNoise:                "BaselineNoise",
	BaselineSharpness:            "BaselineSharpness",
	BayerGreenSplit:              "BayerGreenSplit",
	LinearResponseLimit:          "LinearResponseLimit",
	CameraSerialNumber:           "CameraSerialNumber",
	LensInfo:                     "LensInfo",
	ChromaBlurRadius:             "ChromaBlurRadius",
	AntiAliasStrength:            "AntiAliasStrength",
	ShadowScale:                  "ShadowScale",
	DNGPrivateData:               "DNGPrivateData",
	MakerNoteSafety:              "MakerNoteSafety",
	CalibrationIlluminant1:       "CalibrationIlluminant1",
	CalibrationIlluminant2:       "CalibrationIlluminant2",
	BestQualityScale:             "BestQualityScale",
	RawDataUniqueID:              "RawDataUniqueID",
	AliasLayerMetadata:           "AliasLayerMetadata",
	OriginalRawFileName:          "OriginalRawFileName",
	OriginalRawFileData:          "OriginalRawFileData",
	ActiveArea:                   "ActiveArea",
	MaskedAreas:                  "MaskedAreas",
	AsShotICCProfile:             "AsShotICCProfile",
	AsShotPreProfileMatrix:       "AsShotPreProfileMatrix",
	CurrentICCProfile:            "CurrentICCProfile",
	CurrentPreProfileMatrix:      "CurrentPreProfileMatrix",
	ColorimetricReference:        "ColorimetricReference",
	TIFF_RSID:                    "TIFF_RSID",
	GEO_METADATA:                 "GEO_METADATA",
	CameraCalibrationSignature:   "CameraCalibrationSignature",
	ProfileCalibrationSignature:  "ProfileCalibrationSignature",
	ExtraCameraProfiles:          "ExtraCameraProfiles",
	AsShotProfileName:            "AsShotProfileName",
	NoiseReductionApplied:        "NoiseReductionApplied",
	ProfileName:                  "ProfileName",
	ProfileHueSatMapDims:         "ProfileHueSatMapDims",
	ProfileHueSatMapData1:        "ProfileHueSatMapData1",
	ProfileHueSatMapData2:        "ProfileHueSatMapData2",
	ProfileToneCurve:             "ProfileToneCurve",
	ProfileEmbedPolicy:           "ProfileEmbedPolicy",
	ProfileCopyright:             "ProfileCopyright",
	ForwardMatrix1:               "ForwardMatrix1",
	ForwardMatrix2:               "ForwardMatrix2",
	PreviewApplicationName:       "PreviewApplicationName",
	PreviewApplicationVersion:    "PreviewApplicationVersion",
	PreviewSettingsName:          "PreviewSettingsName",
	PreviewSettingsDigest:        "PreviewSettingsDigest",
	PreviewColorSpace:            "PreviewColorSpace",
	PreviewDateTime:              "PreviewDateTime",
	RawImageDigest:               "RawImageDigest",
	OriginalRawFileDigest:        "OriginalRawFileDigest",
	SubTileBlockSize:             "SubTileBlockSize",
	RowInterleaveFactor:          "RowInterleaveFactor",
	ProfileLookTableDims:         "ProfileLookTableDims",
	ProfileLookTableData:         "ProfileLookTableData",
	OpcodeList1:                  "OpcodeList1",
	OpcodeList2:                  "OpcodeList2",
	OpcodeList3:                  "OpcodeList3",
	NoiseProfile:                 "NoiseProfile",
	OriginalDefaultFinalSize:     "OriginalDefaultFinalSize",
	OriginalBestQualityFinalSize: "OriginalBestQualityFinalSize",
	OriginalDefaultCropSize:      "OriginalDefaultCropSize",
	ProfileHueSatMapEncoding:     "ProfileHueSatMapEncoding",
	ProfileLookTableEncoding:     "ProfileLookTableEncoding",
	BaselineExposureOffset:       "BaselineExposureOffset",
	DefaultBlackRender:           "DefaultBlackRender",
	NewRawImageDigest:            "NewRawImageDigest",
	RawToPreviewGain:             "RawToPreviewGain",
	DefaultUserCrop:              "DefaultUserCrop",
	DepthFormat:                  "DepthFormat",
	DepthNear:                    "DepthNear",
	DepthFar:                     "DepthFar",
	DepthUnits:                   "DepthUnits",
	DepthMeasureType:             "DepthMeasureType",
	EnhanceParams:                "EnhanceParams",

	// EXIF Tags.
	ExposureTime:             "ExposureTime",
	FNumber:                  "FNumber",
	ExposureProgram:          "ExposureProgram",
	SpectralSensitivity:      "SpectralSensitivity",
	ISOSpeedRatings:          "ISOSpeedRatings",
	OECF:                     "OECF",
	ExifVersion:              "ExifVersion",
	DateTimeOriginal:         "DateTimeOriginal",
	DateTimeDigitized:        "DateTimeDigitized",
	ComponentsConfiguration:  "ComponentsConfiguration",
	CompressedBitsPerPixel:   "CompressedBitsPerPixel",
	ShutterSpeedValue:        "ShutterSpeedValue",
	ApertureValue:            "ApertureValue",
	BrightnessValue:          "BrightnessValue",
	ExposureBiasValue:        "ExposureBiasValue",
	MaxApertureValue:         "MaxApertureValue",
	SubjectDistance:          "SubjectDistance",
	MeteringMode:             "MeteringMode",
	LightSource:              "LightSource",
	Flash:                    "Flash",
	FocalLength:              "FocalLength",
	SubjectArea:              "SubjectArea",
	MakerNote:                "MakerNote",
	UserComment:              "UserComment",
	SubsecTime:               "SubsecTime",
	SubsecTimeOriginal:       "SubsecTimeOriginal",
	SubsecTimeDigitized:      "SubsecTimeDigitized",
	FlashpixVersion:          "FlashpixVersion",
	ColorSpace:               "ColorSpace",
	PixelXDimension:          "PixelXDimension",
	PixelYDimension:          "PixelYDimension",
	RelatedSoundFile:         "RelatedSoundFile",
	FlashEnergy:              "FlashEnergy",
	SpatialFrequencyResponse: "SpatialFrequencyResponse",
	FocalPlaneXResolution:    "FocalPlaneXResolution",
	FocalPlaneYResolution:    "FocalPlaneYResolution",
	FocalPlaneResolutionUnit: "FocalPlaneResolutionUnit",
	SubjectLocation:          "SubjectLocation",
	ExposureIndex:            "ExposureIndex",
	SensingMethod:            "SensingMethod",
	FileSource:               "FileSource",
	SceneType:                "SceneType",
	CFAPattern:               "CFAPattern",
	CustomRendered:           "CustomRendered",
	ExposureMode:             "ExposureMode",
	WhiteBalance:             "WhiteBalance",
	DigitalZoomRatio:         "DigitalZoomRatio",
	FocalLengthIn35mmFilm:    "FocalLengthIn35mmFilm",
	SceneCaptureType:         "SceneCaptureType",
	GainControl:              "GainControl",
	Contrast:                 "Contrast",
	Saturation:               "Saturation",
	Sharpness:                "Sharpness",
	DeviceSettingDescription: "DeviceSettingDescription",
	SubjectDistanceRange:     "SubjectDistanceRange",
	ImageUniqueID:            "ImageUniqueID",

	// GPS Tags.
	GPSVersionID:        "GPSVersionID",
	GPSLatitudeRef:      "GPSLatitudeRef",
	GPSLatitude:         "GPSLatitude",
	GPSLongitudeRef:     "GPSLongitudeRef",
	GPSLongitude:        "GPSLongitude",
	GPSAltitudeRef:      "GPSAltitudeRef",
	GPSAltitude:         "GPSAltitude",
	GPSTimeStamp:        "GPSTimeStamp",
	GPSSatellites:       "GPSSatellites",
	GPSStatus:           "GPSStatus",
	GPSMeasureMode:      "GPSMeasureMode",
	GPSDOP:              "GPSDOP",
	GPSSpeedRef:         "GPSSpeedRef",
	GPSSpeed:            "GPSSpeed",
	GPSTrackRef:         "GPSTrackRef",
	GPSTrack:            "GPSTrack",
	GPSImgDirectionRef:  "GPSImgDirectionRef",
	GPSImgDirection:     "GPSImgDirection",
	GPSMapDatum:         "GPSMapDatum",
	GPSDestLatitudeRef:  "GPSDestLatitudeRef",
	GPSDestLatitude:     "GPSDestLatitude",
	GPSDestLongitudeRef: "GPSDestLongitudeRef",
	GPSDestLongitude:    "GPSDestLongitude",
	GPSDestBearingRef:   "GPSDestBearingRef",
	GPSDestBearing:      "GPSDestBearing",
	GPSDestDistanceRef:  "GPSDestDistanceRef",
	GPSDestDistance:     "GPSDestDistance",
	GPSProcessingMethod: "GPSProcessingMethod",
	GPSAreaInformation:  "GPSAreaInformation",
	GPSDateStamp:        "GPSDateStamp",
	GPSDifferential:     "GPSDifferential",
}

// HumanReadableTagNames shows a list of all possible human-readable tag names.
// While Golang does not allow to make a variable constant or read-only like in
// C# language, we use a wrapper-function to show variables in a read-only
// manner.
func HumanReadableTagNames() map[Tag]string {
	return humanReadableTagNames
}

var tagsUsingSubIFDStyle = []Tag{
	SubIFDs,             // https://www.awaresystems.be/imaging/tiff/tifftags/subifds.html
	GlobalParametersIFD, // https://www.awaresystems.be/imaging/tiff/tifftags/globalparametersifd.html
	ExifIFD,             // https://www.awaresystems.be/imaging/tiff/tifftags/exififd.html
	GPSIFD,              // https://www.awaresystems.be/imaging/tiff/tifftags/gpsifd.html
	InteroperabilityIFD, // https://www.awaresystems.be/imaging/tiff/tifftags/interoperabilityifd.html
}

// TagsUsingSubIFDStyle shows a list of all possible tags which use sub-IFDs.
// While Golang does not allow to make a variable constant or read-only like in
// C# language, we use a wrapper-function to show variables in a read-only
// manner.
func TagsUsingSubIFDStyle() []Tag {
	return tagsUsingSubIFDStyle
}

// IsSubIFDTag tells whether the specified tag supports the 'sub-IFD' feature.
func IsSubIFDTag(t Tag) bool {
	for _, tt := range tagsUsingSubIFDStyle {
		if t == tt {
			return true
		}
	}
	return false
}
