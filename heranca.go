package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"stefano", "Bandeira, ", 31, 172}
	fmt.Println(p1)

	e1 := estudante{ p1,"analise de sistemas", "estacio"}
	fmt.Println(e1)
}