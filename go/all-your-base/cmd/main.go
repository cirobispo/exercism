package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	name string 
	amount int
}

func main() {
	fmt.Println(decode("coding_qual_input.txt"))
}

func parseFile(fileName string) []Item {
	result:=make([]Item, 0)
	if file, err:=os.Open(fileName); err == nil {
		scanner:=bufio.NewScanner(file)
		for scanner.Scan() {
			line:=scanner.Text()
			data:=strings.Split(line, " ")
			if len(data) == 2 {
				number, err:=strconv.ParseInt(data[0], 10, 64)
				if err == nil {
					result=append(result, Item{name: data[1], amount: int(number)})
				}
			}
		}
		file.Close()
	}
	return result
}

func sortItems(data []Item) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].amount < data[j].amount
	})
}

func extractStair(data []Item) string {
	step:=1
	position:=0
	result:=""
	for {
		if position < len(data) {
			result=result+" "+data[position].name
			step++
			position+=step
		} else {
			break
		}
	}

	if len(result) > 0 {
		result=result[1:]
	}

	return result
}

func decode(message_file string) string {
	items:=parseFile(message_file)
	sortItems(items)
	result:=extractStair(items)

	return result
}
