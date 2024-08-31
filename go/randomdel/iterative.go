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
	for index, _ := range filesToDelete {
		file := fileList[index]
		os.Remove(file)
		filesDeleted = append(filesDeleted, file)
	}
	return filesDeleted
}

func getFiles(baseDir string) []string {
	baseDirs := []string{baseDir}
	var fileList []string

	for len(baseDirs) > 0 {
		baseDir := baseDirs[0]
        baseDirs = baseDirs[1:]
		files, err := os.ReadDir(baseDir)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			dirPath := fmt.Sprintf("%s/%s", baseDir, file.Name())
			if file.IsDir() {
				baseDirs = append(baseDirs, dirPath)
				continue
			}
			fullPath, err := filepath.Abs(dirPath)
			if err != nil {
				panic(err)
			}
			fileList = append(fileList, fullPath)
		}
	}

	return fileList
}

func main() {
	userInput := flag.String("path", ".", "The top-level directory to run this on.")
	flag.Parse()

	files := getFiles(*userInput)
	fmt.Println(files)
	deletedFiles := randomDelete(files)
	
	for _, file := range deletedFiles {
        fmt.Println("Deleted:", file)
	}
}
