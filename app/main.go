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

		if args[0] == "exit" {
			os.Exit(0)
		} else {
			fmt.Println(input[:len(input)-1] + ": command not found")
		}
	}
}
