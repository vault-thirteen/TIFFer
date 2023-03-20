package ifd

type Statistics struct {
	// KnownTagsCount is the number of tags which have a known name, i.e.
	// textual alias.
	KnownTagsCount int

	// UnKnownTagsCount is the number of tags which have no known name, i.e.
	// no textual alias.
	UnKnownTagsCount int

	// Those tags who have no type rule are automatically considered as valid
	// as stated in the TIFF 6.0 Specification. To count those "shadow" tags,
	// we use following counters.
	CountOfTagsWithRegisteredType   int
	CountOfTagsWithUnRegisteredType int
}
