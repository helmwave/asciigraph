package ascii

type DrawOptions struct {
	CellHeight   int
	MinCellWidth int
	MaxWidth     int
	Padding      int
}

var Default = &DrawOptions{
	CellHeight:   3,
	MinCellWidth: 3,
	MaxWidth:     18,
	Padding:      1,
}
