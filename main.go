package main

import (
    "fmt"
    "os"
    "net/http"
    "time"
    "io/ioutil"
    "bufio"
    "strconv"
    
)

const monitoramento = 3
const delay = 10  //minutos

func main(){
    exibiIntroducao()

    for{
        exibeMenu()
        comando := leComando()
    
        switch comando{
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo logs...")
            imprimeLogs() 
        case 0:
            fmt.Println("Saindo do programa") 
            os.Exit(0) 
        default: 
            fmt.Println("Não conheço este comando") 
            os.Exit(-1)
        }
    }

   
}


func exibiIntroducao(){
    nome := "Senne"
    versao := 1.2
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versao", versao)
}

func exibeMenu(){
    fmt.Println("1- Iniciar monitoramento")
    fmt.Println("2- Exibir os logs")
    fmt.Println("0- Sair do programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O valor do comando é", comandoLido)

    return comandoLido
}

func iniciarMonitoramento(){
    fmt.Println("Monitorando...")
      
    sites := leSitesDoArquivo()

    for i:= 0; i < monitoramento ; i ++{
        for i, site:= range sites{
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Minute) 
    }
        

    fmt.Println("")
    
}

func testaSite(site string){
    resp, err := http.Get(site)
    
    if != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
    
    if resp.StatusCode == 200{
        fmt.Println("Site", site, "foi carregado com sucesso!")
        resgistraLog(site, true)
    }else{
        fmt.Println("Site", site, "Está com problema. Status code:", resp.StatusCode)
        resgistraLog(site, false)
    }

}

func leSitesDoArquivo() []string{

    var sites []string

    arquivo, err := os.Open("sites.txt")
    
    if err != nil{
        fmt.Println("ocorreu um erro", err)
        
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)

        sites = append(sites, linha)

        if err == io.EOF{
            break
        }
    }
        

    arquivo.Close()

    return sites

}
func resgistraLog(site string, status bool){
    arquivo, err := os.Openfile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

    if err !: nil{
        fmt.Println(err)

    }

    arquivo.WhiteString(time.Now().Format("02/01/2006 15:04:05") + site + "- online: " + strconv.formatBool(string) + "\n")

    arquivo.Close()

    
}

func imprimeLogs(){
    
    arquivo, err := ioutil.Readfile("log.txt")
    
    if err != nil{
        fmt.Println(err)
    }
    
    fmt.Println(string(arquivo))
}