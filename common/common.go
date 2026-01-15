package common

type Location struct {
	_sourceOffset int
	_tokenLength  int

	Line   int
	Column int
}

func NewLocation(Line int, Column int, _sourceOffset int, _tokenLength int) *Location {
	loc := new(Location)

	loc.Line = Line
	loc.Column = Column

	loc._sourceOffset = _sourceOffset
	loc._tokenLength = _tokenLength

	return loc
}
