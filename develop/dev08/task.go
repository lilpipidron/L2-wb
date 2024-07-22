package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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
	case "kill":
		if len(args) != 1 {
			fmt.Println("kill: invalid number of parameters")
			return
		}
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("kill: %v\n", err)
			return
		}
		err = syscall.Kill(pid, syscall.SIGKILL)
		if err != nil {
			fmt.Printf("kill: %v\n", err)
		}
	case "ps":
		cmd := exec.Command("ps", "-e")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("ps: %v\n", err)
		}
	default:
		externalCmd := exec.Command(cmd, args...)
		externalCmd.Stdout = os.Stdout
		externalCmd.Stderr = os.Stderr
		err := externalCmd.Run()
		if err != nil {
			fmt.Printf("%s: %v\n", cmd, err)
		}
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
