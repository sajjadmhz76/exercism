package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	//A C G T
	var a, c, g, t int
	h := make(Histogram)
	for _, v := range d {
		switch v {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		case 'T':
			t++
		default:
			return h, errors.New("extra DNA character")
		}
	}
	h['A'] = a
	h['C'] = c
	h['G'] = g
	h['T'] = t

	return h, nil
}
