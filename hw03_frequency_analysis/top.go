package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordData struct {
	Word string
	Freq int
}

func Top10(text string) []string {
	wordFreq := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		wordFreq[word]++
	}

	wordFreqSort := make([]WordData, 0)

	for word, freq := range wordFreq {
		wordFreqSort = append(wordFreqSort, WordData{word, freq})
	}

	sort.Slice(wordFreqSort, func(i, j int) bool {
		if wordFreqSort[i].Freq == wordFreqSort[j].Freq {
			return wordFreqSort[i].Word <= wordFreqSort[j].Word
		}
		return wordFreqSort[i].Freq > wordFreqSort[j].Freq
	})
	result := make([]string, 0, 10)
	for _, wordData := range wordFreqSort {
		result = append(result, wordData.Word)
	}

	if len(result) > 10 {
		result = result[0:10]
	}
	return result
}
