package main

import "fmt"

//Não tem operador ternário em GO

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(4.2))
}
