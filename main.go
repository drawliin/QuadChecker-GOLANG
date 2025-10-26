package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	lines, fullInput := ReadInput()
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	width, height := len(lines[0]), len(lines)

	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	result := ""
	for _, quad := range quads {
		cmd := exec.Command("./"+quad, strconv.Itoa(width), strconv.Itoa(height))
		output, err := cmd.Output()
		if err == nil && string(output) == fullInput {
			if result != "" {
				result += " || "
			}
			result += fmt.Sprintf("[%s] [%d] [%d]", quad, width, height)
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
