package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ffuf",
		"-u", "https://example.com/FUZZ",
		"-w", "wordlist.txt",
		"-mc", "200",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed ..:", err)
	}
}
