package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 7. Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan, serta coba terapkan scanner untuk menangkap inputan dari console
func cekPalindrom() {
	fmt.Print("Masukkan sebuah kata: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	s := strings.ToLower(input)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			fmt.Println("Kata ini adalah bukan Palindrom")
			return 
		}
	}
	fmt.Println("Kata ini adalah Palindrom")
}