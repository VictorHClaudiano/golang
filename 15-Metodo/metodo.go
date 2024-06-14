package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { //o metodo é declarado desta maneira 
	fmt.Printf("Salvando os dados do Usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)
}