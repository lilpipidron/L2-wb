package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Используем библиотеку flag для парсинга timeout
// Библиотеку os/signal для завершения по ctrl+D
// net для подключения к серверу

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatal("incorrect number of arguments")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		select {
		case <-sigChan:
			fmt.Println("Received signal, exiting...")
			os.Exit(0)
		}
	}()

	conn, err := net.DialTimeout("tcp", args[0]+":"+args[1], *timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}

		if _, err := fmt.Fprintf(conn, text); err != nil {
			fmt.Println("Error writing to connection:", err)
			break
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by server")
			break
		}
		fmt.Print("Server response: " + message)
	}

	fmt.Println("Exiting program")
}
