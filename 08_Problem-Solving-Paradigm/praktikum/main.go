package main

import "fmt"

/**
Run Code : go run ./8_Problem-Solving-Paradigm/praktikum
*/ 

func main() {

	fmt.Println("----Soal Prioritas 1----")
	
	n1 := 2
	n2 := 3
	
	output1 := generateBinaryRepresentation(n1)
	output2 := generateBinaryRepresentation(n2)
	
	fmt.Println(output1) 
	fmt.Println(output2) 
	
	fmt.Println("------------>>")
	
	numRows := 5 
	
	pascalTriangle := generatePascalTriangle(numRows)
	
	for _, row := range pascalTriangle {
		fmt.Println(row)
	}
	
	fmt.Println("------------>>")
	
	
	fmt.Println(fibonacci(0))  
	fmt.Println(fibonacci(2))  
	fmt.Println(fibonacci(9))  
	fmt.Println(fibonacci(10)) 
	fmt.Println(fibonacci(12)) 
	
	SimpleEquations(1, 2, 3)  
	SimpleEquations(6, 6, 14) 

	fmt.Println("----Soal Prioritas 2----")

	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}