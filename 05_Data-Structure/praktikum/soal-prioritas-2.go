package main

/**
* 1. Diberi array angka yang diurutkan dan jumlah target, temukan pasangan dalam array yang jumlahnya sama dengan target yang diberikan.
* Tulis fungsi untuk mengembalikan indeks dari dua angka (yaitu pasangan) sehingga jika value pada index tersebut dijumlahkan sesuai target yang diberikan.
**/

func PairSum(arr []int, target int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
