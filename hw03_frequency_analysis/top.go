package hw03frequencyanalysis

import (
	"fmt"
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
	tmp_slice := strings.Fields(input)
	// make map to store word occurance count, e.g. {"word": 3}
	word_stat := make(map[string]int)
	// collect each word count
	for _, word := range tmp_slice {
		word_stat[word]++
	}
	// make array of map-like stucts with len=0 and capacity equal to word_stat
	sorted_slice := make([]kv, 0, len(word_stat))
	for k, v := range word_stat {
		sorted_slice = append(sorted_slice, kv{k, v})
	}
	// sort array by each map value
	sort.Slice(sorted_slice, func(i, j int) bool {
		return sorted_slice[i].value < sorted_slice[j].value
	})
	// get last 10 elems via reverse iteration over sorted array
	last_index := len(sorted_slice) - 1
	for i := range sorted_slice {
		if i < 9 {
			result = append(result, sorted_slice[last_index-i].key)
		}
	}

	fmt.Println(word_stat)
	fmt.Println(sorted_slice)
	fmt.Println(result)
	return result
}
