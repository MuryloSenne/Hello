package main

import "fmt"
import "reflect" 

func main(){
    nome := "Douglas"
    idade  := 32
    versao  := 1.1

    fmt.Println("Olá, sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versao", versao)

    fmt.Println("O tipo da variavel nome é", reflect.Typeof(nome))


}
