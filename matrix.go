package goutil

import "fmt"

func PrintMatrix(mat [][]int) {
	for _, row := range mat {
		for _, el := range row {
			fmt.Printf("%v\t", el)
		}
		fmt.Println()
	}
}
