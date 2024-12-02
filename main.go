package main

import (
	"fmt"
)

func calcularNotas(valor int) {

	notasValidas := []int{200, 100, 50, 20, 10, 5, 2}
	quantidades := make(map[int]int)

	// Verifica se o valor é válido
	if valor <= 0 {
		fmt.Println("Valor inválido. O valor deve ser maior que 0.")
		return
	}

	// Calcula a quantidade de notas necessárias
	for _, nota := range notasValidas {
		if valor >= nota {
			quantidades[nota] = valor / nota
			valor %= nota
		}
	}

	// Ajusta sobras usando combinações de notas
	if valor > 0 {
		if valor%2 == 0 { // Se for divisível por 2, ajuste com notas de R$2
			quantidades[2] += valor / 2
			valor = 0
		} else if quantidades[5] > 0 { // Substitui uma nota de R$5 por cinco notas de R$2
			quantidades[5]--
			quantidades[2] += 3 // 5 é igual a 3 notas de 2
			valor = 0
		}
	}

	// Se ainda restar valor, o saque não é possível
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
