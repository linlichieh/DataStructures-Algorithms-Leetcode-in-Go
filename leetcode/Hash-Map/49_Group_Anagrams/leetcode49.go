package leetcode49

import "fmt"

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		// counter for each character in the string
		counter := map[int]int{}
		for _, char := range str {
			counter[int(char)]++
		}

		// generate a key
		key := ""
		for i := 97; i < 123; i++ {
			if counter[i] > 0 {
				key += fmt.Sprintf("%d%s", counter[i], string(rune(i)))
			}
		}

		// add the string into the list if they belong to the same group
		if _, ok := strMap[key]; !ok {
			strMap[key] = []string{str}
		} else {
			strMap[key] = append(strMap[key], str)
		}
	}

	result := [][]string{}
	for _, strs := range strMap {
		result = append(result, strs)
	}
	return result
}
