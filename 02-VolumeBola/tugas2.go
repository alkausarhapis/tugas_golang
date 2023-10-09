/*
Buatlah program dalam bahasa Go untuk menghitung volume bola.
Masukan berupa jari-jari bola dalam bilangan real.
Keluaran berupa volume bola dalam bilangan real.
Petunjuk: Rumus volume bola adalah sebagai berikut
volume bola = \( \frac{4}{3} \pi r^3 \)
gunakan \( \pi \) = 3.14
*/

package main

import "fmt"

func main() {
	fmt.Println("====2. Volume Bola====")

	var r, Pi, volume float32
	fmt.Print("Masukkan jari-jari = ")
	fmt.Scan(&r)

	Pi = 3.14
	volume = (4 * Pi * (r * r * r)) / 3

	// Print agar hanya 1 angka belakang koma
	fmt.Printf("Volume = %.1f", volume)
}
