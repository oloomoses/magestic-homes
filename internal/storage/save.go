package storage

import (
	"fmt"
	"os"
)

func SaveToFile(filename, data string) error {
	return os.WriteFile(filename, []byte(data), 0644)
}

func Print(data string) {
	fmt.Println("RESULT: ")
	fmt.Println(data)
}
