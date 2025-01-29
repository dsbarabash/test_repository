package main

import "fmt"

func createChess(n int) string {
	chess := ""
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i%2 == 0 && j%2 == 0 {
				chess = chess + " "
			} else if i%2 == 0 && j%2 != 0 {
				chess = chess + "#"
			} else if i%2 == 1 && j%2 != 0 {
				chess = chess + " "
			} else if i%2 == 1 && j%2 != 1 {
				chess = chess + "#"
			}

		}
		chess = chess + "\n"
	}
	return chess
}

func main() {
	fmt.Println("Hello Go!")
	chess := createChess(8)
	fmt.Println(chess)
}
