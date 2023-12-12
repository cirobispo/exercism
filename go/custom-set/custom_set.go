package stringset

import (
	"strings"
)

type Set []string

func New() Set {
	result:=make([]string, 0)
	return result
}

func NewFromSlice(l []string) Set {
	result:=New()

	if len(l) > 0 {
		for i:=range l {
			result.Add(l[i])
		}
	}

	return result
}

func (s Set) String() string {
	var b strings.Builder
	for i:=range s {
		b.Write([]byte(", \""))
		b.Write([]byte(s[i]))
		b.Write([]byte("\""))
	}

	result:=b.String()
	if len(result) > 0 {
		return "{" + result[2:] + "}"
	}
	return "{}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	result:=false;
	for i:= range s {
		if s[i] == elem {
			result=true
			break
		}
	}

	return result
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		*s = append(*s, elem)
	}
}

func Subset(s1, s2 Set) bool {
	sSet:=s1
	bSet:=s2
	if len(s1) > len(s2) {
		sSet=s2
		bSet=s1
	}

	result:=false
	if len(s1) == 0 && len(s2) >= 0 {
		result=true
	}

	if len(s1) > 0 && len(s2) > 0 {
		for i, j, ssize:=0,0,len(sSet); i<len(bSet); i++ {
			if sSet[j] == bSet[i] {
				j++
				if j == ssize {
					result = true
					break
				}
			} else {
				j=0
			}
		}
	}

	return result
}

func Disjoint(s1, s2 Set) bool {
	result:=Intersection(s1, s2)
	return result.IsEmpty()
}

func Equal(s1, s2 Set) bool {
	result :=false
	if len(s1) == len(s2) {
		eCount:=0
		for idx:= range s2 {
			if s1.Has(s2[idx]) {
				eCount++
				continue
			}
			break;
		}
		result=(eCount == len(s1))
	}

	return result
}

func Intersection(s1, s2 Set) Set {
	result:=New()

	sSet:=s1
	bSet:=s2
	if len(s1) > len(s2) {
		sSet=s2
		bSet=s1
	}

	for idx:=range sSet {
		if bSet.Has(sSet[idx]) {
			result.Add(sSet[idx])
		}
	}

	return result
}

func Difference(s1, s2 Set) Set {
	notIn:=Intersection(s1, s2)

	getNotIn:=func(data []string) []string {
		result:=make([]string, 0)
		for i:= range data {
			if !notIn.Has(data[i]) {
				result =append(result, data[i])
			}
		}
		return result
	}	

	allData:=make([]string, 0)
	if len(s1) > 0 {
		allData = append(allData, getNotIn(s1)...)
	}
	return NewFromSlice(allData)
}

func Union(s1, s2 Set) Set {
	in:=Intersection(s1, s2)
	diff:=Difference(s1, s2)
	diff2:=Difference(s2, s1)

	result:=New()
	for i:=range diff {
		result = append(result, diff[i])
	}
	
	for i:=range in {
		result = append(result, in[i])
	}
	
	for i:=range diff2 {
		result = append(result, diff2[i])
	}
	
	return result
}
