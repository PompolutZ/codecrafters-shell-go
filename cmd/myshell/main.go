package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(0)
	}

	command := strings.Trim(reader, "\r\n");

	fmt.Printf("%s: command not found\r\n", command);
}
