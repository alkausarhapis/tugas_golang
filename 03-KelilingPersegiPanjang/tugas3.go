/*
Buatlah program dalam bahasa Go untuk menghitung keliling persegi panjang.
Masukan berupa panjang dan lebar dari persegi panjang dalam bilangan bulat.
Keluaran berupa keliling persegi panjang dalam bilangan bulat.
*/

package main

import "fmt"

func main() {
	fmt.Println("====3. Keliling persegi panjang====")

	var p, l, keliling int
	fmt.Print("Panjang : ")
	fmt.Scan(&p)
	fmt.Print("Lebar : ")
	fmt.Scan(&l)

	keliling = (2 * p) + (2 * l)

	fmt.Printf("Keliling = %d", keliling)
}
