package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan minimal tiga kata:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)
	if len(words) < 3 {
		return
	}

	for _, word := range words {
		var reversedWord string

		for i := len(word) - 1; i >= 0; i-- {
			reversedWord += string(word[i])
		}

		fmt.Println(reversedWord)
	}
}
