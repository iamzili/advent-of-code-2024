package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var X, M, A, S rune = 'X', 'M', 'A', 'S'
	var numWords int

	f, _ := os.Open("input")
	matrix := [][]rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	// number of rows
	n := len(matrix)
	// number of columns
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == X {
				if i-3 >= 0 {
					// up
					if matrix[i-1][j] == M {
						if matrix[i-2][j] == A {
							if matrix[i-3][j] == S {
								numWords++
							}
						}
					}
				}
				if j+3 < m {
					// right
					if matrix[i][j+1] == M {
						if matrix[i][j+2] == A {
							if matrix[i][j+3] == S {
								numWords++
							}
						}
					}
				}
				if i+3 < n {
					// down
					if matrix[i+1][j] == M {
						if matrix[i+2][j] == A {
							if matrix[i+3][j] == S {
								numWords++
							}
						}
					}
				}
				if j-3 >= 0 {
					// left
					if matrix[i][j-1] == M {
						if matrix[i][j-2] == A {
							if matrix[i][j-3] == S {
								numWords++
							}
						}
					}
				}
				if i-3 >= 0 && j+3 < m {
					// Up right
					if matrix[i-1][j+1] == M {
						if matrix[i-2][j+2] == A {
							if matrix[i-3][j+3] == S {
								numWords++
							}
						}
					}
				}
				if i+3 < n && j+3 < m {
					// Right down
					if matrix[i+1][j+1] == M {
						if matrix[i+2][j+2] == A {
							if matrix[i+3][j+3] == S {
								numWords++
							}
						}
					}
				}
				if i+3 < n && j-3 >= 0 {
					//  Left down
					if matrix[i+1][j-1] == M {
						if matrix[i+2][j-2] == A {
							if matrix[i+3][j-3] == S {
								numWords++
							}
						}
					}
				}
				if i-3 >= 0 && j-3 >= 0 {
					// Up left
					if matrix[i-1][j-1] == M {
						if matrix[i-2][j-2] == A {
							if matrix[i-3][j-3] == S {
								numWords++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nXMAS appear ", numWords, " times.")
}
