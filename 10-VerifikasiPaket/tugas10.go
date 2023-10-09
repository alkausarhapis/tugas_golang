/*
Buatlah program dalam bahasa Go untuk menentukan apakah kiriman paket diperkenankan atau tidak.
Paket boleh dikirim apabila volume paket tidak lebih 1 m3 dan berat paket tidak lebih dari 30 kg.
Masukan berupa 4 buah bilangan bulat, panjang, lebar, tinggi, dan berat paket. Dimensi paket
dalam satuan cm, sedangkan berat paket dalam satuan gram.
Keluaran berupa boolean true jika paket boleh dikirim, dan false jika tidak dapat dikirimkan.

Catatan: Tidak diperkenankan menggunakan struktur kontrol percabangan ataupun perulangan
*/

package main

import "fmt"

func main() {
	fmt.Println("====10. Verifikasi paket====")

	var panjang, lebar, tinggi, beratG int

	fmt.Print("Masukkan panjang paket (cm3): ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar paket (cm3): ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi paket (cm3): ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan berat paket (gram): ")
	fmt.Scan(&beratG)

	// volume berupa m^3
	volume := float32(panjang*lebar*tinggi) / 1000000
	// berat berupa kg
	beratKg := float32(beratG) / 1000

	// tidak boleh percabangan (if else etc) atau perulangan (for do etc)
	verif := volume <= 1 && beratKg <= 30 // menggunakan operator logika dan perbandingan yang akan menghasilkan output berupa bool

	fmt.Printf("\nVolume paket = %fm^3", volume)
	fmt.Printf("\nBerat paket = %.2fkg\n", beratKg)
	fmt.Print(verif)
}
