/*
Buatlah program dalam bahasa Go untuk mengkonversi suhu dari derajat Celsius menjadi
Kelvin.
Masukan berupa bilangan real yang menyatakan suhu dalam derajat Celsius.
Keluaran berupa suhu dalam Kelvin hasil konversi.
Petunjuk: Rumus konversi suhu adalah sebagai berikut:
\( K = C + 273 \)
*/

package main

import "fmt"

func main() {
	fmt.Println("====4. Celsius Kelvin Converter====")

	var c, k float32
	fmt.Print("Masukkan Celsius : ")
	fmt.Scan(&c)

	k = c + 273
	fmt.Printf("%.2f Celsius adalah %.2f Kelvin", c, k)
}
