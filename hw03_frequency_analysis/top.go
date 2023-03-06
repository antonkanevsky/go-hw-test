package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordFrequency struct {
	Word string
	Freq int
}

const MaxCnt = 10

func Top10(s string) []string {
	givenWords := strings.Fields(s)
	wordsFreqMap := map[string]int{}
	for _, word := range givenWords {
		wordsFreqMap[word]++
	}

	wordFrequencyList := make([]WordFrequency, len(wordsFreqMap))

	i := 0
	for word, freq := range wordsFreqMap {
		wordFrequencyList[i].Word = word
		wordFrequencyList[i].Freq = freq
		i++
	}

	sort.Slice(wordFrequencyList, func(i, j int) bool {
		if wordFrequencyList[i].Freq != wordFrequencyList[j].Freq {
			return wordFrequencyList[i].Freq > wordFrequencyList[j].Freq
		}

		return strings.Compare(wordFrequencyList[i].Word, wordFrequencyList[j].Word) == -1
	})

	topWords := []string{}
	for i := 0; i < len(wordFrequencyList); i++ {
		if i == MaxCnt {
			break
		}
		topWords = append(topWords, wordFrequencyList[i].Word)
	}

	return topWords
}
