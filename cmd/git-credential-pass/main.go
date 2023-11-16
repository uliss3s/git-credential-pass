package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	operation := os.Args[1]

	// Read input from Git about the credential operation
	input := readInput()

	// Parse the input into a map
	inputMap := parseInput(input)

	// Check the operation from the args (only interested in "get" operation)
	if operation == "get" {
		// Retrieve the password using Pass
		username := inputMap["username"]
		host := inputMap["host"]
		password := getPasswordFromPass(host, username)

		// Output the password to Git
		outputPassword(username, password)
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	return input
}

func parseInput(input string) map[string]string {
	lines := strings.Split(input, "\n")
	inputMap := make(map[string]string)

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			inputMap[key] = value
		}
	}

	return inputMap
}

func getPasswordFromPass(host string, username string) string {
	// Run 'pass' command to retrieve the password
	cmd := exec.Command("pass", "show", host+"/"+username)
	passwordBytes, err := cmd.Output()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving password: %v\n", err)
		os.Exit(1)
	}

	return string(passwordBytes)
}

func outputPassword(username string, password string) {
	fmt.Printf("username=%s\npassword=%s\n", username, password)
}
