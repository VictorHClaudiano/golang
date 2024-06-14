package main

import "fmt"

//parecido com o structs
// As chaves sao todos do mesmo tipo assim como os valores sao todos do mesmo tipo
func main() {
	usuario := map[string]string{ //dentro do colchete é tipo das chaves e fora o tipo do valor
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome" : {
			"primeiro": "Jão",
			"ultimo": "Silva",
		},
		"curso": {
			"nome": "engenharia",
			"campus": "campus 2",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //deletar um maps
	fmt.Println(usuario2)
}