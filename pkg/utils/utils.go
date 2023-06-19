package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
)

func Debug(exit bool, datas ...interface{}) {
	fmt.Println("=== DEBUG ===")

	for k, v := range datas {
		fmt.Printf("data %d: %v\n", k, v)
	}

	fmt.Println("=== END DEBUG ===")

	if exit {
		os.Exit(1)
	}
}

func GetContentTypeFromFile(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the file offset back to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	// Detect the content type based on the file's magic number
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
