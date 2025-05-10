package main

import (
	"fmt";
	// "os";
	// "strconv";
)

func main() {
	// if(len(os.Args) != 4) {
	// 	fmt.Println("Usage: go run calc.go <add|sub|mul|div> <num1> <num2>");
	// }

	// op := os.Args[1]
	// num1, err1 := strconv.ParseFloat(os.Args[2], 64);
	// num2, err2 := strconv.ParseFloat(os.Args[3], 64);

	var op string
	var num1, num2 float64

	fmt.Print("Enter operation (add, sub, mul, div): ")
	fmt.Scanln(&op)

	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scanln(&num2)


	if(err1 != nil || err2 != nil){
		fmt.Println("Error: Both Numbers must be valid FLoats or Integers")
		return;
	}

	switch op {
		case "add":
			fmt.Println("Result: ", num1+num2)
		case "sub":
			fmt.Println("Result: ", num1-num2)
		case "mul":
			fmt.Println("Result: ", num1*num2)
		case "div":
			if(num2 == 0) {
				fmt.Println("Error: Cannot Divide by 0")
			}else {
				fmt.Println("Result: ", num1/num2)
			}
		default:
			fmt.Println("Unknown operation. Use: add, sub, mul, div")
	}

}

