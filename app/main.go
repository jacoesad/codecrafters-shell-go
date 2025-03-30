package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var systemPaths []string

func init() {
	initPaths()
}

func initPaths() {
	systemPaths = strings.Split(os.Getenv("PATH"), ":")
}

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
				found, fullPath := findCmd(args[1])
				if found {
					fmt.Println(args[1] + " is " + fullPath)
				} else {
					fmt.Println(args[1] + ": not found")
				}
			}
		default:
			execCmd(args[0], args[1:])
		}
	}
}

func findCmd(cmd string) (bool, string) {
	for _, path := range systemPaths {
		fullPath := filepath.Join(path, cmd)
		if _, err := os.Stat(fullPath); err == nil {
			return true, fullPath
		}
	}
	return false, ""
}

func execCmd(cmd string, args []string) {
	found, _ := findCmd(cmd)
	if !found {
		fmt.Println(cmd + ": command not found")
		return
	}

	execCmd := exec.Command(cmd, args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Run()
}
