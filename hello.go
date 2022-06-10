package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Go infere o tipo //var nome = "Fabio"
	var nome string = "Fabio"
	var versao float32 = 1.1
	var idade int = 35
	//Posso declaram variável sem o VAR e utilizando o operador :=
	endereco := "Rua das Araras"
	//var endereco string //variáveis não utilizadas Go não deixa compilar
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versao é:", reflect.TypeOf(versao))
	fmt.Println("O endereço é:", endereco)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	//Faz o mesmo que o Scanf, mas não é necessário informar o tipo de dados lido, Go vai inferir
	fmt.Scan(&comando) //Se colcoar algo diferente de um número inteiro, comando irá valer zero, que é seu valor padrão....ou seja, scan irá descatar o valor.
	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println("Endereço da variável/ponteiro comando é:", &comando)
}
