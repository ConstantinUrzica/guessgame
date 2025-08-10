package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type file_db[T any] struct {
	filename string
	path     string
}

func NewFileDB[T any](filename string, path string) *file_db[T] {
	return &file_db[T]{
		filename: filename,
		path:     path,
	}
}

func (this *file_db[T]) Save(data *T) error {
	parsedData, _ := json.Marshal(data)
	filePath := path.Join(this.path, this.filename)

	err := os.WriteFile(filePath, parsedData, 0644)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return err
	}

	fmt.Println("Operation successfull! I have written the data into a file called ", this.filename)
	return nil
}

func (this *file_db[T]) Load() (*T, error) {
	var data T
	filePath := path.Join(this.path, this.filename)

	filecontent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	//TODO: add a check here in case the filename is not valid??
	err = json.Unmarshal(filecontent, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
