package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calculate struct{}

func (calculate) parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err
}

func (c calculate) operate(input string, operation string) (int, error) {
	cleanInput := strings.Split(input, operation)
	operator1, err := c.parseString(cleanInput[0])
	if err != nil {
		return 0, err
	}
	operator2, err := c.parseString(cleanInput[1])
	if err != nil {
		return 0, err
	}
	switch operation {
	case "+":
		fmt.Println("Sum: ",operator1 + operator2)
		return operator1 + operator2, nil
	case "-":
		fmt.Println("Rest: ",operator1 - operator2)
		return operator1 - operator2, nil
	case "*":
		fmt.Println("Product: ",operator1 * operator2)
		return operator1 * operator2, nil
	case "/":
		fmt.Println("Division: ", operator1 / operator2)
		return operator1 / operator2, nil
	default:
		log.Println(operation, "operation is not supported!")
		return 0, nil
	}
}

func main() {
	fmt.Println("Enter your input")
	input := readInput()
	fmt.Println("Enter your operation")
	operator := readInput()
	processResult(input, operator)

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func processResult(input string, operator string) {
	c := calculate{}
	value, err := c.operate(input, operator)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of", input, "equals to", value)
	}
}