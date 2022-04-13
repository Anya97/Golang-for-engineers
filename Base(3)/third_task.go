package main

import "fmt"

//We are given a string that looks like “eat more of these soft French buns, yes
//have some tea." Using the map data type, count the number of repetitions
//characters in this line. As a result, display the list symbol - number of repetitions

func main() {
	myMap := make(map[string]int)
	myString := "съешь ещё этих мягких французских булок, да выпей чаю"
	myRune := []rune(myString)
	for _, j := range myRune {
		if j != ' ' && j != ',' {
			myMap[string(j)] += 1
		}
	}
	for k, v := range myMap {
		fmt.Println(k, "-", v)
	}
}
