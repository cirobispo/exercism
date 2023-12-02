package summultiples

type multiples []int

func SumMultiples(limit int, divisors ...int) int {
	searchMultiples:=func(limit, multiple int, result chan<- multiples) {
		found:=make(multiples, 0, 4)
		if multiple != 0 {
			for i:=1; i < limit; i++ {
				if i % multiple == 0 {
					found = append(found, i)
				}
			}
		}

		result <- found
	}

	sumMultiples:=func(chanCount int, result <-chan multiples, sum chan<- int) {
		count:=0
		all:=make(map[int]int)
		for lido:=0; lido < chanCount; lido++ {
			divisors:=<-result 
			for i:=range divisors {
				all[divisors[i]]++
				if all[divisors[i]] == 1 {
					count+=divisors[i]
				}
			}
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
