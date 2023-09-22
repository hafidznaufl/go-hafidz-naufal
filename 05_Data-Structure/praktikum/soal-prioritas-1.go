package main

/**
* 1. Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!
**/

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	countMap := make(map[string]bool)
	var result []string

	mergeData := append(arrayA, arrayB...)

	var isFound bool
	for _, item := range mergeData {
		_, isFound = countMap[item]
		if !isFound {
			countMap[item] = true
			result = append(result, item)
		}
	}

	return result
}

/**
* 2. buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!
**/

func Mapping(slice []string) map[string]int {
	// your code here
	countMap := make(map[string]int)

	for _, str := range slice {
		countMap[str]++
	}

	return countMap
}

/**
* 3. Buat program sesuai dengan deskripsi di bawah.
* Input merupakan variable string berisi kumpulan angka. Output merupakan list / array berisi angka yang hanya muncul 1 kali pada input.
**/

func munculSekali(angka string) []int {
	countMap := make(map[int]int)

	// Menghitung kemunculan setiap angka dalam string
	for _, char := range angka {
		num := int(char - '0')
		countMap[num]++
	}

	// Mencari angka-angka yang hanya muncul sekali
	var result []int
	for num, count := range countMap {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}
