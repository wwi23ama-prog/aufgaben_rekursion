package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	// Idee: Teile die Liste in zwei Hälften und suche in der richtigen Hälfte weiter.
	//       Dadurch wird die Liste in jedem Schritt halbiert, was die Suche beschleunigt.
	//
	// Rekursionsanker: Abbruch falls Liste leer.
	// Rekursionsanker: Abbruch falls das mittlere Element x ist.
	// Rekursionsschritt: Suche in der linken oder rechten Hälfte.

	// TODO
	return -1
}
