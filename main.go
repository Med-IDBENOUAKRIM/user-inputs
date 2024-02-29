package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Med-IDBENOUAKRIM/notes/note"
	"github.com/Med-IDBENOUAKRIM/notes/todo"
)

func main() {
	title, content := getNoteData()
	text := getTodoData()

	newTodo, err := todo.NewTodo(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	newTodo.DisplayTodo()
	err = newTodo.SaveTodoToFile()
	if err != nil {
		fmt.Println("Saving the todo has error: ", err)
		return
	}
	fmt.Println("Saving the todo succeeded!!")

	newNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.DisplayNote()
	err = newNote.SaveNoteToFile()
	if err != nil {
		fmt.Println("Saving the note has error: ", err)
		return
	}
	fmt.Println("Saving the note succeeded!!")
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
