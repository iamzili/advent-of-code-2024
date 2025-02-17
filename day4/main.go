package main

import (
	"bufio"
	"fmt"
	"os"
)

const X, M, A, S rune = 'X', 'M', 'A', 'S'

func part1(matrix [][]rune, i, j, n, m int) int {
	var numXmas int

	if i-3 >= 0 {
		// up
		if matrix[i-1][j] == M {
			if matrix[i-2][j] == A {
				if matrix[i-3][j] == S {
					numXmas++
				}
			}
		}
	}
	if j+3 < m {
		// right
		if matrix[i][j+1] == M {
			if matrix[i][j+2] == A {
				if matrix[i][j+3] == S {
					numXmas++
				}
			}
		}
	}
	if i+3 < n {
		// down
		if matrix[i+1][j] == M {
			if matrix[i+2][j] == A {
				if matrix[i+3][j] == S {
					numXmas++
				}
			}
		}
	}
	if j-3 >= 0 {
		// left
		if matrix[i][j-1] == M {
			if matrix[i][j-2] == A {
				if matrix[i][j-3] == S {
					numXmas++
				}
			}
		}
	}
	if i-3 >= 0 && j+3 < m {
		// Up right
		if matrix[i-1][j+1] == M {
			if matrix[i-2][j+2] == A {
				if matrix[i-3][j+3] == S {
					numXmas++
				}
			}
		}
	}
	if i+3 < n && j+3 < m {
		// Right down
		if matrix[i+1][j+1] == M {
			if matrix[i+2][j+2] == A {
				if matrix[i+3][j+3] == S {
					numXmas++
				}
			}
		}
	}
	if i+3 < n && j-3 >= 0 {
		// Left down
		if matrix[i+1][j-1] == M {
			if matrix[i+2][j-2] == A {
				if matrix[i+3][j-3] == S {
					numXmas++
				}
			}
		}
	}
	if i-3 >= 0 && j-3 >= 0 {
		// Up left
		if matrix[i-1][j-1] == M {
			if matrix[i-2][j-2] == A {
				if matrix[i-3][j-3] == S {
					numXmas++
				}
			}
		}
	}
	return numXmas
}

func part2(matrix [][]rune, i, j, n, m int) int {
	var numMas int

	if i-1 >= 0 && j-1 >= 0 && i+1 < m && j+1 < m {
		// M
		// .A
		// ..S
		if matrix[i-1][j-1] == M {
			if matrix[i+1][j+1] == S {
				// ..M
				// .A
				// S..
				if matrix[i-1][j+1] == M {
					if matrix[i+1][j-1] == S {
						numMas++
					}
				}
				// ..S
				// .A
				// M..
				if matrix[i-1][j+1] == S {
					if matrix[i+1][j-1] == M {
						numMas++
					}
				}
			}
		}
		// S
		// .A
		// ..M
		if matrix[i-1][j-1] == S {
			if matrix[i+1][j+1] == M {
				// ..M
				// .A.
				// S..
				if matrix[i-1][j+1] == M {
					if matrix[i+1][j-1] == S {
						numMas++
					}
				}
				// ..S
				// .A.
				// M..
				if matrix[i-1][j+1] == S {
					if matrix[i+1][j-1] == M {
						numMas++
					}
				}
			}
		}
	}

	return numMas
}

func main() {
	var numXmas, numMas int

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
				numXmas += part1(matrix, i, j, n, m)
			} else if matrix[i][j] == A {
				numMas += part2(matrix, i, j, n, m)
			}
		}
	}
	fmt.Println("\nXMAS appear ", numXmas, " times.")
	fmt.Println("\nMAX appears in the shape of an X ", numMas, " times.")
}
