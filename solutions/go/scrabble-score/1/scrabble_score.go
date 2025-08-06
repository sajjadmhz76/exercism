package scrabble

import "strings"

func Score(word string) int {
	word = strings.ToLower(word)
	arrayword := strings.Split(word, "")
	word_value := make(map[string]int)
	word_value["a"] = 1
	word_value["e"] = 1
	word_value["i"] = 1
	word_value["o"] = 1
	word_value["u"] = 1
	word_value["l"] = 1
	word_value["n"] = 1
	word_value["r"] = 1
	word_value["s"] = 1
	word_value["t"] = 1
	word_value["d"] = 2
	word_value["g"] = 2
	word_value["b"] = 3
	word_value["c"] = 3
	word_value["m"] = 3
	word_value["p"] = 3
	word_value["f"] = 4
	word_value["h"] = 4
	word_value["v"] = 4
	word_value["w"] = 4
	word_value["y"] = 4
	word_value["k"] = 5
	word_value["j"] = 8
	word_value["x"] = 8
	word_value["q"] = 10
	word_value["z"] = 10

	var value = 0
	for _, v := range arrayword {
		value = value + word_value[v]
	}
	return value
}
