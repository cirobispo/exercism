package series

func All(n int, s string) []string {
	textSize:=len(s)

	result :=make([]string, 0, 1)
	for i:=0; i+n <= textSize; i++ {
		result = append(result, s[i:i+n])
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[0:0+n]
}
