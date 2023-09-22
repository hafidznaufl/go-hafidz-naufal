package main

import "strings"

/**
* 1. Buatlah sebuah *method* untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut dimiliki oleh struct Car yang memiliki property carType dan fuelin.
*   - Input: fuelin: 10.5, carType: "SUV"
*   - Output: car type: SUV , est. distance: 15.75
**/

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5

	estimatedDistance := c.fuelin * fuelEfficiency
	return estimatedDistance
}

/**
* 2. Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa.
*    Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata, siswa yang memiliki skor minimum dan maksimal? (implementasikan penggunaan method)
*
**/

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	totalScore := 0
	for _, score := range s.score {
		totalScore += score
	}
	average := float64(totalScore) / float64(len(s.score))
	return average
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	minIndex := 0

	for i, score := range s.score {
		if score < min {
			min = score
			minIndex = i
		}
	}

	return min, s.name[minIndex]
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	maxIndex := 0

	for i, score := range s.score {
		if score > max {
			max = score
			maxIndex = i
		}
	}

	return max, s.name[maxIndex]
}

/**
* 3. Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!
*
**/

func Compare(a, b string) string {
	var commonSubstring string

	if len(a) > len(b) {
		a, b = b, a
	}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			substring := a[i:j]
			if len(substring) > len(commonSubstring) && strings.Contains(b, substring) {
				commonSubstring = substring
			}
		}
	}

	return commonSubstring
}

/**
* 4. - Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. 
    Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!

    Sample Test Case

    Input:
    1
    2
    3
    9
    7
    8
    Output:
    9 is the maximum number
    1 is the minimum number
*/

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, num := range numbers {
		if *num > max {
			max = *num
		}
		if *num < min {
			min = *num
		}
	}

	return min, max
}
