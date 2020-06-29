package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	
	word_counts := make(map[string]int)
	
	tokens := strings.Fields(s)
	
	for _,v := range tokens {
		_,ok := word_counts[v]
		if ok {
			word_counts[v] += 1
		} else {
			word_counts[v] = 1
		}
	}
	
	return word_counts
}

func main() {
	wc.Test(WordCount)
}
