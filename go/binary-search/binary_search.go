package binarysearch

import (
	"math"
)

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}

	subList:=list
	startIndex:=0
	result:=-1
	for {
		size:=float64(len(subList))
		half:=int(math.Round(size/2))

		centerValue:=subList[half-1]

		if centerValue == key { 
			result=startIndex + half-1
			break
		} else {
			if size == 1 {
				break
			}
		}

		temp:=subList[:half]
		if key > centerValue {
			temp=subList[half:]
			startIndex+=half
		}
		subList=temp
	}

	return result
}