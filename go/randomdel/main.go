package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func randomDelete(fileList []string) []string {
	/* Randomly delete a random number of files from the fileList,
	 * Return the names of deleted files
	 */
	n := len(fileList)
	numFilesToDelete := rand.Intn(n)
	fmt.Println("Randomly deleting", numFilesToDelete, "files")

	filesToDelete := make(map[int]bool)
	for len(filesToDelete) < numFilesToDelete {
		number := rand.Intn(n)
		if _, ok := filesToDelete[number]; !ok {
			filesToDelete[number] = true
		}
	}

	var filesDeleted []string
	for index, shouldDelete := range filesToDelete {
		if !shouldDelete {
			continue
		}
		file := fileList[index]
		os.Remove(file)
		filesDeleted = append(filesDeleted, file)
	}
	return filesDeleted
}

func getFiles(baseDir string) []string {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}

	var fileList []string
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", baseDir, file.Name())
		if file.IsDir() {
			fileList = append(fileList, getFiles(filePath)...)
			continue
		}

		fullPath, err := filepath.Abs(filePath)
		if err != nil {
			panic(err)
		}

		fileList = append(fileList, fullPath)
	}
	return fileList
}

func main() {
	userInput := flag.String("path", ".", "The top-level directory to run this on.")
	flag.Parse()

	files := getFiles(*userInput)
	deletedFiles := randomDelete(files)
	for _, file := range deletedFiles {
		fmt.Println("Deleted:", file)
	}
}
