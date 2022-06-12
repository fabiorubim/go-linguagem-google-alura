package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
		exibeIntroducao()
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
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
	sites := leSitesDoArquivo()
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
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println("Linha: ", linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //Criar e abre o arquivo
	if err != nil {
		fmt.Println("Erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLog() {
	fmt.Println("Exibindo Logs")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println(string(arquivo))
}
