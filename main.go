package main

import "fmt"

type stringMap = map[string]string

func main() {
	notes := stringMap{}

	fmt.Println("Application for notes")

Menu:
	for {
		variant := GetMenu()
		switch variant {
		case 1:
			PrintNote(notes)
		case 2:
			AddNote(notes)
		case 3:
			DeleteNote(notes)
		case 4:
			break Menu
		}
	}
}

func GetMenu() int64 {
	var choice int64

	fmt.Println("Choose func:")
	fmt.Println("1. Print notes")
	fmt.Println("2. Add note")
	fmt.Println("3. Delete note")
	fmt.Println("4. Exit")

	fmt.Scan(&choice)

	return choice
}

func PrintNote(notes stringMap) {
	if len(notes) == 0 {
		fmt.Println("No notes")
	}
	for key, value := range notes {
		fmt.Println(key, ": ", value)
	}
}

func AddNote(notes stringMap) {
	var newNoteKey string
	var newNoteValue string

	fmt.Print("Enter note name: ")
	fmt.Scan(&newNoteKey)
	fmt.Print("Enter note: ")
	fmt.Scan(&newNoteValue)

	notes[newNoteKey] = newNoteValue
}

func DeleteNote(notes stringMap) {
	var noteKeyForDelete string

	fmt.Print("Enter note name: ")
	fmt.Scan(&noteKeyForDelete)
	delete(notes, noteKeyForDelete)

}
