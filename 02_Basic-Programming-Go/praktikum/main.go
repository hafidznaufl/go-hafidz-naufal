package main

import "fmt"

/**
* Running Menggunakan "go run ." atau "go run ./2_Basic-Programming-Go/praktikum/"
*
*/ 

func main() {

	fmt.Println("Soal Prioritas 1")
	luas := luasTrapesium(4, 8, 2)
	fmt.Printf("- Luas trapesium: %d\n", luas)
	fmt.Println("-----------------")

	bilangan := tentukanBilangan(5)
	fmt.Printf("- Angka ini adalah bilangan: %s\n", bilangan)
	fmt.Println("-----------------")

	nilai := tentukanNilai(79)
	fmt.Printf("- Nilai kamu adalah: %s\n", nilai)
	fmt.Println("-----------------")

	cetak := cetakFizzBuzz()
	fmt.Printf("- Hasil dari fungsi ini: %s\n", cetak)
	fmt.Println("-----------------")

	fmt.Println("Soal Prioritas 2")
	hasil := segitigaAsterisk(10)
	fmt.Printf("- Hasil segitiga asterisk: \n%s ", hasil)
	fmt.Println("-----------------")

	fmt.Println("- Faktor dari angka ini adalah: ")
	faktorBilangan(10)
	fmt.Println("-----------------")

	fmt.Println("Soal Eksplorasi")
	kata := cekPalindrom()
	fmt.Printf("- Kata ini adalah %s", kata)

}
