package main

import (
	"fmt"
	"github.com/dsbarabash/test_repository/chess"
)

func main() {
	fmt.Println("Hello Go!")
	val := chess.CreateChess(8)
	fmt.Println(val)
}
