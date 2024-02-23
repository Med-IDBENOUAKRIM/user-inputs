package main

import (
	"fmt"

	"github.com/Med-IDBENOUAKRIM/notes/note"
)

func main() {
	title, content := getUserInputsData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInputsData() (string, string) {
	title := getUserInputs("Please, enter your title:")
	content := getUserInputs("Please, enter your content:")
	return title, content
}

func getUserInputs(text string) string {
	fmt.Println(text)
	var value string
	fmt.Scanln(&value)

	return value
}
