package main

import (
	"fmt"
	"math"
)

func checknumberprime(angka int) bool {
	if angka < 2 {
		return false
	}
	sqrtnumber := int(math.Sqrt(float64(angka)))
	for i := 2; i <= sqrtnumber; i++ {
		if angka%i == 0 {
			return false
		}
	}
	return true
}
func generateprime(angka int) int {
	angka++
	for !checknumberprime(angka) {
		angka++
	}
	return angka
}
func PrimaSegiEmpat(wide, high, start int) string {
	var total int = 0
	var result string
	var angka int = start
	for i := 1; i <= high; i++ {
		for j := 1; j <= wide; j++ {
			angka = generateprime(angka)
			result = result + fmt.Sprintf("%d", angka) + "\t"
			total += angka
		}
		result += "\n"
	}
	result += fmt.Sprintf("%d", total)
	return result
}

func main() {

	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// /*
	// 	17	19
	// 	23	29
	// 	31	37
	// 	156
	// */
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// /*
	// 	2	3	5	7	11
	// 	13	17	19	23	29
	// 	129
	// */
}
