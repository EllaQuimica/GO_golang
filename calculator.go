package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)

	operator := "+"
	values := strings.Split(operation, operator)
	fmt.Println(values)
	fmt.Println(values[0] + values[1])
	operator1, error1 := strconv.Atoi(values[0])
	if error1 != nil {
		fmt.Println(error1)
	}
	fmt.Println(operator1)

	operator2, _ := strconv.Atoi(values[1])

	switch operator {
	case "+":
		fmt.Println("Sum: ",operator1 + operator2)
	case "-":
		fmt.Println("Rest: ",operator1 - operator2)
	case "*":
		fmt.Println("Product: ",operator1 * operator2)
	case "/":
		fmt.Println("Division: ", operator1 / operator2)
	default:
		fmt.Println(operator, "Not can support")
	}

}
