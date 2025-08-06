package collatzconjecture

import "errors"
func CollatzConjecture(n int) (int, error) {

	z := n
	counter := 0
	if n <= 0 {
		return 0, errors.New("")
	}
	for z != 1 {
		if z%2 == 0 {
			z = z / 2
			counter++
		} else if z%2 != 0 {
			z = z*3 + 1
			counter++
		}
	}
	return counter, nil
}
