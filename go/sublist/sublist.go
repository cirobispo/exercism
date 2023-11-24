package sublist

import "fmt"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if IsEqualSize(l1, l2) {
		if IsAnyEmpty(l1, l2) {
			return RelationEqual
		}
		if !IsSubList(l1, l2) {
			return RelationUnequal
		}
		return RelationEqual
	} else {
		if IsSubList(l1, l2) {
			if len(l2) > len(l1) {
				return RelationSublist
			} else {
				return RelationSuperlist
			}
		}
		return RelationUnequal
	}
	return RelationUnequal
}

func IsEqualSize(list1, list2 []int) bool {
	 size1, size2:=len(list1), len(list2); 
	 return size1 == size2
}

func IsAnyEmpty(list1, list2 []int) bool {
	size1, size2:=len(list1), len(list2); 
	return size1 == 0 || size2 == 0
}

func IsSubList(list1, list2 []int) bool {
	smallList, bigList,sSize, bSize := &list1, &list2, len(list1), len(list2)
	if size1, size2:=len(list1), len(list2); size1>size2  {
		smallList, bigList,sSize, bSize = &list2, &list1, len(list2), len(list1)
	}

	fmt.Printf("list1=%v, list2=%v\n", list1, list2)

	result:=false
	for i:=0; i <=bSize - sSize;i++ {
		result=true
		for j:=i; j< sSize; j++ {
			sl, bl := *smallList, *bigList
			fmt.Printf("%d=%d,", sl[i], sl[j])
			if sl[j] != bl[j] {
				result=false
				break
			}
			fmt.Println()
		}
	}   

	return result
}