package grep

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type GFlag int

const (
	GNumber GFlag = 1
	GName GFlag = 2 
	GInsensitive GFlag = 4
	GInverted GFlag = 8
	GLine GFlag = 16
)

type FileInfo struct {
	lineNumber int
	fileName string
	lineValue string
}

type FileResult struct {
	order int
	fileName string
	result []string
}

func Search(pattern string, flags, files []string) []string {
	searchFlags, displayFlags:=textToFlags(flags)
	result:=make([]string, 0)
	size:=len(files)


	if size > 1 {
		var wg sync.WaitGroup
		dataChan:=make(chan FileResult)
		sorteredResult:=make(chan []FileResult)
		go packtAll(&wg, dataChan, sorteredResult)

		wg.Add(size)
		for i:=range files {
			file:=files[i]
			fr:=FileResult{order: i, fileName: file}
			go getFileResult(&wg, fr, pattern, searchFlags, displayFlags, dataChan)
		}
		wg.Wait()
		close(dataChan)

		subResults:=<-sorteredResult
		for i:=range subResults {
			file:=subResults[i].fileName
			data:=subResults[i].result
			result = append(result, makeFileResult(file, data, func() bool { _, ok:=displayFlags[GName]; return size > 1 && !ok })...)
		}
	} else {
		fr:=FileResult{order: 0, fileName: files[0]}
		data:=readFile(pattern, fr.fileName, searchFlags, displayFlags)
		result = append(result, makeFileResult(fr.fileName, data, func() bool {_, ok:=displayFlags[GName]; return size > 1 && !ok})...)
	}

   	return result
}

func getFileResult(wg *sync.WaitGroup, fr FileResult, pattern string, searchFlags, displayFlags map[GFlag]int, dataChan chan<- FileResult) {
	defer wg.Done()
	data:=readFile(pattern, fr.fileName, searchFlags, displayFlags)
	dataChan<-FileResult{order: fr.order, fileName: fr.fileName, result: data}
}

func packtAll(wg *sync.WaitGroup, dataChan <-chan FileResult, sorteredResult chan<- []FileResult) {
	subResults:=make([]FileResult, 0)
	for item:=range dataChan {
		subResults = append(subResults, item)
	}	

	sort.Slice(subResults, func(i, j int) bool {
		return subResults[i].order < subResults[j].order
	})
	
	sorteredResult<-subResults
}

func makeFileResult(fileName string, data []string, proc func() bool) []string {
	result:=make([]string, 0)
	if proc != nil && proc() {
		for i:=range data {
			result = append(result, fmt.Sprintf("%s:%s", fileName, data[i]))
		}
	} else {
		result = append(result, data...)
	}	
	
	return result
}

func readFile(pattern, fileName string, searchFlags, displayFlags map[GFlag]int) []string {
	result:=[]string{}
	if fileHDL, err:=os.OpenFile(fileName, os.O_RDONLY, 0555); err == nil {
		scanner:=bufio.NewScanner(fileHDL)
		rNumber:=1
		for scanner.Scan() {
			row:=scanner.Text()
			if ok:=miniGrep(pattern, row, searchFlags); ok {
				info:=FileInfo{lineNumber: rNumber, fileName: fileName, lineValue: row}
				result = append(result, resultDisplay(info, displayFlags))
				if _, ok:=displayFlags[GName]; ok {
					break
				}
			}
			rNumber++
		}
		fileHDL.Close()
	}

	return result
} 

func miniGrep(pattern, line string, flags map[GFlag]int) bool{
	_,insensitive:=flags[GInsensitive]
	if insensitive {
		pattern=strings.ToLower(pattern)
		line=strings.ToLower(line)
	}

	_, allLine:=flags[GLine]
	_, inverted:=flags[GInverted]
	if allLine && !inverted {
		return pattern == line
	} else {
		if inverted {
			return !strings.Contains(line, pattern)
		} else {
			return strings.Contains(line, pattern)
		}
	}
}

func resultDisplay(info FileInfo, displayFlags map[GFlag]int) string {
	var result strings.Builder
	_, displayNumber:=displayFlags[GNumber]
	_, displayFileName:=displayFlags[GName]

	if displayNumber && !displayFileName {
		result.WriteString(fmt.Sprintf("%d:", info.lineNumber))
	}

	if displayFileName {
		result.WriteString(info.fileName)
	} else {
		result.WriteString(info.lineValue)
	}

	return result.String()
}

func textToFlags(flags []string) (map[GFlag]int, map[GFlag]int) {
	resultA:=make(map[GFlag]int)
	resultB:=make(map[GFlag]int)
	for i:=range flags {
		flag:=flags[i]
		switch flag {
			case "-n":resultB[GNumber]=1
			case "-l":resultB[GName]=1
			case "-i":resultA[GInsensitive]=1
			case "-v":
				resultA[GInverted]=1
			case "-x":resultA[GLine]=1
		}
	}

	return resultA, resultB
}