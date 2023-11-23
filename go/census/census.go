// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

func isResidentInfoOK(r *Resident, callback func()) bool {
	if hasAgeAndSomeAddress := r.Name != "" && len(r.Address) > 0; hasAgeAndSomeAddress {
		value, found := r.Address["street"]
		if found && len(value) != 0 {
			if callback != nil {
				callback()
			}
			return true
		}
	}
	return false
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return isResidentInfoOK(r, nil)
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var amount int = 0
	for _, r := range residents {
		isResidentInfoOK(r, func() { amount++ })
	}
	return amount
}
