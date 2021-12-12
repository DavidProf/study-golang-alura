package monitoring

import (
	"os"
	"strconv"
)

func getMonitoringQuantity() int {
	quantity, error := strconv.Atoi(os.Getenv("MONITORING_QUANTITY"))
	if error != nil {
		return 3
	}
	return quantity
}

func getDelay() int {
	delay, error := strconv.Atoi(os.Getenv("DELAY"))
	if error != nil {
		return 1
	}
	return delay
}
