package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
	for i := 0; i < s.Length(); i++ {
		result = fn(result, s[i])
	}

	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
	for i := s.Length() - 1; i > -1; i-- {
		result = fn(s[i], result)
	}

	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0)
	for i := range s {
		if fn(s[i]) {
			result = append(result, s[i])
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, 0)
	for i := range s {
		result = append(result, fn(s[i]))
	}

	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, 0, s.Length())
	for i := s.Length() - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, 0, s.Length()+lst.Length())
	result = append(result, s...)
	result = append(result, lst...)
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	size := s.Length()
	for i := range lists {
		size += lists[i].Length()
	}

	result := make(IntList, 0, size)
	result = append(result, s...)
	for i := range lists {
		result = append(result, lists[i]...)
	}

	return result
}
