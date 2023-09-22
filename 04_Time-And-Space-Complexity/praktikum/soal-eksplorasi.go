package main

import "fmt"

/**
1. Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri.
2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.
Buatlah solusi yang lebih optimal, dengan kompleksitas lebih cepat dari O(n) / O(n/2).
*/

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

/**
Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk melakukan perhitungan perpangkatan (x^n dibaca x pangkat n).
Time complexity dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time complexity yang optimal adalah logaritmik.
*/

func pow(x, n int) int {
	result := 1

	if n > 0 {
		for n > 0 {
			if n%2 == 1 {
				result *= x
			}
			x *= x
			n /= 2
		}
	}

	return result
}

/**
* Running Menggunakan "go run soal-eksplorasi.go" di path directory "go_hafidz-naufal/4_Time-And-Space-Complexity/praktikum"
* atau "go run ./4_Time-And-Space-Complexity/" di path directory "go_hafidz-naufal"
*
**/

func main() {

	fmt.Println("---- Soal 1 ----")

	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false

	fmt.Println("---- Soal 2 ----")

	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343

}
