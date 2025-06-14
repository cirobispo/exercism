package rectangles

import (
	"fmt"
	"strings"
)

type Orientation int

const (
	voVer Orientation = 0
	voHor Orientation = iota
)

type edge struct {
	row, col int
}

type vertice struct {
	edgeA, edgeB edge
}

type box struct {
	topLeft, bottomRight edge
}

type board struct {
	rows, cols int
	diagram    []string
	vertices   map[int]*map[string]*vertice
}

func (o Orientation) toString() string {
	result := "H"
	if o == voVer {
		result = "V"
	}
	return result
}

func (v vertice) Orientation() Orientation {
	if v.edgeA.row == v.edgeB.row {
		return voHor
	}

	return voVer
}

func (v vertice) Size() int {
	if v.Orientation() == voVer {
		return v.edgeB.row - v.edgeA.row
	}

	return v.edgeB.col - v.edgeA.col
}

func (v vertice) toString() string {
	dp := fmt.Sprintf("%d,%d", v.edgeA.row, v.edgeA.col)
	if v.Orientation() == voVer {
		dp = fmt.Sprintf("%d,%d", v.edgeA.col, v.edgeA.row)
	}

	return fmt.Sprintf("%s,%s,%d", v.Orientation().toString(), dp, v.Size())
}

func (v vertice) moveTo(pos int) vertice {
	if v.Orientation() == voVer {
		return vertice{edge{row: v.edgeA.row, col: pos}, edge{row: v.edgeB.row, col: pos}}
	}
	return vertice{edge{row: pos, col: v.edgeA.col}, edge{row: pos, col: v.edgeB.col}}
}

func newVerticeByEdge(edgeA, edgeB edge) *vertice {
	if edgeA.row > edgeB.row || edgeA.col > edgeB.col {
		return &vertice{edgeA: edgeB, edgeB: edgeA}
	}

	return &vertice{edgeA: edgeA, edgeB: edgeB}
}

func newVerticeBySize(displ, pos, size int, ori Orientation) *vertice {
	edgeA := edge{row: displ, col: pos}
	edgeB := edge{row: displ, col: pos + size}
	if ori == voVer {
		edgeA.col, edgeA.row = displ, pos
		edgeB.col, edgeB.row = displ, pos+size
	}

	return &vertice{edgeA: edgeA, edgeB: edgeB}
}

func displacedSize(vA, vB *vertice) int {
	if vA.Orientation() == voVer {
		return vB.edgeA.col - vA.edgeA.col
	}

	return vB.edgeA.row - vA.edgeA.row
}

func (b *board) collectVertices(ori Orientation, chars []rune) {
	if (len(b.diagram) < 1) || (len(b.diagram) > 0 && len(b.diagram[0]) < 2) {
		return
	}

	buildLine := func(data *[]int, displ int, ori Orientation) {
		for h := 0; h < len((*data)); h++ {
			for hh := h + 1; hh < len((*data)); hh++ {
				newV := newVerticeBySize(displ, (*data)[h], (*data)[hh]-(*data)[h], ori)
				m, f := b.vertices[newV.Size()]
				if !f {
					g := make(map[string]*vertice)
					b.vertices[newV.Size()] = &g
					m = b.vertices[newV.Size()]
				}
				(*m)[newV.toString()] = newV
			}
		}

		if len(*data) > 0 {
			*data = make([]int, 0)
		}
	}

	getChar := func(row, col int, ori Orientation) rune {
		if ori == voVer {
			row, col = col, row
		}
		return rune(b.diagram[row][col])
	}

	rows, cols := len(b.diagram), len(b.diagram[0])
	if ori == voVer {
		rows, cols = len(b.diagram[0]), len(b.diagram)
	}

	allPos := make([]int, 0)
	for d := 0; d < rows; d++ {
		for p := 0; p < cols; p++ {
			ch := getChar(d, p, ori)
			if !strings.ContainsRune(string(chars), ch) {
				if len(allPos) > 0 {
					buildLine(&allPos, d, ori)
				}
				continue
			}

			if ch == chars[0] {
				allPos = append(allPos, p)
			}
		}

		buildLine(&allPos, d, ori)
	}
}

