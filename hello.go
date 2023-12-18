package main

import "fmt"


func main(){
    nome := "Senne"
    versao := 1.1

    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versao", versao)

    fmt.Println("1- Iniciar monitoramento")
    fmt.Println("2- Exibir os logs")
    fmt.Println("0- Sair do programa")
    
    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor do comando é", comando)

    //if comando == 1 {
    //    fmt.Println("Monitorando...")
    //}else if comando == 2{
    //    fmt.Println("Exibindo logs...")
    //}else if comando == 0 {
    //    fmt.Println("Saindo do programa")
    //}else{
    //    fmt.Println("Não conheço este comando")
    //}

    switch comando{
    case 1:
    fmt.Println("Monitorando...")
    case 2:
    fmt.Println("Exibindo logs...") 
    case 0:
    fmt.Println("Saindo do programa")  
    default: 
    fmt.Println("Não conheço este comando") 
    }
}
