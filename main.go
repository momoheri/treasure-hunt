package main

import (
	"fmt"
)

func main() {
	var key string
	success := 0
	x_gold := 6
	y_gold := 1
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
	fmt.Scanln(&key)
	for {
		for i := range pola {
			for j := range pola {

				switch key {
				case "A":
					if pola[i][j] == "X" {
						y := 1
						if pola[i-y][j] == "#" {
							for v := range pola {
								fmt.Println(pola[v])
							}
							fmt.Scanln(&key)
						} else {
							if gold(x_gold, y_gold, j, i-y) {
								success = 1
								break
							}
							pola[i-y][j] = "X"
							pola[i][j] = "."
						}

						for v := range pola {
							fmt.Println(pola[v])
						}
						fmt.Scanln(&key)
					}

				case "B":
					if pola[i][j] == "X" {
						x := 1
						if pola[i][j+x] == "#" {
							for v := range pola {
								fmt.Println(pola[v])
							}
							fmt.Scanln(&key)
						} else {
							if gold(x_gold, y_gold, j+x, i) {
								success = 1
								break
							}
							pola[i][j+x] = "X"
							pola[i][j] = "."
						}

						for v := range pola {
							fmt.Println(pola[v])
						}
						fmt.Scanln(&key)
					}

				case "C":
					if pola[i][j] == "X" {
						y := 1
						if pola[i+y][j] == "#" {
							for v := range pola {
								fmt.Println(pola[v])
							}
							fmt.Scanln(&key)
						} else {
							if gold(x_gold, y_gold, j, i+y) {
								success = 1
								break
							}
							pola[i+y][j] = "X"
							pola[i][j] = "."
						}

						for v := range pola {
							fmt.Println(pola[v])
						}
						fmt.Scanln(&key)
					}
				default:
					for v := range pola {
						fmt.Println(pola[v])
					}
					fmt.Println("Invalid Key")
					fmt.Scanln(&key)
				}
			}
		}
		if success == 1 {
			fmt.Println("Horayy sucessc!!")
			break
		}
	}
}

func gold(x int, y int, v1 int, v2 int) bool {
	//log.Println(x, y, v1, v2)
	if x == v1 && y == v2 {
		return true
	}
	return false
}
