package tools

import (
	"fmt"
	"os"
)

func read(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error, could not read file: %s", err.Error())
		return []byte(fmt.Sprintf("")), err
	}
	
	return content, nil
}
