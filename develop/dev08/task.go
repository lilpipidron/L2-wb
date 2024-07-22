package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var currDir string

func changeCurrDir() {
	pathToCurrDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirs := strings.Split(pathToCurrDir, "/")
	currDir = dirs[len(dirs)-1]
}

func runPipeline(command string) {}

func runCommand(command string) {
	parts := strings.Fields(command)
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "cd":
		if len(args) > 0 {
			if err := os.Chdir(args[0]); err != nil {
				fmt.Printf("cd: %v\n", err)
				return
			}
			changeCurrDir()
		} else {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("cd: %v\n", err)
				return
			}
			if err := os.Chdir(homeDir); err != nil {
				fmt.Printf("cd: %v\n", err)
				return
			}
			changeCurrDir()
		}
	case "pwd":
		pathToCurrDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("pwd: %v\n", err)
			return
		}
		fmt.Println(pathToCurrDir)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type \\quit to exit")
	changeCurrDir()
	for {
		fmt.Printf("GoShell> %s $", currDir)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		command = strings.TrimSpace(command)
		if command == "\\quit" {
			break
		}

		if command == "" {
			continue
		}

		if strings.Contains(command, "|") {
			runPipeline(command)
		} else {
			runCommand(command)
		}
	}
}
