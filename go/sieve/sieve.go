package sieve

type prime struct {
	number    int
	nextPrime *prime
}

func NewPrime(number int) *prime {
	return &prime{number: number, nextPrime: nil}
}

func (p *prime) CheckAndUpdateList(value int) {
	if value%p.number != 0 {
		if p.nextPrime != nil {
			p.nextPrime.CheckAndUpdateList(value)
		} else {
			p.nextPrime = NewPrime(value)
		}
	}
}

func (p *prime) GetAllPrimes() []int {
	result := make([]int, 0, 5)
	result = append(result, p.number)
	if p.nextPrime != nil {
		result = append(result, p.nextPrime.GetAllPrimes()...)
	}

	return result
}

func Sieve(limit int) []int {
	result := []int{}
	if limit > 1 {
		p := NewPrime(2)
		for i := 3; i <= limit; i++ {
			p.CheckAndUpdateList(i)
		}

		result = p.GetAllPrimes()
	}

	return result
}
