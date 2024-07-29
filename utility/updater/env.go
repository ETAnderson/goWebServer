package updater

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// UpdateExampleEnv updates the example.env file with the type of each environment variable
func UpdateExampleEnv() {
	// Open the .env file
	envFile, err := os.Open(".env")
	if err != nil {
		log.Println("Error opening .env file:", err)
		return
	}
	defer envFile.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				envVars[parts[0]] = parts[1]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading .env file:", err)
		return
	}

	// Create or open the example.env file
	exampleEnvFile, err := os.Create("example.env")
	if err != nil {
		log.Println("Error creating example.env file:", err)
		return
	}
	defer exampleEnvFile.Close()

	// Write the variable names and types to the example.env file
	for key := range envVars {
		varType := "string" // Default type
		if _, err := strconv.Atoi(envVars[key]); err == nil {
			varType = "int"
		} else if _, err := strconv.ParseBool(envVars[key]); err == nil {
			varType = "bool"
		}
		exampleEnvFile.WriteString(fmt.Sprintf("%s=%s\n", key, varType))
	}
}
