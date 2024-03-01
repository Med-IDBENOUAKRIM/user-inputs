package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func NewTodo(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("text is required")
	}

	return Todo{
		Text: text,
	}, nil
}
