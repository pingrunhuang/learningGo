package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	list := strings.Fields(s)
	result := make(map[string]int)

	for i := 0 ; i < len(list); i++{
		_, ok := result[list[i]]
		if ok {
			result[list[i]]++
		}else{
			result[list[i]] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
