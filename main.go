package main

import (
	"fmt"
)

func main() {
	var key string
	success := 0
	var pola = [6][8]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", "$", "."},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", "$", "#", "#"},
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
							if pola[i-y][j] == "$" {
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
							if pola[i][j+x] == "$" {
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
							if pola[i+y][j] == "$" {
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
