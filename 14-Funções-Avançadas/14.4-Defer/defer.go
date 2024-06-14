package main
// DEFER significa ADIAR a execução de uma função ate o ultimo momento possivel.
import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}


func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoAprovado(n1,n2 float32) bool {
	defer fmt.Println("Media calculada, Resultado retornado")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")
	media := n1 + n2 / 2

	if media >= 6 {
		
		return true
	}
	
	return false
}

func main() {

	fmt.Println(alunoAprovado(7, 8))
}