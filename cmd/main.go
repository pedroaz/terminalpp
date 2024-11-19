package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("START")

	// Channel to signal the goroutine to stop
	stopChan := make(chan struct{})

	// Launch the command in a goroutine
	go func() {

		fmt.Println("Starting command...")
		textCommand := "while [ 1 ] ; do echo \"$(date +%T)\" ; sleep 1 ; done"
		textCommand = "echo 'Hello World'"
		cmd := exec.Command("sh", "-c", textCommand)
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error obtaining stdout pipe:", err)
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			select {
			case <-stopChan:
				if err := cmd.Process.Kill(); err != nil {
					fmt.Println("Error killing process:", err)
				}
				return
			default:
				fmt.Println(scanner.Text())
			}
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("Error waiting for command:", err)
		}
		fmt.Println("Cmd finished")
	}()

	// Set up channels to listen for interrupt signal and 'K' key press
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	keyChan := make(chan struct{})

	// Goroutine to listen for 'K' key press
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadByte()
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			if input == 'K' || input == 'k' {
				keyChan <- struct{}{}
				return
			}
		}
	}()

	fmt.Println("Press 'K' to stop the program.")
	select {
	case <-signalChan:
	case <-keyChan:
	}
	close(stopChan) // Signal the goroutine to stop
	fmt.Println("\nEND")
}
