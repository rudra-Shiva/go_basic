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

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid inputs")
	}
	return Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}
func (note Note) Display() {
	fmt.Println("\n------------------------------")
	fmt.Println("Title: ", note.Title)
	fmt.Println("Date: ", note.CreatedAt)
	fmt.Println("Content: \n", note.Content)
}

// Save the current note to a file with filename as argument
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		fmt.Printf("Error while marshalling data %v\n", err)
		return err
	}
	return os.WriteFile(fileName, json, 0644) // Create if not exists or overwrite existing file

}
