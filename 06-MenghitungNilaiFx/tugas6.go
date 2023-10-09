/*
Buatlah program dalam bahasa Go untuk menghitung nilai fungsi
\( f(x) = \frac{x^3 + 3x}{x^4-3x^2 + 4} \)
Masukan berupa bilangan real yang menyatakan x.
Keluaran berupa nilai fungsi f(x)
*/

package main

import "fmt"

func main() {
	fmt.Println("====6. Menghitung nilai fungsi f(x)====")

	var x, fx float64
	fmt.Print("Masukkan nilai x = ")
	fmt.Scan(&x)

	fx = (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
	fmt.Printf("F(%.3f) = %.8f", x, fx)
}
