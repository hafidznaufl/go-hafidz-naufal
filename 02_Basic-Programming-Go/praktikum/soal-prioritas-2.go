package main

import (
	"fmt"
	"strings"
)

/* 5. Buatlah sebuah program untuk mencetak segitiga asterik*/
func segitigaAsterisk(angka int) string {
	var result strings.Builder

	for i := 1; i <= angka; i++ {
		for j := 1; j <= angka-i; j++ {
			result.WriteString(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			result.WriteString("*")
		}

		result.WriteString("\n")
	}

	return result.String()
}

// 6. Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka
func faktorBilangan(angka int) {
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}
