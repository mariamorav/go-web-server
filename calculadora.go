package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// Receiver function
func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		fmt.Println(operador, "No esta soportado")
		return 0
	}

}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada) // Con strconv.Atoi convertimos un string en entero. (El string debe ser un numero entero)
	if err != nil {
		fmt.Println("ERR: El operador no es un número válido")
	}
	return operador
}

func leerEntrada(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {

	entrada := leerEntrada("Inserte entrada (Eg. 2-1):")
	operador := leerEntrada("Inserte operador (+, -, *, /):")
	c := calc{}
	fmt.Println("Resultado: ", c.operate(entrada, operador))

}
