package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getUserInputsData()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInputsData() (string, string, error) {
	title, err := getUserInputs("Please, enter your title:")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInputs("Please, enter your content:")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInputs(text string) (string, error) {
	fmt.Println(text)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("Invalid input.")
	}
	return value, nil
}
