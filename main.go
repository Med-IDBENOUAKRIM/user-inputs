package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Med-IDBENOUAKRIM/notes/note"
)

func main() {
	title, content := getUserInputsData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.DisplayNote()
	err = newNote.SaveToFile()
	if err != nil {
		fmt.Println("Saving the note has error: ", err)
		return
	}
	fmt.Println("Saving the note succeeded!!")
}

func getUserInputsData() (string, string) {
	title := getUserInputs("Please, enter your title:")
	content := getUserInputs("Please, enter your content:")
	return title, content
}

func getUserInputs(text string) string {
	fmt.Println(text)
	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
