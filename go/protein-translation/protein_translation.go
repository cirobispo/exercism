package protein

import (
	"errors"
)

var ErrStop = errors.New("app: stop error")
var ErrInvalidBase = errors.New("app: invalid base error")

// :---                  | :---
// AUG                   | Methionine
// UUU, UUC              | Phenylalanine
// UUA, UUG              | Leucine
// UCU, UCC, UCA, UCG    | Serine
// UAU, UAC              | Tyrosine
// UGU, UGC              | Cysteine
// UGG                   | Tryptophan
// UAA, UAG, UGA         | STOP

type MultiMap struct {
	raw_keys   []*string
	raw_values []*string
	v2k        map[*string][]*string
	k2v        map[*string]*string
}

func NewMultiMap() *MultiMap {
	rk, rv := make([]*string, 0), make([]*string, 0)
	k2, v2 := make(map[*string][]*string), make(map[*string]*string)

	return &MultiMap{rk, rv, k2, v2}
}

func (mm *MultiMap) removeValue(value *string) bool {
	delete(mm.v2k, value)
	_, result := removeItemFrom(mm.raw_values, *value)
	return result
}

func (mm *MultiMap) updateKey(key string, new_value *string) bool {
	result := false
	if addr_k, found := isItIn(mm.raw_keys, key); found {
		//procuro no map
		if valueFK, fk := mm.k2v[addr_k]; fk {
			if keyFV, fv := mm.v2k[valueFK]; fv {
				mm.v2k[valueFK], result = removeItemFrom(keyFV, *valueFK)
				if len(mm.v2k[valueFK]) == 0 {
					mm.removeValue(valueFK)
				}
			}
		}
	}

	return result
}

func (mm *MultiMap) Add(keys []string, value string) {
	var addr_v *string = &value
	if af, fv := isItIn(mm.raw_values, value); fv {
		addr_v = af
	} else {
		mm.raw_values = append(mm.raw_values, addr_v)
		mm.v2k[addr_v] = []*string{}
	}

	for _, item := range keys {
		if ak, fk := isItIn(mm.raw_keys, item); fk {
			mm.updateKey(*ak, addr_v)
		} else {
			key := item
			mm.raw_keys = append(mm.raw_keys, &key)
			mm.k2v[&key] = addr_v
			mm.v2k[addr_v] = append(mm.v2k[addr_v], &key)
		}
	}

}

func (mm *MultiMap) GetValue(key string) (string, bool) {
	if addr, fk := isItIn(mm.raw_keys, key); fk {
		data := mm.k2v[addr]
		return *data, true
	}
	return "", false
}

func (mm *MultiMap) UpdateKeys(keys []string, new_value string) []bool {
	addr_value := &new_value
	nv_exists := false
	if addr, fv := isItIn(mm.raw_values, new_value); !fv {
		mm.raw_values = append(mm.raw_values, addr_value)
		mm.v2k[addr_value] = []*string{}
	} else {
		addr_value = addr
		nv_exists = true
	}

	var result []bool = make([]bool, len(keys))
	hasChanged := false
	for index, item := range keys {
		if mm.updateKey(item, addr_value) {
			result[index] = true
			hasChanged = true
			continue
		}
		result[index] = false
	}

	if !hasChanged && nv_exists {
		mm.removeValue(addr_value)
	}

	return result
}

func (mm *MultiMap) GetKeys(value string) ([]string, bool) {
	result := []string{}

	if addr, fv := isItIn(mm.raw_values, value); fv {
		keys := mm.v2k[addr]
		result = make([]string, len(keys))
		for index, k := range keys {
			result[index] = *k
		}

		return result, true
	}

	return result, false
}

func isItIn(set []*string, search string) (*string, bool) {
	for _, item := range set {
		if *item == search {
			return item, true
		}
	}

	return nil, false
}

func removeItemFrom(set []*string, search string) ([]*string, bool) {
	var deleted bool = false
	var result []*string = make([]*string, len(set))
	count := 0
	for index := 0; index < len(set); index++ {
		if *set[index] != search {
			result[count] = set[index]
			count++
		} else {
			deleted = true
		}
	}

	if deleted {
		result = result[:len(result)-1]
	}

	return result, deleted
}

func buildData() *MultiMap {
	mm := NewMultiMap()
	mm.Add([]string{":---"}, ":---")
	mm.Add([]string{"AUG"}, "Methionine")
	mm.Add([]string{"UUU", "UUC"}, "Phenylalanine")
	mm.Add([]string{"UUA", "UUG"}, "Leucine")
	mm.Add([]string{"UCU", "UCC", "UCA", "UCG"}, "Serine")
	mm.Add([]string{"UAU", "UAC"}, "Tyrosine")
	mm.Add([]string{"UGU", "UGC"}, "Cysteine")
	mm.Add([]string{"UGG"}, "Tryptophan")
	mm.Add([]string{"UAA", "UAG", "UGA"}, "STOP")
	return mm
}

func FromRNA(rna string) ([]string, error) {
	if rest := len(rna) % 3; rest == 0 {
		mm := buildData()
		var result []string = make([]string, 0)

		for i := 0; i < len(rna); i += 3 {
			if value, found := mm.GetValue(rna[i : i+3]); found {
				if value == "STOP" {
					break
				}
				result = append(result, value)
			} else {

			}
		}

		return result, nil
	}

	return []string{}, ErrInvalidBase
}

func FromCodon(codon string) (string, error) {
	mm := buildData()
	if value, found := mm.GetValue(codon); found {
		if value == "STOP" {
			return "", ErrStop
		}
		return value, nil
	} else {
		return "", ErrInvalidBase
	}
}
