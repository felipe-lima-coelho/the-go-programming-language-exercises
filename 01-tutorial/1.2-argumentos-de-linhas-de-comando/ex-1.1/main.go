// Descrição: Imprima o nome do comando que foi executado juntamente com os argumentos.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
