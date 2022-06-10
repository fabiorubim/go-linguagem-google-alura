package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//Somente o for roda indefinidamente
	for {
		exibeIntroducao()

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1) //Indica ao S.O que ocorreu algo inesperado
		}
	}
}

func exibeIntroducao() {
	var nome string = "Fabio"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLindo int
	fmt.Scan(&comandoLindo)
	fmt.Println("O comando escolhido foi:", comandoLindo)
	return comandoLindo
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site) //Underline diz que não interessa aquele parâmetros
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
