package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const maxCnt = 10

func Top10(s string) []string {
	givenWords := strings.Fields(s)
	wordsFreqMap := map[string]int{}
	for _, word := range givenWords {
		wordsFreqMap[word]++
	}

	wordsList := make([]string, 0, len(wordsFreqMap))
	for word := range wordsFreqMap {
		wordsList = append(wordsList, word)
	}

	sort.Slice(wordsList, func(i, j int) bool {
		if wordsFreqMap[wordsList[i]] != wordsFreqMap[wordsList[j]] {
			return wordsFreqMap[wordsList[i]] > wordsFreqMap[wordsList[j]]
		}

		return strings.Compare(wordsList[i], wordsList[j]) == -1
	})

	if len(wordsList) >= maxCnt {
		return wordsList[:maxCnt]
	}

	return wordsList
}
