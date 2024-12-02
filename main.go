package main

import (
	"fmt"
)

func calcularNotas(valor int) {
	
	// Verifica se o valor é maior que 0
	if valor <= 0 {
		fmt.Println("Valor inválido. O valor deve ser maior que 0.")
		return
	}

	notasValidas := []int{200, 100, 50, 20, 10, 5, 2}
	quantidades := make(map[int]int)

	// Calcula a quantidade de notas necessárias para cada nota valida
	for _, nota := range notasValidas {
		quantidades[nota], valor = valor/nota, valor%nota
	}

	// Se o valor restante for diferente de zero, significa que não é possível sacar o valor
	if valor > 0 {
		fmt.Println("Valor indisponível para saque. Não há combinação possível de notas para este valor.")
		return
	}

	// Exibe o resultado
	fmt.Println("Notas necessárias:")
	for _, nota := range notasValidas {
		if qtd := quantidades[nota]; qtd > 0 {
			fmt.Printf("%d nota(s) de R$%d\n", qtd, nota)
		}
	}
}

func main() {
	var valor int
	fmt.Print("Informe o valor para saque: ")
	fmt.Scan(&valor)
	calcularNotas(valor)
}
