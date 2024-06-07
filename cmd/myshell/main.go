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

type CommandFunc interface{}

type Commands map[string]CommandFunc

func (c Commands) add(name string, function CommandFunc) {
    c[name] = function
}

// Setup initializes and returns the Commands map with predefined commands
func setup() Commands {
	commands := Commands{}
	commands.add("exit", func(in string) {
		os.Exit(0)
	})
	commands.add("echo", func(in string) {
		fmt.Println(strings.TrimPrefix(in, "echo "))
	})
	commands.add("type", func(in string) {
		cmd := strings.TrimPrefix(in, "type ")
		if _, ok := commands[cmd]; ok {
			fmt.Printf("%s is a shell builtin\n", cmd)
		} else {
			fmt.Printf("%s not found\n", cmd)
		}
	})
	return commands
}
