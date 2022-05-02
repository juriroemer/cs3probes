package probes

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/Daniel-WWU-IT/cs3probes/pkg/logging"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/tests"
)

func RunNetworkProbe(target string, warnLimit int, percentile int) int {
	// Check if required flags are set
	if target == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
		fmt.Println("Please specify target")
		flag.PrintDefaults()
		os.Exit(checkError)
	}

	// Setup Logger and Log object

	data, logger := log.CreateSystemLog(target, warnLimit)

	// Ping target system
	// the ping library has it's own timing capabilities, those are used instead of Time.TargetFunction,

	time, state := tests.Test_ping(data.Host())
	if state == checkOK {
		data.AddMetric("ping", time)
	} else {
		fmt.Printf("Test_Ping failed\n")
		return state
	}

	// Test if port is available

	response := tests.Test_portscan(target)

	if response == checkError {
		fmt.Printf("Test_portscan failed\n")
		return checkError
	}

	// parse response code to {0,1}-boolean and add to log
	// checkOK -> 1 (true)
	// checkError -> 0 (false)

	data.AddMetric("portscan", (tests.Test_portscan(target)+1)%2)

	// Insert Data into Database and get Outliers in return

	outliers := logger.InsertLog(data, percentile)

	// Return checkOK, if there are no Outliers

	if outliers == nil {
		fmt.Printf("Probe %s ended successfully\n", data.Probe())
		return checkOK
	}

	// Else, write warnings for each outlier to stdout and return checkWarning

	fmt.Printf("Probe %s ended with %d Warnings\n", data.Probe(), len(outliers))
	for test, time := range outliers {
		fmt.Printf("WARNING: Test %s took %d ms\n", test, time)
	}

	return checkWarning
}
