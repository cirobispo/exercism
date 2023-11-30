package summultiples

import "fmt"

type multiples []int

func SumMultiples(limit int, divisors ...int) int {
	fmt.Println("Limit: ", limit, "Divisors: ", divisors)
	searchMultiples:=func(limit, multiple int, result chan<- multiples) {
		found:=make(multiples, 0)
		if multiple != 0 {
			for i:=1; i < limit; i++ {
				if i % multiple == 0 {
					found = append(found, i)
				}
			}
		}

		result <- found
	}

	sumMultiples:=func(chanCount int, result <-chan multiples, sum chan int) {
		all:=make(map[int]int)
		for lido:=0; lido < chanCount; lido++ {
			divisors:=<-result 
			fmt.Println(divisors)
			for i:=range divisors {
				all[divisors[i]]++
			}
		}

		count:=0
		for key:= range all {
			count+=key
		}

		sum<- count
	}

	if countDivisors:=len(divisors); countDivisors > 0 {
		sum:=make(chan int)
		results:=make(chan multiples, countDivisors)
		go sumMultiples(countDivisors, results, sum)
		for i:=0; i < countDivisors; i++ {
			go searchMultiples(limit, divisors[i], results)
		}

		return <-sum
	}

	return 0
}
