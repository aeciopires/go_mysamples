package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Filename struct {
	name        string
	permissions []string
}

// List permissions of file
// Reference: https://stackoverflow.com/questions/45429210/how-do-i-check-a-files-permissions-in-linux-using-go
func listFilePermission(file string) (result Filename) {
	// Creating slice with 9 characters
	permission := make([]string, 9)

	// Get permissions of file if format 'rwxrwxrwx' nine characters
	info, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	mode := info.Mode()

	// Get permissions for owner user, first group of three characters
	for i := 1; i < 4; i++ {
		permission = append(permission, string(mode.String()[i]))
	}

	// Get permissions for group of user of owner, second group of three characters
	for i := 4; i < 7; i++ {
		permission = append(permission, string(mode.String()[i]))
	}

	// Get permissions for others users, third group of three characters
	for i := 7; i < 10; i++ {
		permission = append(permission, string(mode.String()[i]))
	}

	result.name = file
	result.permissions = permission

	return
}

// Check if directory exist, list subdirectory
// Reference: https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html (Option 2: filepath.Walk)
func listSubDirectories(directory string, channelAux chan Filename) {
	// Closing channel to evict deadlock
	defer close(channelAux)

	// filepath.Walk list files in a directory structure, from the filepath Go package.
	// It also allows us to recursively discover directories and files.
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// If info.IsDir() equals false, then path is file
		// If info.IsDir() equals true, then path is directory
		if info.IsDir() == false {
			// Send data to channel
			channelAux <- listFilePermission(path)
		} else {
			log.Printf(
				"[INFO] %s is a directory",
				path,
			)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

// Reference: https://stackoverflow.com/questions/17825857/how-to-make-a-channel-that-receive-multiple-return-values-from-a-goroutine
func main() {
	// Creating a channel with buffer with 5 elements
	mychannel := make(chan Filename, 5)

	path := "/home/aeciopires/Desktop"

	// Creating a go routine
	go listSubDirectories(path, mychannel)

	// Getting data of channel with range statement
	for results := range mychannel {
		log.Printf(
			"[OK] %s has permissions %s",
			results.name, results.permissions,
		)
	}

	fmt.Println("End!")
}
