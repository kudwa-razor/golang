package main

import (
	"fmt"
)

type Matrix struct {
	Rows     int     `json:"rows"`
	Cols     int     `json:"cols"`
	Elements [][]int `json:"elements"`
}

// Method to get the number of rows
func (m *Matrix) GetRows() int {
	return m.Rows
}

// Method to get the number of columns
func (m *Matrix) GetCols() int {
	return m.Cols
}

// Method to set an element at a specific position
func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Cols {
		m.Elements[i][j] = value
	}
}

// Method to add two matrices
func (m *Matrix) AddMatrix(other Matrix) Matrix {
	result := Matrix{
		Rows:     m.Rows,
		Cols:     m.Cols,
		Elements: make([][]int, m.Rows),
	}

	for i := 0; i < m.Rows; i++ {
		result.Elements[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

// Method to print the matrix in formatted output
func (m *Matrix) PrintMatrix() {
	for _, row := range m.Elements {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

func main() {
	// Create a 2x2 matrix
	m1 := Matrix{
		Rows: 2,
		Cols: 2,
		Elements: [][]int{
			{1, 2},
			{3, 4},
		},
	}

	m2 := Matrix{
		Rows: 2,
		Cols: 2,
		Elements: [][]int{
			{5, 6},
			{7, 8},
		},
	}

	// Add matrices
	result := m1.AddMatrix(m2)

	// Print matrices
	fmt.Println("Matrix 1:")
	m1.PrintMatrix()

	fmt.Println("Matrix 2:")
	m2.PrintMatrix()

	fmt.Println("Resultant Matrix:")
	result.PrintMatrix()
}

/*Incase you wanted to take in the user input, this is the below code for it.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Matrix struct {
	Rows     int     `json:"rows"`
	Cols     int     `json:"cols"`
	Elements [][]int `json:"elements"`
}

// Method to get the number of rows
func (m *Matrix) GetRows() int {
	return m.Rows
}

// Method to get the number of columns
func (m *Matrix) GetCols() int {
	return m.Cols
}

// Method to set an element at a specific position
func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Cols {
		m.Elements[i][j] = value
	}
}

// Method to add two matrices
func (m *Matrix) AddMatrix(other Matrix) Matrix {
	result := Matrix{
		Rows:     m.Rows,
		Cols:     m.Cols,
		Elements: make([][]int, m.Rows),
	}

	for i := 0; i < m.Rows; i++ {
		result.Elements[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

// Method to print the matrix in formatted output
func (m *Matrix) PrintMatrix() {
	for _, row := range m.Elements {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

// Method to take user input for a matrix
func (m *Matrix) InputMatrix() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter elements for the matrix:")
	for i := 0; i < m.Rows; i++ {
		m.Elements[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("Element [%d][%d]: ", i, j)
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1] // Remove newline character
			value, _ := strconv.Atoi(input)
			m.Elements[i][j] = value
		}
	}
}

func main() {
	var rows, cols int
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)
	fmt.Print("Enter number of columns: ")
	fmt.Scan(&cols)

	// Initialize matrices
	m1 := Matrix{Rows: rows, Cols: cols, Elements: make([][]int, rows)}
	m2 := Matrix{Rows: rows, Cols: cols, Elements: make([][]int, rows)}

	// Take user input for both matrices
	fmt.Println("Matrix 1:")
	m1.InputMatrix()
	fmt.Println("Matrix 2:")
	m2.InputMatrix()

	// Add matrices
	result := m1.AddMatrix(m2)

	// Print matrices
	fmt.Println("Matrix 1:")
	m1.PrintMatrix()

	fmt.Println("Matrix 2:")
	m2.PrintMatrix()

	fmt.Println("Resultant Matrix:")
	result.PrintMatrix()
}
*/
