package util

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"os/exec"
	"time"
)

func ShowCommandOutputWithSpinner(cmd *exec.Cmd, loadingMessage string) {
	// Start Spinner
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Prefix = loadingMessage + " "
	s.Start()

	// Get the command's output pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe: %v\n", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating StderrPipe: %v\n", err)
		return
	}

	// Start the command
	if err = cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	// Create a scanner to read the command's output
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			s.Suffix = " " + scanner.Text()
		}
	}()

	// Read stderr as well
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			s.Suffix = " " + errScanner.Text()
		}
	}()

	// Wait for the command to finish
	cmd.Wait()
	s.Stop()
}

func ShowCommandOutput(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe: %v\n", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating StderrPipe: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()

	cmd.Wait()
}
