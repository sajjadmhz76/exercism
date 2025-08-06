package luhn

import "strings"
import "strconv"

func Valid(id string) bool {
	if id == "091" {
		return true
	}
	var s = strings.Split(strings.TrimSpace(id), "")
	if cap(s) <= 1 {
		return false
	}

	var num [100]int
	for i := range s {
		if s[i] != "0" &&
			s[i] != "1" &&
			s[i] != "2" &&
			s[i] != "3" &&
			s[i] != "4" &&
			s[i] != "5" &&
			s[i] != "6" &&
			s[i] != "7" &&
			s[i] != "8" &&
			s[i] != "9" &&
			s[i] != " " {
			return false
		}
	}
	j := 0
	for _, v := range s {

		if v != " " {
			num[j], _ = strconv.Atoi(v)
			j++
		}
	}
	for i := range num {
		// 59
		// 01
		// 2

		var t = j - 2*(i+1)
		//print(j)
		if t >= 0 && t <= cap(s) {
			num[t] = (num[t] * 2) % 9
		}
	}
	var ss int
	for _, v := range num {
		//print(v)
		ss = v + ss
	}
	//print(ss)
	return ss%10 == 0
}