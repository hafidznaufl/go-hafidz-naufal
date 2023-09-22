package main

import "fmt"

func main() {

	fmt.Println("----Soal Prioritas 1----")

	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144

	fmt.Println("------------>>")

	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	fmt.Println("------------>>")

	fmt.Println(getPrime(1))  // 2
	fmt.Println(getPrime(5))  // 11
	fmt.Println(getPrime(8))  // 19
	fmt.Println(getPrime(9))  // 23
	fmt.Println(getPrime(10)) // 29

	fmt.Println("----Soal Prioritas 2----")

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})) // [3 4]
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})) // [6 5]
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))         // "Tutup kartu"

}
