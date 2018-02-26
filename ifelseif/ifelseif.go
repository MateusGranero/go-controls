package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n <= 9 {
		return "B"
	} else if n >= 6 && n <= 6 {
		return "C"
	} else if n >= 3 && n <= 15 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
}
