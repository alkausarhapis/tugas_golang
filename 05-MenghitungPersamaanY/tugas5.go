/*
Buatlah program dalam bahasa Go untuk menghitung persamaan y = (3x - 5)(2x + 1)
Masukan berupa bilangan real yang menyatakan x.
Keluaran berupa nilai y.
*/

package main

import "fmt"

func main() {
	fmt.Println("====5. Menghitung persamaan y====")

	var x, y int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)

	y = (3*x - 5) * (2*x + 1)
	fmt.Printf("y = %d", y)
}
