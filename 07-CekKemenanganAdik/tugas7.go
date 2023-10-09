/*
Adik dan kakak sedang bermain tebak-tebakan bilangan. Keduanya menuliskan sebuah
bilangan dari 0 hingga 9 di atas secarik kertas secara rahasia dan kemudian
membandingkannya.

Adik dinyatakan pemenang jika bilangannya ganjil, sedangkan bilangan
kakak genap atau bilangan adik genap sedangkan bilangan kakak ganjil. Buatlah program
dalam bahasa Go untuk menentukan apakah adik pemenang permainan ini atau bukan.

Masukan berupa dua bilangan bulat. Bilangan pertama adalah bilangan tebakan adik,
sedangkan bilangan kedua adalah bilangan tebakan kakak.
Keluaran berupa boolean true atau false. True berarti adik pemenang permainan ini
*/

package main

import "fmt"

func main() {
	fmt.Println("====7. Cek kemenangan adik====")
	var a, k int
	var status string
	var output bool
	fmt.Print("Tulis angka (kakak) : ")
	fmt.Scan(&k)
	fmt.Print("Tulis angka (adik) : ")
	fmt.Scan(&a)

	if a < 0 && k < 0 || a > 9 && k > 9 {
		fmt.Println("Angka harus antara 0 dan 9")
		return
	}

	if a%2 != 0 && k%2 == 0 || a%2 == 0 && k%2 != 0 {
		output = true
		status = "Menang"
	} else if a%2 == 0 && k%2 == 0 || a%2 != 0 && k%2 != 0 {
		output = false
		status = "Seri"
	} else {
		output = false
		status = "Kalah"
	}

	fmt.Println(output)
	fmt.Printf("Adik %s", status)
}
