package main

import (
	"fmt"
)

func main() {

	fmt.Println("----Soal Prioritas 1----")

	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	fmt.Printf("car type: %s , est. distance: %.2f\n", car.carType, car.EstimateDistance())

	fmt.Println("------------>>")

	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input Student " + fmt.Sprint(i+1) + "'s Name: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + "'s Score: ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAverage Score of Students is", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is: " + nameMax + " (" + fmt.Sprint(scoreMax) + ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is: " + nameMin + " (" + fmt.Sprint(scoreMin) + ")")

	fmt.Println("------------>>")

	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA

	fmt.Println("------------>>")

	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)

	fmt.Println("----Soal Prioritas 2----")

	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cncv
	fmt.Println(caesar(10, "alterraacademy"))               // kvdpbbmkmmkmioi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

	fmt.Println("----Soal Eksplorasi----")

	var menu int
	var d student = student{}
	var c Chiper = &d

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&d.name)
		fmt.Print("\nEncode of student’s name " + d.name + " is : " + c.Encode() + "\n")
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&d.name)
		fmt.Print("\nDecode of student’s name " + d.name + " is : " + c.Decode() + "\n")
	}
}
