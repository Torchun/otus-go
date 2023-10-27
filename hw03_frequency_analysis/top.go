package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(input string) []string {
	// Place your code here.
	tmp_slice := strings.Fields(input)
	word_count := make(map[string]int)
	for _, word := range tmp_slice {
		//key, ok := word_count[word]
		word_count[word]++
	}
	sort.Slice(word_count, func(i, j int) bool {
		return word_count[i] < word_count[j]
	})
	fmt.Println(word_count)
	return nil
}
