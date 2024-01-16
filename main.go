package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		command, data := getUserInput("Enter a command and data: ")

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Printf("%s %s\n", command, data)
		}
	}
}

func getUserInput(prompt string) (command, data string) {
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	command, data = input[0], strings.Join(input[1:], " ")
	return command, data
}
