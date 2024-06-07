package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		command := readCommand()
		executeCommand(command)
	}
}

func readCommand() string {
	fmt.Fprint(os.Stdout, "$ ")
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(0)
	}

	return strings.Trim(reader, "\n");	
}

func executeCommand(command string) {
	fmt.Printf("%s: command not found\n", command);
}	
