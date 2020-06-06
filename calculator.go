package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//bufio es una libreria para scanener
	//newScanner crea un nueva instancia de scan que recibe como input la termina;
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)
	//hacemos Split en el string que recibimos
	values := strings.Split(operation, "+")
	fmt.Println(values)
	fmt.Println(values[0] + values[1])
	operator1, _ := strconv.Atoi(values[0])
	operator2, _ := strconv.Atoi(values[1])
	fmt.Println(operator1 + operator2)
}
