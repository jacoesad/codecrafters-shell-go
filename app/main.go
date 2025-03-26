package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		command := strings.TrimSpace(input[:len(input)-1])

		args := strings.Split(command, " ")
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "exit":
			os.Exit(0)
		case "type":
			switch args[1] {
			case "echo", "exit", "type":
				fmt.Println(args[1] + " is a shell builtin")
			default:
				fmt.Println(args[1] + ": not found")
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
