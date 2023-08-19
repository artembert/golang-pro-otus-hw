package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var delimiter = regexp.MustCompile(`[ \n\r\t]`)

const MostFrequentWordsAmout = 10

type WordsByCont struct {
	Count     int
	WordsList []string
}

type WordCounter struct {
	Word  string
	Count int
}

func getSliceFromDictionary(dictionary map[string]int) []WordCounter {
	arr := []WordCounter{}

	for key, value := range dictionary {
		arr = append(arr, WordCounter{Word: key, Count: value})
	}
	return arr
}

func groupWordsByFrequency(arr []WordCounter) map[int][]string {
	frequencyDictionary := map[int][]string{}

	for _, value := range arr {
		frequencyDictionary[value.Count] = append(frequencyDictionary[value.Count], value.Word)
	}

	return frequencyDictionary
}

func getMostFrequentWords(dictionary map[int][]string, amount int) []string {
	sortedFrequencies := make([]int, 0, len(dictionary))
	for frequency := range dictionary {
		sortedFrequencies = append(sortedFrequencies, frequency)
	}
	sort.Slice(sortedFrequencies, func(i, j int) bool {
		return sortedFrequencies[i] > sortedFrequencies[j]
	})

	count := amount
	currentFrequentIndex := 0
	result := make([]string, 0, amount)
	for count > 0 && currentFrequentIndex < len(sortedFrequencies) {
		wordsListByCurrentFrequency := dictionary[sortedFrequencies[currentFrequentIndex]]
		sort.Strings(wordsListByCurrentFrequency)
		index := 0
		for count > 0 && index < len(wordsListByCurrentFrequency) {
			result = append(result, wordsListByCurrentFrequency[index])
			count--
			index++
		}
		currentFrequentIndex++
	}

	return result
}

func Top10(text string) []string {
	dictionary := map[string]int{}

	for _, word := range delimiter.Split(text, -1) {
		if len(word) == 0 || delimiter.MatchString(word) {
			continue
		}
		if number, ok := dictionary[word]; ok {
			dictionary[word] = number + 1
		} else {
			dictionary[word] = 1
		}
	}

	return getMostFrequentWords(groupWordsByFrequency(getSliceFromDictionary(dictionary)), MostFrequentWordsAmout)
}
