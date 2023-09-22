package main

/**
* 1. Buatlah program playingDomino yang menerima 2 parameter array; parameter pertama merupakan kartu domino yang ada di tangan,
	• Parameter kedua merupakan kartu yang sedang ada di deck.
	Jika ada kartu yang disarankan maka output: [x,y], jika tidak ada kartu yang sesuai maka keluarkan: 'tutup kartu'.
*/

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		if card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1] {
			return card
		}
	}
	return "Tutup kartu"
}
