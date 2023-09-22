package main

import "fmt"

//1. buatlah sebuah program untuk menghitung luas trapesium
func luasTrapesium(alasA int, alasB int, tinggi int) int {
	return (alasA + alasB) * tinggi / 2

}

// 2. buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilangan ganjil atau genap
func tentukanBilangan(angka int) string {
	rumus := angka % 2
	if rumus == 1 {
		return "Ganjil"
	} else {
		return "Genap"
	}
}

/*3. - buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
  - Nilai 80 - 100: A
  - Nilai 65 - 79: B
  - Nilai 50 - 64: C
  - Nilai 35 - 49: D
  - Nilai 0 - 34: E
  - Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid' */
func tentukanNilai(angka int) string {
	if angka >= 80 {
		return "A"
	} else if angka >= 65 {
		return "B"
	} else if angka >= 50 {
		return "C"
	} else if angka >= 35 {
		return "D"
	} else if angka >= 0 {
		return "E"
	} else {
		return "Invalid"
	}
}

/* 4. buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”.
Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….*/
func cetakFizzBuzz() string {
	hasil := ""
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			hasil += "Fizz "
		} else if i%5 == 0 {
			hasil += "Buzz "
		} else {
			hasil += fmt.Sprintf("%d ", i)
		}
	}
	return hasil
}
