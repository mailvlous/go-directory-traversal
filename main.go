package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ffuf",
		"-u", "http://testphp.vulnweb.com/FUZZ",
		"-w", "SecLists/Discovery/Web-Content/common.txt",
		"-mc", "200",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed ..:", err)
	}
}
