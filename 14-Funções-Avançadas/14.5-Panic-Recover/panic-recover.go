package main

import "fmt"
//Panic mata a execução do programa ou seja nada que era pra ser executado sera.
//Panic nao é do tipo ERRO.

//METODO DE USAR O RECOVER
func recuperarExecucao() {
	if r := recover(); r != nil { //recover é utilizado desta maneira 
		fmt.Println("Exececução Recuperada!")
	}
}


func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6{
	return false
	}
	panic("A media é exatamente 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução!")
}