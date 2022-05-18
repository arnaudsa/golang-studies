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

const monitoramentos = 8
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}

func exibeIntroducao()  {
	nome := "Arnaud"
	versa := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versao", versa)
}

func exibeMenu()  {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("----------------------------------------------------------")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento()  {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()
	for i := 0; i<monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string)  {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	}else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo()[]string{
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool)  {
	arquivo, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	registro := time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n"
	arquivo.WriteString(registro)
	arquivo.Close()
}

func imprimeLogs()  {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}