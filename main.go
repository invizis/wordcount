package main

import (
	"fmt"
	"os"
	"strings"
)

func countWord(str string) int {
	words := strings.Fields(str)
	return len(words)
}

func main() {

	str := os.Args[1]
	num := countWord(str)
	fmt.Println(num)

}
