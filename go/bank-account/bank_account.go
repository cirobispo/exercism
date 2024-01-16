package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	mt sync.RWMutex
	value int64
	closed bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{ value: amount, closed: false }
}

func (a *Account) Balance() (int64, bool) {
	var result int64
	ok:=false

	a.mt.RLock()
	if !a.closed {
		result, ok= a.value, true		
	}
	a.mt.RUnlock()

	return result, ok
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	var result int64
	ok:=false

	a.mt.Lock()
	if !a.closed {
		if a.value+amount >= 0 {
			a.value+=amount

			result, ok= a.value, true		
		}
	}
	a.mt.Unlock()

	return result, ok
}

func (a *Account) Close() (int64, bool) {
	var result int64
	ok:=false

	a.mt.Lock()
	if !a.closed {
		a.closed=true
		result, ok=a.value, true		
	}
	a.mt.Unlock()

	return result, ok
}
