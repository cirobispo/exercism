package connect

import (
	"errors"
	"fmt"
	_ "fmt"
	"math"
)

type HexGame struct {
	Height, Width int
	Cells []HexCell 
}

type HexCell struct {
	Line, Column int
	Value rune
}

func New(lines []string) HexGame {
	lineCount:=len(lines)
	cells:=make([]HexCell, 0, lineCount * lineCount)
	width:=func() int {
		if lineCount > 0 {
			return len(lines[0])
		}
		return 0
	}()

	if lineCount > 0 {
		for i:=range lines {
			line:=lines[i]		
			if cSize:=len(line); cSize > 0 {
				// deixo a coisa em losango e adiciono B nas areas cinzas. Antes estava como um retangulo
				for pc:=0; pc<i; pc++ {
					cells = append(cells, HexCell{Line: i, Column: pc, Value: 'N' })
				}

				for c:=range line {
					r:=rune(line[c])
					cells = append(cells, HexCell{Line: i, Column: c+i, Value: r })
				}

				for pc:=width+i; pc<2*width;pc++ {
					cells = append(cells, HexCell{Line: i, Column: pc, Value: 'N' })
				}
			}
		}
	}

	return HexGame{Height: lineCount, Width: width, Cells: cells}
}

func PrintMapCells(hg HexGame) {
	lastLine:=0;
	for i:=range hg.Cells {
		cell:=hg.Cells[i]
		if cell.Line != lastLine {
			fmt.Println()
			lastLine=cell.Line
		}

		fmt.Print(string(cell.Value), " ")
	}
	fmt.Println()
}

func (h HexGame) IsOverTheBoard(line, column int) bool {
	otl:=line >= 0 && line < h.Height
	otc:=column >= 0 && column < 2*h.Width

	return otl && otc
}

func(h HexGame) getValue(line, column int) (rune) {
	index:= line * (2*h.Width)
	index+=column

	return h.Cells[index].Value
}

func (h HexGame) TryGetValue(line, column int) (rune, error) {
	if h.IsOverTheBoard(line, column) {
		return h.getValue(line, column), nil
	}
	return '?', errors.New("position out of bounds")
}

func (h HexGame) EdgeCells() []HexCell {
	result:=make([]HexCell, 0, h.Height+h.Width)

	line:=h.Cells[1:h.Width]
	for i:= range line {
		if line[i].Value != '.' {
			result = append(result, line[i])
		}
	}

	size:=(h.Height*(h.Width*2))
	for i,k:=0, 0; i<size;i+=(2*h.Width) {
		cell:=h.Cells[i+k]
		if cell.Value != '.' {
			result = append(result, cell)
		}
		k++
	}

	return result
}

func (h HexGame) IsEdge(cell HexCell) bool {
	lEdge:=cell.Line == 0 || cell.Line == h.Height-1
	cEdge:=cell.Column-cell.Line == 0 || cell.Column-cell.Line == h.Width-1

	return lEdge || cEdge
}

func (h HexGame) IsParentOpposite(cell, parent HexCell) bool {
	vertical:=parent.Column > 0 && parent.Line == 0
	horizontal:=parent.Line > 0 && parent.Column-parent.Line == 0

	if !vertical && !horizontal {
		return (cell.Column == h.Width-1 + cell.Line) ||
			(cell.Line == h.Height-1)
	} else if vertical || horizontal {
		return (vertical && cell.Line == h.Height-1) || 
			(horizontal && cell.Column == h.Width-1 + cell.Line)
	}

	return false
}

func (h HexGame) GetNextPos(line, column int, r rune, prior []HexCell) []HexCell {
	exists:=func (cell HexCell, priorCells []HexCell) bool {
		result:=false
		for i:=range priorCells {
			if cell.Line == priorCells[i].Line && cell.Column == priorCells[i].Column {
				result=true
				break
			}
		}
		return result
	}
	
	result:=make([]HexCell, 0)
	for i, piRad:=float64(0), float64(0); i<360; i+=45 {
		piRad=(2*math.Pi*i)/360
		fc:=int(math.Round(math.Cos(piRad)))
		fl:=int(math.Round(math.Sin(piRad)))*-1

		if h.IsOverTheBoard(line+fl, column+fc) {
			value:=h.getValue(line+fl, column+fc)
			if value == r {
				if !exists(HexCell{Line: line+fl, Column: column+fc }, prior) {
					result = append(result, HexCell{Line: line+fl, Column: column+fc, Value: r })
				}
			}
		}
	}

	return result
}

func (h HexGame) IsWinningCell(parent HexCell, priorCells []HexCell) (HexCell, bool) {
	var result HexCell
	var found bool

	localPriorCells:=make([]HexCell, 0)
	localPriorCells = append(localPriorCells, priorCells...)
	
	vizinhos:=h.GetNextPos(parent.Line, parent.Column, parent.Value, localPriorCells)
	if vSize:=len(vizinhos); vSize > 0 {
		for i:=range vizinhos {
			vizinho:=vizinhos[i]
			if index:=len(localPriorCells)-1; index == -1 || index > -1 && localPriorCells[index] !=parent {
				localPriorCells = append(localPriorCells, parent)
			}
			result, found=h.IsWinningCell(vizinho, localPriorCells)
			if found {
				break
			}
		}
	} else {
		isEdge:=h.IsEdge(parent)
		if isEdge {
			var firstParent HexCell
			if len(localPriorCells) > 0 {
				firstParent=localPriorCells[0]
			}
			isOpposite:=h.IsParentOpposite(parent, firstParent)
			if found=isOpposite; found {
				result = parent
			}
		}
	}

	return result,found
}

func ResultOf(lines []string) (string, error) {
	hg:=New(lines)
	result:=""
	var err error
	if hg.Height > 0 {
		//caso em que há somente 1 celula.
		if hg.Height == 1 && hg.Width == 1 {
			return string(hg.Cells[0].Value), nil
		}

		edges:=hg.EdgeCells()
		priorCells:=make([]HexCell, 0)
		for i:=range edges {
			edge:=edges[i]
			data,found:=hg.IsWinningCell(edge, priorCells)
			if found {
				result=string(data.Value)
				break
			}
		}

		return result, err
	}

	return result, errors.New("empty board")
}
