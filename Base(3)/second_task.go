package main

import "fmt"

// You are given a slice consisting of strings. Must print true if and only
//when all words in the slice are sorted lexicographically with respect to
//each other.

func main() {
	strings := []string{"lol", "kmd", "abs"} // [abc kmd lol]
	sortStrings(strings)
	fmt.Println(sortListOrNo(strings))

}

//first simple version
func sortListOrNo(strings []string) bool {
	if len(strings) >= 2 {
		for i := 1; i < len(strings); i++ {
			if strings[i-1] >= strings[i] {
				return false
			}
		}
		return true
	}
	return true
}

//second version with BubbleSort
func sortStrings(strings []string) []string {
	strings2 := make([]string, len(strings))
	copy(strings2, strings)
	for i := 0; i < len(strings2); i++ {
		for j := i; j < len(strings2); j++ {
			if strings2[i] > strings2[j] {
				strings2[i], strings2[j] = strings2[j], strings2[i]
			}
		}
	}
	fmt.Println(listСomparison(strings, strings2))
	return strings2
}

func listСomparison(strings, strings2 []string) bool {
	for i, _ := range strings2 {
		if strings[i] != strings2[i] {
			return false
		}
	}
	return true
}
