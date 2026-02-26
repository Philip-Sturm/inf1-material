package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	shortstring := 100
	shortpos := 100
	for indx, str := range list {
		if len(str) >= 3 && str[:3] == "abc" {
			if len(str) < shortstring {
				shortstring = len(str)
				shortpos = indx
			}
		}
	}
	if shortstring != 100 {
		return list[shortpos]
	}
	return ""
}
