package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) DisplayNote() {
	fmt.Printf("Your note titled: %v, has following content: %v\n\n", note.Title, note.Content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
