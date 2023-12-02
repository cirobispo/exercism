package kindergarten

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

type Child struct {
	name   string
	plants []string
}

// Define the Garden type here.
type Garden struct {
	children []Child
}

func isPlantOK(plantRows []string) bool {
	result := true
	pc := len(plantRows[0]) / 2

	for i := range plantRows {
		if c := len(plantRows[i]); c/2 != pc || c%2 != 0 || unicode.IsLower(rune(plantRows[i][0])) {
			result = false
			break
		}
	}

	return result
}

func isChildrenOK(children []string) bool {
	allChild := make(map[string]int)

	result := true
	for i := range children {
		allChild[children[i]]++
		if value, found := allChild[children[i]]; found && value > 1 {
			result = false
			break
		}
	}

	return result
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	plantRows := strings.Split(diagram, "\n")
	wrongFormat := plantRows[0] != ""
	plantRows = plantRows[1:]

	if wrongFormat || !isPlantOK(plantRows) || !isChildrenOK(children) {
		return nil, errors.New("plant count or child is wrong")
	}

	splitPlants := func(plantRows []string, children []string) []Child {
		result := make([]Child, 0, len(children))
		for i := range children {
			result = append(result, Child{name: children[i]})
		}

		sort.Slice(result, func(i int, j int) bool {
			return result[i].name < result[j].name
		})

		if len(plantRows) > 0 {
			plantsPerRow := len(plantRows[0]) / 2
			for j := 0; j < plantsPerRow; j++ {
				plants := make([]string, 0, plantsPerRow)
				for i := range plantRows {
					plants = append(plants, plantRows[i][j*2:j*2+2])
				}
				result[j].plants = plants
			}
		}
		return result
	}

	return &Garden{children: splitPlants(plantRows, children)}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	getPlantNames := func(plants []string) []string {
		plantNames := map[rune]string{'C': "clover", 'G': "grass", 'R': "radishes", 'V': "violets"}

		result := make([]string, 0, len(plants)*2)
		for i := range plants {
			plant := plants[i]
			for j := range plant {
				result = append(result, plantNames[rune(plant[j])])
			}
		}
		return result
	}

	for i := range g.children {
		if g.children[i].name == child {
			return getPlantNames(g.children[i].plants), true
		}
	}

	return []string{}, false
}
