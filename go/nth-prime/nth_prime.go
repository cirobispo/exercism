package prime

import "errors"

type prime struct {
	divider int64
	node    *prime
}

func (p *prime) IsItPrime(value int64) bool {
	if yep := value%p.divider == 0; !yep {
		if p.node != nil {
			return p.node.IsItPrime(value)
		}

		return false
	}

	return true
}

func (p *prime) AddSibling(next *prime) {
	if p.node != nil {
		p.node.AddSibling(next)
	} else {
		p.node = next
	}
}

func (p *prime) Size() int {
	if p.node != nil {
		return 1 + p.node.Size()
	}

	return 1
}

func (p *prime) Last() int64 {
	if p.node != nil {
		return p.node.Last()
	}

	return p.divider
}

func (p *prime) Nth(pos int) int {
	if pos > 0 {
		if p.node != nil {
			p.node.Nth(pos - 1)
		}
	}
}

func NewPrime(value int64) *prime {
	return &prime{value, nil}
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	prime2 := NewPrime(2)

	if n > 0 {
		var number int64 = 2
		for {
			if !prime2.IsItPrime(number) {
				prime2.AddSibling(NewPrime(number))
			}

			if prime2.Size() == n {
				break
			}

			number++
		}

		return int(prime2.Last()), nil
	}

	return 0, errors.New("it cant compute negative number")
}
