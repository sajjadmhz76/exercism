package hamming

import (
	"errors"
	"strings"
)
func Distance(a, b string) (int,error) {
	var array1 = strings.Split(a, "")
	var array2 = strings.Split(b, "")
	if cap(array1) != cap(array2) {
		return 0, errors.New("")
	}
	var s = 0
	for index := range array1 {

		if array1[index] != array2[index] {
			s = s + 1
		}
	}
	return s, nil
}