func (b *board) CollectAllVertices() {
	b.vertices = make(map[int]*map[string]*vertice)
	b.collectVertices(voVer, []rune{'+', '|'})
	b.collectVertices(voHor, []rune{'+', '-'})
}

func (b board) placedVertices(v *vertice, vertices *map[string]*vertice) *[][]*vertice {
	result := make([][]*vertice, 0)

	max := b.rows
	if v.Orientation() == voVer {
		max = b.cols
	}

	vid := v.toString()
	for i := 0; i < max; i++ {
		id := v.moveTo(i).toString()
		if id == vid {
			continue
		}

		if found := (*vertices)[id]; found != nil {
			item := []*vertice{v, found}
			if (v.Orientation() == voHor && v.edgeA.row > found.edgeA.row) ||
				(v.Orientation() == voVer && v.edgeA.col > found.edgeA.col) {
				item[0], item[1] = item[1], item[0]
			}
			result = append(result, item)
		}
	}

	return &result
}

func (b board) displacedVertices(allPlaceds *[][]*vertice) *[][]*vertice {
	result := make([][]*vertice, 0)

	for i := 0; i < len((*allPlaceds)); i++ {
		placed := (*allPlaceds)[i]

		subResult := make([]*vertice, 0, 2)
		if len(placed) == 2 {
			dSize := displacedSize(placed[0], placed[1])
			vertices := b.vertices[dSize]

			findSert := func(edgeA, edgeB edge) {
				id := (newVerticeByEdge(edgeA, edgeB)).toString()
				if data, found := (*vertices)[id]; found {
					subResult = append(subResult, data)
				}
			}

			findSert(placed[0].edgeA, placed[1].edgeA)
			findSert(placed[0].edgeB, placed[1].edgeB)
		}
		result = append(result, subResult)
	}

	return &result
}

func (b box) toString() string {
	result := fmt.Sprintf("(%d, %d) (%d, %d)", b.topLeft.row, b.topLeft.col, b.bottomRight.row, b.bottomRight.col)
	return result
}

func (b box) Height() int {
	return (b.bottomRight.row - b.topLeft.row)
}

func (b box) Width() int {
	return (b.bottomRight.col - b.topLeft.col)
}

func (b box) Area() int {
	return b.Height() * b.Width()
}

func verticesToBox(placed, displaced []*vertice) *box {
	if len(placed) != 2 || len(displaced) != 2 {
		return nil
	}

	result := &box{topLeft: placed[0].edgeA, bottomRight: displaced[1].edgeB}

	return result
}

func (b *board) buildBoxes() *map[string]*box {
	boxes := make(map[string]*box)
	for s := range b.vertices {
		vertices := b.vertices[s]
		for _, vv := range *vertices {
			allPlaced := b.placedVertices(vv, vertices)
			allDisplaced := b.displacedVertices(allPlaced)
			for p := 0; p < len((*allPlaced)); p++ {
				placed := (*allPlaced)[p]
				displaced := (*allDisplaced)[p]
				box := verticesToBox(placed, displaced)
				if box != nil {
					if _, found := boxes[box.toString()]; !found {
						boxes[box.toString()] = box
					}
				}
			}
		}
	}

	return &boxes
}

func newBoard(diagram *[]string) board {
	rows := len(*diagram)
	cols := 0
	if rows > 0 {
		cols = len((*diagram)[0])
	}

	result := board{rows: rows, cols: cols, diagram: *diagram, vertices: make(map[int]*map[string]*vertice)}
	return result
}

func Count(diagram []string) int {
	b := newBoard(&diagram)
	b.CollectAllVertices()
	boxes := b.buildBoxes()

	return len(*boxes)
}
