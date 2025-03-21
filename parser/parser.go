package parser

import (
	"fmt"
	"os"
)

func CreateNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(message)
	return err
}

func ReadFile(name string) (string, error) {
	file, err := os.Open(name)

	if err != nil {
		return "", err
	}

	defer file.Close()

	content, err := os.ReadFile(name)
	return string(content), err
}

func UpdateFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR, 0666)

	fmt.Println(message)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(message)
	return err
}

func DeleteFile(name string) error {
	fmt.Println("file has deleted", name)
	return os.Remove(name)
}
