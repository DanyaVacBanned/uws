package scripts

import (
	"bufio"
	"fmt"
	"os"
)

func CreateScript(name string) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	fmt.Println("Enter shell commands.")
	fmt.Println("Type 'done' to finish.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "done" {
			break
		}
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	scriptName := fmt.Sprintf("%s.sh", name)
	file, err := os.Create(scriptName)

	if err != nil {
		fmt.Printf("Error creating script file: %v\n", err)
		return
	}
	err = os.Chmod(scriptName, 0755)
	if err != nil {
		fmt.Printf("Error setting script permissions: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n\n")

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to script file: %v\n", err)
			return
		}
	}
	fmt.Printf("Script created: %s.sh\n", name)
}
