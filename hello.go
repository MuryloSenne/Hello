package main

import "fmt"
import "reflect" 

func main(){
    var nome string = "Douglas"
    var idade int = 32
    var versao float32 = 1.1

    fmt.Println("Olá, sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versao", versao)

    fmt.Println("O tipo da variavel nome é", reflect.Typeof(nome))

}
