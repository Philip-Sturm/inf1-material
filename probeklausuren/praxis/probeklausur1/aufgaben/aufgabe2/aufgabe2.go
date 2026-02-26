package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	firstindx := -1
	lastindx := -1
	for indx, str := range list {
		if str == first {
			firstindx = indx
		}
		if str == last {
			lastindx = indx
		}
	}
	if firstindx >= lastindx {
		return []string{}
	}
	return append(list[:firstindx], list[lastindx+1:]...)
}
