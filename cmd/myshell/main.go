package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cmds := setup()

	for {
		command := readCommand()
		executeCommand(command, cmds)
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

func executeCommand(command string, commands Commands) {
	cmd := strings.Split(command, " ")[0]
	if _, ok := commands[cmd]; ok {
		commands[cmd].(func(string))(command)
	} else {
		fmt.Printf("%s: command not found\n", cmd);
	}
}	