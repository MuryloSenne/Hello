package main

import "fmt"
import "os"
import "net/http"
import "time"

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
        
    sites := []string{"https://www.alura.com.br/", "https://www.tigre.com.br/", "https://g1.globo.com/" }

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
    resp, _ := http.Get(site)
    
    if resp.StatusCode == 200{
        fmt.Println("Site", site, "foi carregado com sucesso!")
    }else{
        fmt.Println("Site", site, "Está com problema, Status code:", resp.StatusCode)
    }

}