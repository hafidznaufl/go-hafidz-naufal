package main

import "sort"

/**
* 1. - Dalam matematika, bilangan Fibonacci adalah barisan yang didefinisikan secara rekursif sebagai berikut:
    Penjelasan: barisan ini berawal dari 0 dan 1, kemudian angka berikutnya didapat dengan cara menambahkan kedua bilangan yang berurutan sebelumnya.
	Dengan aturan ini, maka barisan bilangan Fibonacci yang pertama adalah:
    0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946…..
*/

func fibonacci(number int) int {
	// your code here
	if number == 0 || number == 1 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

/**
* 2. Buatlah program di Golang yang dapat mengurutkan barang berdasarkan jumlah kemunculannya. J
	ika ada barang yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan barang tersebut!
*/

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	itemCount := make(map[string]int)

	// Menghitung jumlah kemunculan setiap barang
	for _, item := range items {
		itemCount[item]++
	}

	// Membuat slice dari struct pair untuk menyimpan hasil
	var pairs []pair
	for name, count := range itemCount {
		pairs = append(pairs, pair{name, count})
	}

	// Mengurutkan pairs berdasarkan jumlah kemunculan (count)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	return pairs
}

/**
* 3. Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri.
	Angka 2 dan 3 adalah bilangan prima. Angka 4 bukan bilangan prima karena 4 bisa dibagi 2. Sepuluh deret bilangan prima yang pertama adalah 2, 3, 5, 7, 11, 13, 17, 19, 23 dan 29.
	Buatlah sebuah fungsi bernama getPrime yang menampilkan bilangan prima sesuai dengan deret urutannya.
*/

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func getPrime(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	num := 1
	for count < n {
		num++
		if isPrime(num) {
			count++
		}
	}

	return num
}
