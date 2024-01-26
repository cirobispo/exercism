package pascal

func Triangle(n int) [][]int {
	result:=make([][]int, 0)

	for i:=0; i<n; i++ {
		result = append(result, make([]int, 0))
		linha:=make([]int, 0, n)
		for j:=0; j<(i+1); j++ {
			b4:=0
			if i-1 >=0 && j-1 >=0 {
				b4=result[i-1][j-1]
			}
			ac:=0
			if (i-1) >=0 && j <=(i-1) {
				ac=result[i-1][j]
			}
			value:=b4+ac
			if value == 0 {
				value = 1
			}
			linha = append(linha, value)
		}
		result[i] = append(result[i], linha...)
	}

	return result
}
