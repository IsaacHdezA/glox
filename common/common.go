package common

type Location struct {
	_sourceOffset int
	_tokenLength  int

	line   int
	column int
}

func NewLocation(line int, column int, _sourceOffset int, _tokenLength int) *Location {
	loc := new(Location)

	loc.line = line
	loc.column = column

	loc._sourceOffset = _sourceOffset
	loc._tokenLength = _tokenLength

	return loc
}
