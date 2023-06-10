package utils

import (
	"fmt"
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
