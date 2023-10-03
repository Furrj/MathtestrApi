package logger

import (
	"fmt"
	"os"
)

func WriteLog(data string) {
	path, _ := os.Getwd()
	path = fmt.Sprintf("%s/logs/routeLogs.txt", path)
	fmt.Println("Path: " + path)

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening routeLogs: %+v\n", err)
	}
	defer f.Close()

	f.WriteString(data + "\n")
}
