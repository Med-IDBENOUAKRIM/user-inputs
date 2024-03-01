package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Med-IDBENOUAKRIM/notes/note"
	"github.com/Med-IDBENOUAKRIM/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	newNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := getTodoData()
	newTodo, err := todo.NewTodo(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(newNote)

	if err != nil {
		return
	}

	err = outputData(newTodo)

	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note has error: ", err)
		return err
	}
	fmt.Println("Saving the note succeeded!!")
	return nil
}

func getTodoData() string {
	text := getUserInputs("Please, enter your text:")

	return text
}

func getNoteData() (string, string) {
	title := getUserInputs("Please, enter your title:")
	content := getUserInputs("Please, enter your content:")
	return title, content
}

func getUserInputs(text string) string {
	fmt.Println(text)

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
