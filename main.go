package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

func banner(blue string, reset string) {
	fmt.Printf(blue)
	banner := `                                             
      _ )  |              _ \ _)  |  | _ _|      _)             |              
      _ \  |  |  |   -_)  __/  |  |  |   |     \  |   -_)   _|   _|   _ \   _| 
     ___/ _| \_,_| \___| _|   _| _| _| ___| _| _| | \___| \__| \__| \___/ _|   
                                                __/                             
  
	`
	fmt.Println(banner)
	fmt.Printf(reset)
}

func main() {
	// sudo your program
	if os.Getuid() != 0 {
		cmd := exec.Command("sudo", os.Args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		return
	}

	blue := "\033[34m"
	reset := "\033[0m"

	banner(blue, reset)
	fmt.Printf("[ ! ] About to rewrite files with the hex message and erase the file. Continue? (y/N): ")
	var response string
	fmt.Scanln(&response)
	if response != "y" && response != "Y" {
		os.Exit(0)
	}

	//os.Exit(0) // BUG!!!!!

	// DONT USE THIS CODE IF YOU ARE NOT SURE!!!
	// rewrite file and erase with loading bar.
	files := []string{}
	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	totalFiles := len(files)
	// create and start new bar
	bar := pb.StartNew(totalFiles)
	// bar will format numbers as bytes (B, KiB, MiB, etc)
	bar.Set(pb.Bytes, true)
	// start bar from 'default' template
	// bar := pb.Default.Start(totalFiles)

	// start bar from 'simple' template
	// bar := pb.Simple.Start(totalFiles)

	// start bar from 'full' template
	//bar := pb.Full.Start(totalFiles)

	for _, file := range files {
		bar.Increment()
		rewriteFileWithHexMessage(file)
		os.Remove(file)
	}
	// finish bar
	bar.Finish()
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
	fmt.Println("                  ---: YOU HAVE GIVEN THE BLUE PILL :---")
	fmt.Printf(blue)
	fmt.Println(`
                                   .-.
                                  /:::\
                                 /::::/
                                / '-:/
                               /    /
                               \   /
                                '"'

	`)
	fmt.Printf(reset)
}

func rewriteFileWithHexMessage(filePath string) {
	// Step 1: Open the file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Step 2: Read all bytes from the file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 3: Create the hex encoded message
	message := "You will only find the blue pills ..."
	hexMessage := hex.EncodeToString([]byte(message))

	// Step 4: Write the hex encoded message to all bytes in the file
	for i := range fileBytes {
		fileBytes[i] = hexMessage[i%len(hexMessage)]
	}

	// Step 5: Move file pointer to the beginning of the file
	file.Seek(0, 0)

	// Step 6: Write the modified bytes back to the file
	_, err = file.Write(fileBytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Step 7: Truncate the file to the new length
	err = file.Truncate(int64(len(fileBytes)))
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}
}
