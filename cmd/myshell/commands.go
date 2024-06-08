package main

import (
	"fmt"
	"os"
	"strings"
)

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
			return
		}

		folders := strings.Split(os.Getenv("PATH"), ":")
		for _, folder := range folders {
			if _, err := os.Stat(folder + "/" + cmd); err == nil {
				fmt.Printf("%s is %s/%s\n", cmd, folder, cmd)
				return
			}
		}
			
		fmt.Printf("%s: command not found", cmd)
	})
	return commands
}
