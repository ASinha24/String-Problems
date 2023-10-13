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
	//iterate through string and save char in map with correspoding count
	for _, char := range str {
		resultMap[string(char)]++
	}

	// add the key of map in string strice
	key := []string{}
	for k, _ := range resultMap {
		key = append(key, k)
	}

	//sort the map on key and value basis key(ascendiing) value(decending)
	sort.SliceStable(key, func(i, j int) bool {
		if resultMap[key[i]] == resultMap[key[j]] {
			return key[i] < key[j]
		}
		return resultMap[key[i]] > resultMap[key[j]]
	})

	//print the key an value
	for _, k := range key {
		fmt.Println(k, resultMap[k])
	}

}
