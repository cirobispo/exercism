package prime

type PrimeNumbers []int64

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

func(pn PrimeNumbers) SearchNewOne() (int64, bool) {
	size:=len(pn)
	lastOne:=pn[size-1]
	divisors:=[]int64{1}
	divisors=append(divisors, pn...)
	for i:=0; i<size; i++ {
		nums2Sum:=1
		for nums2Sum+i < size {
			var soma int64=0
			for j:=i;j<i+nums2Sum; j++ {
				soma+=divisors[j]
			}
			nums2Sum++
			if value, is:=pn.IsPrime(lastOne + soma); is {
				return value, true
			}
		}
	}
	return 0, false
}

func subDivisors(n int64, pn PrimeNumbers) []int64 {
	result:=[]int64{}
again:
	if value, isPrime:=pn.IsPrime(n); !isPrime {
		if value > 1 {
			result = append(result, value)
			result = append(result, subDivisors(n/value, pn)...)
		}
	} else {
		if value, is:=pn.SearchNewOne(); is {
			pn.Add(value)
			goto again
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
