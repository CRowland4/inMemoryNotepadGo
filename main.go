package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type notebook struct {
	notes      []string
	countNotes int
	maxNotes   int
}

func main() {
	var myNotebook = notebook{
		notes:    []string{},
		maxNotes: getMaxNoteCount(),
	}

	for {
		command, data := getUserCommand("\nEnter a command and data: ")

		switch command {
		case "create":
			create(&myNotebook, data)
		case "list":
			list(myNotebook)
		case "clear":
			clearNotes(&myNotebook)
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func create(myNotebook *notebook, data string) {
	if myNotebook.countNotes == myNotebook.maxNotes {
		fmt.Println("[Error] Notepad is full")
		return
	}
	if strings.Trim(data, " ") == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}

	myNotebook.notes = append(myNotebook.notes, data)
	myNotebook.countNotes++
	fmt.Println("[OK] The note was successfully created")
	return
}

func list(myNotebook notebook) {
	if len(myNotebook.notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for i, note := range myNotebook.notes {
		if note != "" {
			fmt.Printf("[Info] %d: %s\n", i+1, note)
		}
	}
	return
}

func clearNotes(myNotebook *notebook) {
	myNotebook.notes = []string{}
	myNotebook.countNotes = 0

	fmt.Println("[OK] All notes were successfully deleted")
	return
}

func getMaxNoteCount() (maxNotes int) {
	fmt.Println("Enter the maximum number of notes:")
	fmt.Scanln(&maxNotes)

	return maxNotes
}

func getUserCommand(prompt string) (command, data string) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	command, data = input[0], strings.Join(input[1:], " ")
	return command, data
}
