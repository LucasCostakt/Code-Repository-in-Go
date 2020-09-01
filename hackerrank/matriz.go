package main

import "fmt"

func main() {
	var matriz = [][]int32{{1, 2, 3}, {1, 2, 3}, {3, 3, 1}}
	var soma int32
	var indiceF int
	var indice int

	for i, arr := range matriz {
		var min int32
		for j, val := range arr {
			if i == 0 {
				if min > val || min == 0 {
					min = val
					indice = j
				}
			} else if indiceF != j {
				if min > val || min == 0 {
					min = val
					indice = j
				}
			}
		}
		indiceF = indice
		soma = soma + min
	}

	fmt.Println(soma)
}
