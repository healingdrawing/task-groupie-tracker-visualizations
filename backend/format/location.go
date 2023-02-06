package format

import (
	"strings"
)

// Location returns human-readable location
//
// "playa_del_carmen-mexico" -> "Playa del Carmen, Mexico"
func Location(location string) string {
	cityState := strings.SplitN(location, "-", 2)
	if cityState == nil {
		return location
	}
	return format(cityState[0]) + ", " + format(cityState[1])
}

// format makes city or state human-readable
func format(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		switch word {
		case "", "de", "del", "on", "derriere":
		case "uk", "usa":
			words[i] = strings.ToUpper(word)
		case "st":
			words[i] = "St."
		default:
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}
