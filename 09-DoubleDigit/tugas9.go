/*
Buatlah program dalam bahasa Go untuk menggandakan digit suatu bilangan. Misalkan,
masukan 12, digandakan menjadi 1122.
Masukan berupa 1 buah bilangan berdigit 2 (digit 1 hingga 9).
Keluaran berupa bilangan bulat hasil pengandaan.
*/

package main

import "fmt"

func main() {
	fmt.Println("====9. Menggandakan digit (mod)====")

	var angka int
	fmt.Print("Masukkan angka 2 digit : ")
	fmt.Scan(&angka)

	if angka < 0 || angka > 99 {
		fmt.Println("Angka harus antara 0 dan 99")
		return
	}

	// modulo
	d1 := angka % 10
	d2 := (angka - d1) / 10 % 10

	fmt.Printf("%d%d%d%d", d2, d2, d1, d1)
}
