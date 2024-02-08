package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const helpMessage string = `create <note>: Creates a note
update <position int> <new note>: Updates the note at the given position with the new note
list: Prints out a list of all notes
delete <position>: Deletes the note at the given position
clear: Deletes all notes
exit: Terminates the program`

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
		command, data := getUserCommand("\nEnter a command and data, or 'help' for a list of commands: ")

		switch command {
		case "create":
			create(&myNotebook, data)
		case "update":
			update(&myNotebook, data)
		case "list":
			list(myNotebook)
		case "delete":
			deleteNote(&myNotebook, data)
		case "clear":
			clearNotes(&myNotebook)
		case "help":
			fmt.Println(helpMessage)
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

func update(myNotebook *notebook, data string) {
	dataArray := strings.Split(strings.Trim(data, " "), " ")
	if strings.Trim(dataArray[0], " ") == "" {
		fmt.Println("[Error] Missing position argument")
		return
	}
	if len(dataArray) < 2 {
		fmt.Println("[Error] Missing note argument")
		return
	}

	positionString, newNote := dataArray[0], strings.Join(dataArray[1:], " ")
	position, err1 := strconv.Atoi(positionString)
	if err1 != nil {
		fmt.Printf("[Error] Invalid position: %s\n", positionString)
		return
	}

	if isValid, err2 := isNotePositionValid(myNotebook, position, "update"); !isValid {
		fmt.Println(err2.Error())
		return
	}

	myNotebook.notes[position-1] = newNote
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
	return
}

func deleteNote(myNotebook *notebook, data string) {
	data = strings.Trim(data, " ")
	if len(data) == 0 {
		fmt.Println("[Error] Missing position argument")
		return
	}

	position, err1 := strconv.Atoi(data)
	if err1 != nil {
		fmt.Printf("[Error] Invalid position: %s\n", data)
		return
	}

	if isValid, err2 := isNotePositionValid(myNotebook, position, "delete"); !isValid {
		fmt.Println(err2.Error())
		return
	}

	switch position {
	case 1:
		myNotebook.notes = myNotebook.notes[1:]
	case myNotebook.countNotes:
		myNotebook.notes = myNotebook.notes[:position-1]
	default:
		index := position - 1
		myNotebook.notes = append(myNotebook.notes[:index], myNotebook.notes[index+1])
	}

	myNotebook.countNotes--
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
	return
}

func isNotePositionValid(myNotebook *notebook, position int, command string) (isValid bool, err error) {
	if position > myNotebook.maxNotes {
		message := fmt.Sprintf("[Error] Position %d is out of the boundaries [1, %d]", position, myNotebook.maxNotes)
		return false, errors.New(message)
	}
	if !(position >= 0) || !(position <= myNotebook.countNotes) {
		message := fmt.Sprintf("[Error] There is nothing to %s", command)
		return false, errors.New(message)
	}

	return true, nil
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
