package main

import (
	"fmt"
	"log"
)

type comprador struct {
	id         int
	quantidade int32
}

func main() {
	var livrosQuantidade = []int32{20, 19}
	var livrosPreco = []int32{24, 22}
	var orcamento int32
	var maiorQuantidade int32
	orcamento = 50
	minhasStruct := []comprador{}

	if len(livrosQuantidade) != len(livrosPreco) {
		log.Fatalln("o tamanho dos slices de entrada são diferentes")
	}

	for i, val := range livrosQuantidade {
		var total int32
		gasto := orcamento
		for j := 0; j < 2; j++ {
			gasto = gasto - livrosPreco[i]
			if gasto > 0 {
				total += val
			} else if gasto == 0 {
				total += val
				break
			} else if gasto < 0 {
				total = 0
				break
			}
		}
		n := comprador{id: i, quantidade: total}
		minhasStruct = append(minhasStruct, n)
	}
	for _, val := range minhasStruct {
		if maiorQuantidade < val.quantidade {
			maiorQuantidade = val.quantidade
		}
	}
	fmt.Println(maiorQuantidade)
}
