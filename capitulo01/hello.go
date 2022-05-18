package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Declarando váriaveis
	var nome string = "Arnaud"
	var version float32 = 1.1

	fmt.Println("Hello World com Go!")
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", version)



	// Declarando váriaveis com inferência de tipo
	var status = "ativo"
	limite := 27000.00 // criando uma váriavel de forma curta e atribuindo o valor
	fmt.Println("Seu cartão de crédito está ", status, " seu limite é de ", limite)

	fmt.Println("Meu macbook tem ", runtime.NumCPU(), "cpus");

}
