package main

import (
	"fmt"
	"os"
	"strconv"

	"study-golang-alura/logs"
	"study-golang-alura/monitoring"

	"github.com/joho/godotenv"
)

func main() {
	command := getCommandFromArgv()
	loadEnv()
	switch command {
	case 0:
		monitoring.StartMonitoring()
	case 1:
		logs.PrintLogs()
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func loadEnv() {
	// load environment variables
	// if error then exit
	error := godotenv.Load()
	if error != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}

func getCommandFromArgv() int {
	if len(os.Args) < 2 {
		return 0
	}
	command, error := strconv.Atoi(os.Args[1])
	if error != nil {
		return 0
	}
	return command
}
