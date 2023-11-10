package fileoperations

import (
	"fmt"
	"os"
)

func CreateFilePaths(directories []string) {
	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
			} else {
				fmt.Printf("Directory %s created successfully.\n", dir)
			}
		} else if err != nil {
			fmt.Printf("Error checking directory %s: %v\n", dir, err)
		}
	}
}
