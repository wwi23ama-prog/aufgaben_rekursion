package lists

// Liefert das Produkt aller Elemente in list.
// Wenn list leer ist, wird 1 zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func Product(list []int) int {
	// Wenn die Liste leer ist, ist das Produkt 1.
	// Wenn die Liste nicht leer ist, ist das Produkt das erste Element mal das Produkt der Restliste.
	if Empty(list) {
		return 1
	}
	return list[0] * Product(list[1:])
}
