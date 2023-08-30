package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var delimiter = regexp.MustCompile(`[ \n\r\t]`)

const MostFrequentWordsAmount = 10

type WordCounter struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	wordsCountMap := map[string]int{}

	for _, word := range delimiter.Split(text, -1) {
		if len(word) == 0 || delimiter.MatchString(word) {
			continue
		}
		wordsCountMap[word]++
	}

	wordsSlice := make([]string, 0, len(wordsCountMap))
	for word := range wordsCountMap {
		wordsSlice = append(wordsSlice, word)
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		a, b := wordsSlice[i], wordsSlice[j]
		if wordsCountMap[a] == wordsCountMap[b] {
			return a < b
		} else {
			return wordsCountMap[a] > wordsCountMap[b]
		}
	})

	if len(wordsCountMap) > MostFrequentWordsAmount {
		return wordsSlice[:MostFrequentWordsAmount]
	}
	return wordsSlice[:len(wordsCountMap)]
}
