package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	result := make(map[string]int)
	result["quarter_of_a_dozen"] = 3
	result["half_of_a_dozen"] = 6
	result["dozen"] = 12
	result["small_gross"] = 120
	result["gross"] = 144
	result["great_gross"] = 1728
	return result
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	result := make(map[string]int)
	return result
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	billTotal, foundItem := bill[item]
	unitValue, foundUnit := units[unit]
	if foundUnit {
		if foundItem {
			unitValue += billTotal
		}
		bill[item] = unitValue
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	sum, foundItem := bill[item]
	if foundItem {
		unitValue, foundUnit := units[unit]
		if foundUnit {
			if unitValue == sum {
				delete(bill, item)
				return true
			} else if unitValue < sum {
				bill[item] = sum - unitValue
				return true
			}
			return false
		} else {
			return false
		}
	}
	return foundItem
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, found := bill[item]
	return value, found
}
