package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	sizeA, sizeB:=len(l1),len(l2)

	if sizeA == sizeB  {
		if sizeA == 0 && sizeB == 0 {
			return RelationEqual
		} else {
			result:=RelationEqual
			for i:=0; i<sizeA; i++ {
				if l1[i] != l2[i] {
					result = RelationUnequal
					break
				}
			}
			return result
		}
	} else {
		if (sizeA > 0 && sizeB == 0) || (sizeB > 0 && sizeA == 0) {
			if sizeA > 0 {
				return RelationSuperlist
			} else {
				return RelationSublist
			}
		} else {
			return RelationUnequal
		}
	}
}

func IsSubList(smallList, bigList []int) bool {
	bSize, sSize:=len(bigList), len(smallList)

	result:=false
	for i:=0; i <=bSize - sSize;i++ {
		result=true
		for j:=i; j<= sSize; j++ {
			if smallList[j] != bigList[j] {
				result=false
				break
			}
		}
	}   

	return result
}