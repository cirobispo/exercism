package transpose

import "strings"

type BendSliceString [][]rune

func New(s []string) BendSliceString {
	if len(s) == 0 {
		return make([][]rune, 0)
	}

	strSizes := getAllSizes(s)
	for i := len(strSizes) - 1; i > 0; i-- {
		if strSizes[i-1] < strSizes[i] {
			diff := strSizes[i] - strSizes[i-1]
			s[i-1] = s[i-1] + strings.Repeat(" ", diff)
			strSizes[i-1] += diff
		}
	}

	size := strSizes[getBiggest(strSizes)]
	result := make([][]rune, 0, size)
	for c := 0; c < size; c++ {
		result = append(result, []rune{})
		for l := range s {
			line := s[l]
			if c >= strSizes[l] {
				continue
			}

			result[c] = append(result[c], rune(line[c]))
		}
	}
	return result
}

func (t BendSliceString) Juxtapose() []string {
	result := make([]string, 0, len(t))
	for i := range t {
		result = append(result, string([]rune(t[i])))
	}

	return result
}

func Transpose(input []string) []string {
	t := New(input)
	return t.Juxtapose()
}

func getAllSizes(s []string) []int {
	result := make([]int, 0, len(s))
	for i := range s {
		result = append(result, len(s[i]))
	}

	return result
}

func getBiggest(sizes []int) int {
	if len(sizes) == 0 {
		return -1
	}

	result := 0
	for i := range sizes {
		if sizes[i] > sizes[result] {
			result = i
		}
	}

	return result
}
