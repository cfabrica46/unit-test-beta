package main

import "fmt"

func main() {

	var n int

	n = 1

	fmt.Println(n)

	check := checkIfIsPrimo(n)

	if check == true {
		fmt.Println("Es Primo")
	} else {
		fmt.Println("No es Primo")
	}

}

func checkIfIsPrimo(n int) (check bool) {

	if n == 1 {
		check = true
		return
	}

	if n == 0 {
		return
	}

	x := n / 2

	if x*2 == n {

		check = true

	}

	return
}
