// Descrição: Imprimir o índice e o valor de cada argumento de linha de comando, um por linha.

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
