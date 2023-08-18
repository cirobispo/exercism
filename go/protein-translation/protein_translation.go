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
	v2k        map[string]*string
	k2v        map[string]*string
}

func NewMultiMap() *MultiMap {
	rk, rv := make([]*string, 0), make([]*string, 0)
	k2, v2 := make(map[string]*string), make(map[string]*string)

	return &MultiMap{rk, rv, k2, v2}
}

func (mm *MultiMap) Add(keys []string, value string) {
	if !isItIn(mm.raw_values, value) {
		exist_keys := []string{}
		for _, item := range keys {
			if !isItIn(mm.raw_keys, item) {
				mm.k2v[item] = &value
				continue
			}
			exist_keys = append(exist_keys, item)
		}

	}
	//if keys, fk := mm.values[value]
}

func isItIn(set []*string, search string) bool {
	for _, item := range set {
		if *item == search {
			return true
		}
	}

	return false
}

/**/

func FromRNA(rna string) ([]string, error) {
	if rest := len(rna) % 3; rest == 0 {
		var result []string = make([]string, len(rna)*3)

		for i := 0; i < len(rna); i += 3 {
			result[i/3] = rna[i : i+3]
		}

		return result, nil
	}

	return []string{}, errors.New("wrong number of rna sequence")
}

func FromCodon(codon string) (string, error) {
	//	fmt.Println(p)

	return "", nil
}
