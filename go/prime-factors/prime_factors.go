package prime

import (
	"sort"
	"sync"
)

type PrimeNumbers []int64
type PAV struct {
	position, acumulator, foundValue int64
}

func (pn PrimeNumbers) IsPrime(value int64) (int64, bool) {
	result, OK:=value, true
	for i:= range pn {
		if value % pn[i] == 0 {
			result, OK=pn[i], false
			break
		} 
	}

	return result, OK
}

func (pn *PrimeNumbers) Add(value int64) {
	*pn=append(*pn, value)
}

func(pn PrimeNumbers) ResearchNewPrime(pavs []PAV) ([]int64, bool) {
	search:=func(wg *sync.WaitGroup, pav *PAV, found chan bool) {
		defer wg.Done()

		for i:=pav.position; i< pav.position+500; i+=pav.acumulator {
			select {
			case <-found:
				goto quit
			default:
				if value, f:=pn.IsPrime(i); f {
					pav.foundValue=value
					goto quit
				}
			}
		}
quit:
		found<- true
		//Avise ao proximo!
		<-found
	}

	result:=make(map[int64]int)
	size:=len(pavs)
	found:=make(chan bool, size)
	var wg sync.WaitGroup
	for i:=0; i< size; i++ {
		pav:=&pavs[i]
		wg.Add(1)
		go search(&wg, pav, found)		
	}
	wg.Wait()

	for i:=range pavs {
		if pavs[i].foundValue > 0 {
			result[pavs[i].foundValue]++
		}
	}

	keys := make([]int64, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}
	
	return keys, len(keys) > 0
}

func subDivisors(n int64, pn PrimeNumbers) []int64 {
	result:=[]int64{}

	getPAVs:=func(n int64, pavCount int, pavs []PAV) []PAV {
		size:=len(pavs)
		if size == 0 { 
			biggestPrime:=pn[len(pn)-1]
			pav:=PAV{position: biggestPrime}
			pavs=[]PAV{pav, pav, pav}
			size=len(pavs)
			for i:=0 ; i<size; i++ {
				pavs[i].acumulator=pn[i]
			}
		}

		result:=make([]PAV, 0, size)
		for i:=0; i<size; i++ {
			newPosition:=pavs[i].position
			if pavs[i].foundValue > 0 {
				newPosition=pavs[i].foundValue
				if newPosition > n {
					newPosition=pn[i + size]
					pavs[i].acumulator=pn[i + size]
				}
			}
			pav:=PAV{position: newPosition, acumulator: pavs[i].acumulator, foundValue: 0}
			result = append(result, pav)
		}

		return result
	}


	i:=n

	pavs:=getPAVs(i, 6, []PAV{})

	for i > 1 {
		if value, isPrime:=pn.IsPrime(i); !isPrime {
			result = append(result, value)
			i=i/value
		} else {
			for {
				if value, is:=pn.ResearchNewPrime(pavs); is {
					for j:=range value {
						pn.Add(value[j])
					}
					sort.Slice(pn, func(i, j int) bool {
						return pn[i] < pn[j]
					})

					pavs=getPAVs(i, 6, pavs)
					break
				}
			}
		}
	}

	return result
}

func Factors(n int64) []int64 {
	var pn PrimeNumbers = PrimeNumbers{}
	for _, value:=range []int64{2, 3, 5, 7, 11, 13} {
		pn.Add(value)
	}

	result:=subDivisors(n, pn)

	return result
}
