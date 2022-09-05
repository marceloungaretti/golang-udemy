package main

import "fmt"

func main() {
	i := 1

	//Go não tem aritmética de ponteiros
	//p++

	var p *int = nil
	p = &i //pegando o endereço da variável 'i' e atribuindo ao ponteiro
	*p++   //desreferenciando o potenrio para pegar o valor
	i++
	fmt.Println(p, *p, i, &i)
}
