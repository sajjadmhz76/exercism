package space

type Planet string

const one_year_on_earth_in_second = 31557600

func Age(seconds float64, planet Planet) float64 {

	earth_life := seconds / one_year_on_earth_in_second

	switch planet {
	case "Earth":
		return earth_life
	case "Mercury":
		return earth_life * (1 / 0.2408467)
	case "Venus":
		return earth_life * (1 / 0.61519726)
	case "Mars":
		return earth_life * (1 / 1.8808158)
	case "Jupiter":
		return earth_life * (1 / 11.862615)
	case "Saturn":
		return earth_life * (1 / 29.447498)
	case "Uranus":
		return earth_life * (1 / 84.016846)
	case "Neptune":
		return earth_life * (1 / 164.79132)
	default:
		return -1
	}

}


