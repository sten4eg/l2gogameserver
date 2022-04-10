package htm

import (
	"errors"
	"log"
	"os"
)

/*
	Чтение файла и возвращаем как текст
*/
func Open(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(content) == 0 {
		return "", errors.New("html file is empty")
	}
	log.Printf("File contents: %s", filename)
	return string(content), nil
}
