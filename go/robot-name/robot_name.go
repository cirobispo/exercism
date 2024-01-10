package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

const maxRobotNames = 26 * 26 * 10 * 10 * 10

var (
	namePool = generateRobotNames()
	idx      = 0
)

func generateRobotNames() []string {
	pos := 0
	names := make([]string, maxRobotNames)

	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			for k := 0; k < 1000; k++ {
				names[pos] = fmt.Sprintf("%s%s%03d", string(i), string(j), k)
				pos++
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}

// Name returns the existing name or returns a newly assigned name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if idx >= maxRobotNames {
		return "", fmt.Errorf("uniqueness exhausted")
	}

	r.name = namePool[idx]
	idx++

	return r.name, nil
}

// Reset erases the existing name without returning it to unused names.
func (r *Robot) Reset() {
	r.name = ""
}