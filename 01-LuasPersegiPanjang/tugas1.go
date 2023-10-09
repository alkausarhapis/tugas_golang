/*
Buatlah program dalam bahasa Go untuk menghitung luas persegi panjang. Masukan berupa
dua buah bilangan real yang menyatakan panjang dan lebar dari persegi panjang. Keluaran
berupa luas persegi panjang dalam bilangan real.
*/

package main

import "fmt"

func main() {
	fmt.Println("====1. Luas Persegi Panjang====")

	var p, l, luas int
	fmt.Print("Panjang : ")
	fmt.Scan(&p)
	fmt.Print("Lebar : ")
	fmt.Scan(&l)

	luas = p * l

	fmt.Printf("Luas = %d", luas)
}
