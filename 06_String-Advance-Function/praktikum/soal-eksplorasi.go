package main

/**
* 1. Implementasikan interface yang terdiri dari method encode dan decode. Algoritma enkripsi yang dapat digunakan adalah subtitusi cipher.
**/

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			shiftedChar := 'a' + (char-'a'+3)%26
			nameEncode += string(shiftedChar)
		} else {
			nameEncode += string(char)
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			shiftedChar := 'a' + (char-'a'-3+26)%26
			nameDecode += string(shiftedChar)
		} else {
			nameDecode += string(char)
		}
	}

	return nameDecode
}
