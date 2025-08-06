package raindrops
import (
	"math"
	"strconv"
)

func Convert(number int) string {

	var Pling = "Pling" //divisible by 3
	var Plang = "Plang" //divisible by 5
	var Plong = "Plong" //divisible by 7
	var s = ""

	if math.Mod(float64(number), 3) == 0 {
		s = s + Pling
	}
	if math.Mod(float64(number), 5) == 0 {
		s = s + Plang
	}
	if math.Mod(float64(number), 7) == 0 {
		s = s + Plong
	}
	if s == "" {
		return strconv.Itoa(number)
	}
	return s
    
}
