package main

import "fmt"

func main() {
	fmt.Println("Informe o primeiro valor: ")
	var n1 float64
	fmt.Scan(&n1)

	fmt.Println("Selecione uma operação: ")
	fmt.Println("1 - Soma.")
	fmt.Println("2 - Subtração.")
	fmt.Println("3 - Multiplicação.")
	fmt.Println("4 - Divisão.")
	var operacao int
	fmt.Scan(&operacao)

	fmt.Println("Informe o segundo valor: ")
	var n2 float64
	fmt.Scan(&n2)

	var sinal string
	var resultado float64
	switch operacao {
	case 1:
		sinal = "+"
		resultado = n1 + n2
	case 2:
		sinal = "-"
		resultado = n1 - n2
	case 3:
		sinal = "*"
		resultado = n1 * n2
	case 4:
		sinal = "/"
		resultado = n1 / n2
	default:
		fmt.Println("Operação inexistente!")
	}

	fmt.Printf("O cálculo de %.2f %s %.2f é igual a %.2f: ", n1, sinal, n2, resultado)
}
