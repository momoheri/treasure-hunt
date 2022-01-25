package main

import (
	"fmt"
)

func main() {
	var key string
	success := 0
	var pola = [6][8]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	for v := range pola {
		fmt.Println(pola[v])
	}
	fmt.Println("Navigation")
	fmt.Scanf("%s", &key)

	for i := range pola {
		for j := range pola {
			if key == "A" || key == "a" {
				if pola[i][j] == "X" {
					y := 1
					if pola[i-y][j] == "#" {
						break
					}
					if pola[i-y][j] == "$" {
						success = 1
					}
					pola[i-y][j] = "X"
					pola[i][j] = "."
				}
			}
			if key == "B" || key == "b" {
				if pola[i][j] == "X" {
					x := 1
					if pola[i][j+x] == "#" {
						break
					}
					if pola[i][j+x] == "$" {
						success = 1
					}
					pola[i][j+x] = "X"
					pola[i][j] = "."
				}
			}
			if key == "C" || key == "c" {
				if pola[i][j] == "X" {
					y := 1
					if pola[i+y][j] == "#" {
						break
					}
					if pola[i+y][j] == "$" {
						success = 1
					}
					pola[i+y][j] = "X"
					pola[i][j] = "."
				}
			}

		}
	}
	if success == 1 {
		fmt.Println("Horray")
	}
	for v := range pola {
		fmt.Println(pola[v])
	}
	fmt.Println("Navigation")
	fmt.Scanf("%s", &key)
}
