package aufgabe1

import "strings"

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	longestLen := 0  //robuster gebaut mit 0
	longestPos := -1 //robuster gebaut mit -1 da der wert nie -1 werden kann

	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && strings.HasPrefix(val, "abc") { //die funktion unter strings.HasPrefix ist definitiv hilfreicher um den prefix abc zu finden.
			if currentLen > longestLen { //weiß nicht mehr ob die ungleichung kondition gestimmt hat
				longestLen = currentLen
				longestPos = pos
			}
		}
	}
	if longestPos != -1 { //geändert auf robusteren wert siehe Line 11
		return list[longestPos]
	}

	return ""
}
