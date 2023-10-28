package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type kv struct {
	key   string
	value int
}

func Top10(input string) []string {
	result := make([]string, 0, 10)
	// split into words by whitespaces
	tmpSlice := strings.Fields(input)
	// make map to store word occurrence count, e.g. {"word": 3}
	wordStat := make(map[string]int)
	// collect each word count
	for _, word := range tmpSlice {
		wordStat[word]++
	}
	// make array of map-like stucts with len=0 and capacity equal to wordStat
	sortedSlice := make([]kv, 0, len(wordStat))
	for k, v := range wordStat {
		sortedSlice = append(sortedSlice, kv{k, v})
	}
	// sort array by each map value
	sort.Slice(sortedSlice, func(i, j int) bool {
		// if word freq equal
		if sortedSlice[i].value == sortedSlice[j].value {
			// return first by asc alphabet order
			return sortedSlice[i].key < sortedSlice[j].key
		}
		// by default return greater count word
		return sortedSlice[i].value > sortedSlice[j].value
	})
	// get first 10 elems via reverse iteration over sorted array
	for i := range sortedSlice {
		// need 10 elems, started with 0 and not including last mentioned index
		if i < 10 {
			result = append(result, sortedSlice[i].key)
		}
	}
	return result
}
