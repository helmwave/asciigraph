package dgraph

import (
	"github.com/helmwave/asciigraph/ascii"
	"github.com/helmwave/asciigraph/core"
)

func DrawGraph(list []core.NodeInput, o ascii.DrawOptions) (canvas *ascii.Canvas, err error) {
	g, err := core.NewGraph(list)
	if err != nil {
		return
	}
	mtx, err := g.Traverse()
	if err != nil {
		return
	}
	canvas, err = ascii.DrawAsciiMatrix(mtx, o)
	if err != nil {
		return
	}
	return
}
