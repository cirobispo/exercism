package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile("^\\[(TRC|DBG|INF|WRN|ERR|FTL)\\]")
	return r.Find([]byte(text)) != nil
}

func SplitLogLine(text string) []string {
	//r := regexp.MustCompile("(\\w* \\d)|(\\<\\w+\\>)")
	r := regexp.MustCompile("(<\\W{1,6}>)|(<>)")
	result := r.Split(text, -1)
	return result
}

func CountQuotedPasswords(lines []string) int {
	r := regexp.MustCompile(" [p|P][a|P][s|S][s|S][w|W][o|O][r|R][d|D]")
	var result int = 0
	for _, item := range lines {
		result += len(r.FindAllString(item, -1))
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile("end-of-line\\d*")
	result := r.ReplaceAllString(text, "")
	return result
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile("(User\\s+\\w*)")
	ru := regexp.MustCompile("(User\\s*)")
	result := make([]string, 0, len(lines))
	for _, item := range lines {
		if r.MatchString(item) {
			result = append(result, string(ru.ReplaceAll([]byte(r.FindString(item)), []byte("[USR] ")))+" "+item)
		} else {
			result = append(result, item)
		}
	}
	return result
}
