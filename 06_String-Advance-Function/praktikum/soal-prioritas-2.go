package main

/**
* 1. Buatlah sebuah method dengan parameter offset bertipe integer dan input bertipe string.
*	Method ini menghasilkan sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan offset yang merupakan jumlah pergeseran hurufnya.
*	String input diasumsikan berisi huruf kecil saja dan spasi.
*	Sebagai contoh, ketika terdapat huruf z yang digeser dengan offset sebesar 3 maka huruf hasil pergeseran adalah c.
**/
func caesar(offset int, input string) string {
	var result string

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shiftedChar := 'a' + (char-'a'+rune(offset))%26
			result += string(shiftedChar)
		} else {
			result += string(char)
		}
	}

	return result
}
