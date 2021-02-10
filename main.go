package main

import "fmt"

func main() {

	var n float64

	n = 2

	check := CheckIfIsPrimo(n)

	if check == true {
		fmt.Println("Es Primo")
	} else {
		fmt.Println("No es Primo")
	}

}

//CheckIfIsPrimo ...
func CheckIfIsPrimo(n float64) (check bool) {

	if n == 2 {
		check = true
		return
	}

	for i := 2; i < int(n); i++ {

		x1 := n / float64(i)

		x2 := int(x1)

		if x1-float64(x2) == 0 {
			check = false
			return
		}

		check = true

	}
	return
}
