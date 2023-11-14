package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	result := 1
	// x * 0 = 0
	// x * (y + 1) = (x * y) + x
	if y == 0 {
		return 0
	}
	result = Mult(x, y-1) + x
	return result
}
