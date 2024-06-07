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
	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(0)
	}

	return strings.TrimSpace(cmd);	
}

func executeCommand(command string) {
	if strings.HasPrefix(command, "echo") {
		fmt.Println(strings.TrimPrefix(command, "echo "));
	} else if strings.HasPrefix(command, "exit") {
		os.Exit(0);
	} else {
		fmt.Printf("%s: command not found\n", command);
	}
}	
