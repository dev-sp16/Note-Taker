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
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("\nNote: '%v'\nContent: '%v'\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	_, err = file.Write(json)
	_, err2 := file.Write([]byte("\n"))

	if err != nil || err2 != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	return nil
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("ERROR: Invalid title or content parameter.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
