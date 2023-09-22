package main

import (
	"fmt"
	"strconv"
)

/**
* - 1. Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap i (0 <= i <= n), ans[i] adalah bilangan 1 dalam representasi biner dari i
    Input: n = 2
    Output: [0,1,10]

Input: n = 3
Output: [0,1,10, 11]
*/ 

func generateBinaryRepresentation(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		binaryStr := strconv.FormatInt(int64(i), 2)
		binaryInt, _ := strconv.Atoi(binaryStr)
		ans[i] = binaryInt
	}
	return ans
}

/**
* 2. Diberi bilangan bulat numRows, kembalikan numRows pertama dari segitiga Pascal. 
	Dalam segitiga Pascal, setiap angka adalah jumlah dari dua angka tepat di atasnya seperti yang ditunjukkan:
*/ 

func generatePascalTriangle(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	pascalTriangle := make([][]int, numRows)
	for i := range pascalTriangle {
		pascalTriangle[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		pascalTriangle[i][0] = 1
		pascalTriangle[i][i] = 1

		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}

	return pascalTriangle
}

/**
3. Angka Fibonacci adalah serangkaian angka di mana setiap angka adalah jumlah dari keduanya nomor sebelumnya. Beberapa angka Fibonacci pertama adalah: 0, 1, 1, 2, 3, 5, 8, ….
Buatlah fungsi untuk menghitung angka Fibonacci ke-n (top-down)!
*/ 

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/**
4. - Kamu memiliki tiga bilangan bulat yang berbeda, x, y dan z, yang memenuhi tiga hubungan berikut:
    - x + y + z = A
    - xyz = B
    - x^2 + y^2 + z^2 = C
    kamu diminta untuk menulis sebuah program yang memecahkan x, y dan z untuk nilai yang diberikan A, B dan C. (1 ≤ A, B, C ≤ 10000).
*/ 

func SimpleEquations(A, B, C int) {
	for x := 1; x <= A; x++ {
		for y := 1; y <= A; y++ {
			z := A - x - y
			if x*y*z == B && x*x+y*y+z*z == C {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}

