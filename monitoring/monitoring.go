package monitoring

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"study-golang-alura/logs"
	"study-golang-alura/types"
	"time"
)

func StartMonitoring() {
	println("Starting monitoring")
	// set sites to monitor with pointers to their status
	var sites []types.SiteToMonitor = getSitesFromFiles()
	// if sites is empty then exit
	if len(sites) == 0 {
		println("No sites to monitor")
		return
	}

	// loop five times
	for times := 0; times < getMonitoringQuantity(); times++ {
		// for each site, start monitoring
		for index := 0; index < len(sites); index++ {
			site := sites[index]
			// request site
			res, error := http.Get(site.Url)
			if error != nil {
				println("Error: " + error.Error())
				sites[index].Status = -1
			} else {
				println("Status: " + res.Status)
				sites[index].Status = res.StatusCode
			}
		}
		time.Sleep(time.Second * time.Duration(getDelay()))
	}

	fmt.Println(sites)
	logs.RegisterMonitoringLogs(sites)
}

func getSitesFromFiles() []types.SiteToMonitor {
	var sites []types.SiteToMonitor
	// find files with prefix "site"
	// for each file, read it and add to sites
	dirEntries, error := os.ReadDir(".")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	for index := 0; index < len(dirEntries); index++ {
		filename := dirEntries[index].Name()
		isSitesList := strings.HasPrefix(filename, "site")
		// if file is a sites list
		if isSitesList {
			// open file
			file, error := os.Open(filename)
			if error != nil {
				fmt.Println(error)
				os.Exit(1)
			}
			defer file.Close()
			// read file line by line
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				site := scanner.Text()
				sites = append(sites, types.SiteToMonitor{Url: site, Status: 0})
			}
		}
	}
	return sites
}
