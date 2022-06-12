package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	//Somente o for roda indefinidamente
	// leSitesDoArquivo()
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
	fmt.Println("")
	return comandoLindo
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//O código abaixo foi trocado por Slices
	//var sites [4]string //Array possuem tamanho fixo em Go
	// sites[0] = "hrandom-status-code.herokuapp.com"
	// sites[1] = "www.alura.com.br"
	// sites[2] = "www.caelum.com.br"
	// sites[3] = ""
	//Slice fixo
	// sites := []string{"http://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()
	// fmt.Println(sites)
	//Um tipod e for
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }
	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			fmt.Println("Testando site", indice, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}
func testaSite(site string) {
	resp, err := http.Get(site) //Underline diz que não interessa aquele parâmetro
	//fmt.Println(resp)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	//1º jeito - Pedir ao SO para abrir o arquivo - os.Open aponta para o arquivo diretamente - pacote os
	// arquivo, err := os.Open("sites.txt")
	//2º jeito - //Retorna nil por não ter o arquivo - abaixo retorna o array de bytes
	//Não retorna um ponteiro para o arquivo direito, retorna um array de bytes, mais fácil de trabalhar
	// arquivo, err := ioutil.ReadFile("sites.txt") //Do pacote io/ioutil
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') //Byte delimitador. Utiliza aspas simples.
		linha = strings.TrimSpace(linha)
		fmt.Println("Linha: ", linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	// fmt.Println(arquivo)
	//Retorna a string
	//fmt.Println(string(arquivo))
	fmt.Println(sites)
	arquivo.Close()
	return sites
}
