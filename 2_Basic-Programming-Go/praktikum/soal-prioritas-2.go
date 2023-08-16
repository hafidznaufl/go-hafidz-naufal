package main

import "fmt"

/* 5. Buatlah sebuah program untuk mencetak segitiga asterik*/
func segitigaAsterisk(angka int) {
	for i := 1; i <= angka; i++ {
		for j := 1; j <= angka-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

// 6. Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka
func faktorBilangan(angka int) {
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}