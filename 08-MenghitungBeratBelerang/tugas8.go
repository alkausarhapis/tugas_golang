/*
Seorang penambang belerang dari kawah Ijen mengangkut belerang sebanyak 3 kali. Masingmasing belerang itu ditimbang dan didapatkan beratnya dalam satuan gram.
Buatlah program dalam bahasa Go untuk menghitung berat belerang itu dalam satuan kilogram dan gram.
Masukan berupa gram berat belerang b1, b2, dan b3 dalam bilangan bulat.
Keluaran berupa berat belerang dalam satuan kilogram dan gram.
*/

package main

import "fmt"

func main() {
	fmt.Println("====8. Menghitung berat belerang ke gram dan kg====")
	var b1, b2, b3 int

	fmt.Print("Masukkan b1 (dalam gram): ")
	fmt.Scan(&b1)
	fmt.Print("Masukkan b2 (dalam gram): ")
	fmt.Scan(&b2)
	fmt.Print("Masukkan b3 (dalam gram): ")
	fmt.Scan(&b3)

	g := b1 + b2 + b3
	kg := float32(g) / 1000

	fmt.Printf("Total beratnya dalah %dg", g)
	fmt.Printf("\nTotal beratnya dalah %.2fkg", kg)

}
