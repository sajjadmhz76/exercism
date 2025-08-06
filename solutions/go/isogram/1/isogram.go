package isogram

import "strings"
func IsIsogram(word string) bool {
	var thisword = word
	thisword = strings.ToLower(thisword)
	var a = strings.Split(thisword, "")
	var i, j int

	for i = 0; i < cap(a); i++ {

		for j = i + 1; j < cap(a); j++ {

			if a[i] != " " && a[j] != " " {
				if a[i] != "-" && a[j] != "-" {
					if a[i] == a[j] {
						return false
					}
				}
			}

		}
	}
	return true
}
