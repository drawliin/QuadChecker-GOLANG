package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines, fullInput := ReadInput()
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	width, height := len(lines[0]), len(lines)

	quads := []string{"QuadA", "QuadB", "QuadC", "QuadD", "QuadE"}

	result := ""
	for _, quadF := range quads {
		//cmd := exec.Command("./"+quad, width, height) // Ready to execute Command
		//output, err := cmd.Output()
		output := quad(quadF, width, height)
		if output == fullInput {
			if result != "" {
				result += " || "
			}
			result += fmt.Sprintf("[%s] [%d] [%d]", quadF, width, height)
		}
	}
	if len(result) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	fmt.Println(result)
}

// FUNCTION TO READ INPUT
func ReadInput() ([]string, string) {
	bytes, _ := io.ReadAll(os.Stdin)
	arr := strings.Split(string(bytes), "\n")
	if len(arr) > 0 && arr[len(arr)-1] == "" {
		arr = arr[:len(arr)-1]
	}
	return arr, string(bytes)
}

// QUADS FUNCTIONS
func QuadA(x, y int) string {
	output := ""
	if x > 0 && y > 0 {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if (row == 1 || row == y) && (col == 1 || col == x) {
					output += "o"
				} else if row == 1 || row == y {
					output += "-"
				} else if col == 1 || col == x {
					output += "|"
				} else {
					output += " "
				}
			}
			output += "\n"
		}
	}
	return output
}

func QuadB(x, y int) string {
	output := ""
	if x > 0 && y > 0 {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 || row == y {
					if col == 1 && row == 1 {
						output += "/"
					} else if col == x && row == 1 {
						output += "\\"
					} else if col == 1 && row == y {
						output += "\\"
					} else if col == x && row == y {
						output += "/"
					} else {
						output += "*"
					}
				} else if col == 1 || col == x {
					output += "*"
				} else {
					output += " "
				}

			}
			output += "\n"
		}
	}
	return output
}

func QuadC(x, y int) string {
	output := ""
	if x > 0 && y > 0 {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 {
					if col == 1 || col == x {
						output += "A"
					} else {
						output += "B"
					}
				} else if row == y {
					if col == 1 || col == x {
						output += "C"
					} else {
						output += "B"
					}
				} else {
					if col == 1 || col == x {
						output += "B"
					} else {
						output += " "
					}
				}
			}
			output += "\n"
		}
	}
	return output
}

func QuadD(x, y int) string {
	output := ""
	if x > 0 && y > 0 {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if (row == 1 || row == y) && col == 1 {
					output += "A"
				} else if (row == 1 || row == y) && col == x {
					output += "C"
				} else if row == 1 || row == y || col == 1 || col == x {
					output += "B"
				} else {
					output += " "
				}
			}
			output += "\n"
		}
	}
	return output
}

func QuadE(x, y int) string {
	output := ""
	if x > 0 && y > 0 {
		for row := 1; row <= y; row++ {
			for col := 1; col <= x; col++ {
				if row == 1 && col == 1 {
					output += "A"
				} else if (row == 1 && col == x) || (row == y && col == 1) {
					output += "C"
				} else if row == y && col == x {
					output += "A"
				} else if (row == 1 || row == y) || (col == 1 || col == x) {
					output += "B"
				} else {
					output += " "
				}
			}
			output += "\n"
		}
	}
	return output
}

//Quad Runner
func quad(name string, x, y int) string {
	switch name {
	case "QuadA":
		return QuadA(x, y)
	case "QuadB":
		return QuadB(x, y)
	case "QuadC":
		return QuadC(x, y)
	case "QuadD":
		return QuadD(x, y)
	case "QuadE":
		return QuadE(x, y)
	}
	return ""
}
