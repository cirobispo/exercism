package pythagorean

import (
	"fmt"
	"math"
	"runtime"
	"slices"
	"sort"
	"sync"
)

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	result:=make([]Triplet, 0)

	for i:=min; i<=max; i++ {
		if ((i / 3)*3 >= min) && ((i / 3)*5 <= max) { 
			k:=i/3
			result=append(result, makeTriplet(12*k))
			i+=2
		}
	}

	return result
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	result:=make([]Triplet, 0)

	// r3,r4,r5:=p%3,p%4,p%5
	type info struct {
		worker int
		triplet Triplet
	}

	coreCount:=runtime.NumCPU()
	var wg sync.WaitGroup
	channel:=make(chan info)

	worker:=func(start, all int, wg *sync.WaitGroup, channel chan<- info) {
		workerNum:=(start/(all/coreCount))+1

		triplet:=[3]int{start, (all-start)/2, (all-start) - ((all-start)/2)}
		for isATriangle(triplet) {
			for isATriangle(triplet) {
				if isPithagorean(triplet) {
					sortTriplet(&triplet)		
					channel<-info{workerNum, triplet}
				}	
				triplet[1]-=1
				triplet[2]+=1
			}
			start++
			triplet=[3]int{start, (all-start)/2, (all-start) - ((all-start)/2)}
			sortTriplet(&triplet)
		}
		wg.Done()
	}

	processResponse:=func(wg *sync.WaitGroup, channel <-chan info) {
		for item:=range channel {
			fmt.Printf("worker #%d gets triplet: %v \n", item.worker, item.triplet)
			if !slices.Contains(result, item.triplet) {
				result = append(result, item.triplet)
			}
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0] && result[i][1] < result[j][1] && result[i][2] < result[j][2]  
		})
		
		wg.Done()
	}

	fmt.Println(math.Round(float64(p-(p/coreCount))))

	go processResponse(&wg, channel)

	wg.Add(coreCount)
	for i,b:=0,p/coreCount; i/b < p/b; i+=b {
		go worker(i+3, p, &wg, channel)
	}
	wg.Wait()
	wg.Add(1)
	close(channel)
	wg.Wait()

	return result
}

func makeTriplet(values ...int) Triplet {
	if len(values) == 0 {
		panic(fmt.Errorf("no argument was given. It needs to be one or three"))
	}

	if len(values) == 3 {
		return Triplet{values[0], values[1], values[2]}
	}

	k:=values[0]/12
	a,b,c:=3*k,4*k,5*k
	return Triplet{a, b, c}
}

/**/
func isATriangle(t [3]int) bool {
	isOK := func(a, b, c int) bool {
		return a < b+c
	}

	return isOK(t[0], t[1], t[2]) && isOK(t[1], t[0], t[2]) && isOK(t[2], t[0], t[1])
}

func isPithagorean(t [3]int) bool {
	return float64(t[2]) == math.Sqrt(math.Pow(float64(t[0]), 2) + math.Pow(float64(t[1]), 2))
}

func sortTriplet(t *[3]int) {
	for i := 0; i < 3; i++ {
		for j := 2; j > i; j-- {
			if t[i] > t[j] {
				t[j], t[i] = t[i], t[j]
			}
		}
	}
}
/**/