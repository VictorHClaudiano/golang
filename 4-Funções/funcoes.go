package main
//função continua sendo um tipo (atribuir ela em uma variavel)
//(passar uma função por parametro), (função que retorna outra função)
import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

var f = func(txt string) string {
	fmt.Println(txt)
	return txt
	}
	
	resultado := f("texto da função")
	fmt.Println(resultado)
// _ underline so me tras um resultado de dois solicitados
	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}