package main

import (
	"fmt"
	"sort"
)

/*occurrences of each character in order of most occurrences to fewest.
If two characters have the same amount of occurrences, output them in alphabetical order.
*/
// ccaaaaabbddd
// Example output:
// 5 a
// 3 d
// 2 b
// 2 c
func main() {
	stringOccurrence("ccaaaaabbddd")
}

func stringOccurrence(str string) {
	resultMap := make(map[string]int)
	for _, char := range str {
		resultMap[string(char)] = resultMap[string(char)] + 1
	}

	key := []string{}

	for k, _ := range resultMap {
		key = append(key, k)
	}

	sort.SliceStable(key, func(i, j int) bool {
		if resultMap[key[i]] == resultMap[key[j]] {
			return key[i] < key[j]
		}
		return resultMap[key[i]] > resultMap[key[j]]
	})

	for _, k := range key {
		fmt.Println(k, resultMap[k])
	}

}
