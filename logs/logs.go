package logs

import (
	"bufio"
	"os"
	"strconv"
	"study-golang-alura/types"
	"time"
)

const filename = "monitoring.logs"

func PrintLogs() {
	// open file
	file, error := os.Open(filename)
	if error != nil {
		println("Error: " + error.Error())
		os.Exit(1)
	}
	defer file.Close()
	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func RegisterMonitoringLogs(sites []types.SiteToMonitor) {
	// append to logs file, date, site and site status
	// if error then exit
	file, error := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if error != nil {
		println("Error: " + error.Error())
		os.Exit(1)
	}
	defer file.Close()
	// write to file
	for index := 0; index < len(sites); index++ {
		site := sites[index]
		_, error = file.WriteString(time.Now().Format(time.RFC3339) + " " + site.Url + " " + strconv.Itoa(site.Status) + "\n")
		if error != nil {
			println("Error: " + error.Error())
			os.Exit(1)
		}
	}
}

// export printLogs
