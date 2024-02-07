package spiralmatrix

import "math"

type psrow int
type pscol int

type spirral struct {
	matrix [][]int 
	size int
}

func NewSpirral(size int) *spirral {
	matrix:=make([][]int, size)
	for i:=0; i<size; i++ {
		line:=make([]int, size)
		matrix[i] = append(matrix[i], line...)
	}

	return &spirral{size: size, matrix: matrix }
}

func (p *spirral) GetMatrix() [][]int {
	return p.matrix
}

func (p *spirral) BuildMatrix() {
	var col pscol = -1
	var row psrow = 0
	cellValue:=1
	placesToMove:=p.size
	angle:=float64(0)
	turns:=1
	for placesToMove > 0 {
		vWay, hWay:=math.Sin(angle)*-1, math.Cos(angle)
		for i:=0; i < placesToMove; i++ {
			row+=psrow(vWay)
			col+=pscol(hWay)
			p.matrix[int(row)][int(col)]=cellValue
			cellValue++
		}
		angle-=math.Pi/2
		turns--
		if turns == 0 {
			placesToMove--
			turns=2
		}
	}
}

func SpiralMatrix(size int) [][]int {
	sm:=NewSpirral(size)
	sm.BuildMatrix()
	return sm.GetMatrix()
}
