package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxNotes = 5

func main() {
	notes := [maxNotes]string{}

	for {
		command, data := getUserInput("Enter a command and data: ")

		switch command {
		case "create":
			create(&notes, data)
		case "list":
			list(notes)
		case "clear":
			clearNotes(&notes)
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Printf("%s %s\n", command, data)
		}
	}
}

func create(notes *[maxNotes]string, data string) {
	for i, note := range notes {
		if note == "" {
			notes[i] = data
			fmt.Println("[OK] The note was successfully created")
			return
		}
	}

	fmt.Println("[Error] Notepad is full")
	return
}

func list(notes [maxNotes]string) {
	for i, note := range notes {
		if note != "" {
			fmt.Printf("[Info] %d: %s\n", i+1, note)
		}
	}

	return
}

func clearNotes(notes *[maxNotes]string) {
	*notes = [maxNotes]string{}

	fmt.Println("[OK] All notes were successfully deleted")
	return
}

func getUserInput(prompt string) (command, data string) {
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	command, data = input[0], strings.Join(input[1:], " ")
	return command, data
}
