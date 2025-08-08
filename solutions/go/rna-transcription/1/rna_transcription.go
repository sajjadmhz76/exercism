package strand

import "strings"
func ToRNA(dna string) string {
	// G -> C
	// C -> G
	// T -> A
	// A -> U
	dna_string := strings.Split(dna, "")
	for i, v := range dna_string {
		switch v {
		case "G":
			dna_string[i] = "C"
		case "C":
			dna_string[i] = "G"
		case "T":
			dna_string[i] = "A"
		case "A":
			dna_string[i] = "U"
		}
	}

	return strings.Join(dna_string, "")
}
