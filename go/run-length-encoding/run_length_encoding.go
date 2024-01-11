package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var result strings.Builder

	if size:=len(input); size > 0 {
		sum:=0
		lc:=input[0]
		for i:=range input {
			cc:=input[i]
			if cc == lc {
				sum++
				if i == (size-1) {
					result.WriteString(fmt.Sprintf("%d%c", sum, lc))
				}
			} else {
				if sum > 1 {
					result.WriteString(fmt.Sprintf("%d", sum))
				}
				result.WriteString(fmt.Sprintf("%c", lc))
				if i == (size-1) {
					result.WriteString(fmt.Sprintf("%c", cc))
				} else {
					lc=cc
					sum=1
				}
			}
		}
	}

	return result.String()
}

func RunLengthDecode(input string) string {
	var result strings.Builder

	if len(input) > 0 {
		reg:=regexp.MustCompile(`[0-9]*[\w ]`)
		data:=reg.FindAllString(input, -1)
		for i:=range data {
			item:=data[i]
			c:=item
			if size:=len(item); size > 1 {
				c=item[len(item)-1:]
				count, _:=strconv.ParseInt(item[0:size-1], 10, 64)
				result.WriteString(strings.Repeat(c, int(count)-1))
			}
			result.WriteString(c)
		}
	}

	return result.String()
}