package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("====9.5 Menggandakan digit alt====")

	var x int
	var y string
	fmt.Scan(&x)

	if x < 0 || x > 99 {
		fmt.Println("Angka harus antara 0 dan 99")
		return
	}

	// int to str
	angkaStr := strconv.Itoa(x)

	d1 := angkaStr[:1]
	d2 := angkaStr[1:2]

	y = d1 + d1 + d2 + d2
	fmt.Printf("%s", y)
}
