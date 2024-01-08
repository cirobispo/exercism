package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	dataChan:=make(chan FreqMap, len(texts))
	
	go goRoutine(dataChan, texts)

	var allFreqMap=make(FreqMap)
	for data:= range dataChan {
		for k, v:=range data {
			allFreqMap[k]+=v
		}
	}

	return allFreqMap
}

func goRoutine(data chan<- FreqMap, texts []string) {
	var wg sync.WaitGroup
	for i:=range texts {
		text:=texts[i]
		wg.Add(1)
		go func (wg *sync.WaitGroup, text string) {
			defer wg.Done()
			data<-Frequency(text)
		}(&wg, text)
	}

	wg.Wait()
	close(data)
}

