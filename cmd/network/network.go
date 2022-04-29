package main

/*
Usage example: >$ ./network -target=reva:443
*/

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/juriroemer/cs3probes/pkg/logging"
	tests "github.com/juriroemer/cs3probes/pkg/tests"
)

// Setup variables

var (
	target 				string
	insecure           bool
	warnLimit, percentile		int

)

// Setup return values

const (
	checkOK      = 0
	checkWarning = 1
	checkError   = 2
	checkUnknown = 3
)



func init() {

	// Setup commandline flags

	flag.StringVar(&target, "target", "", "[required] the target, [host]:[port]")
	flag.IntVar(&warnLimit, "warnlimit", 100, "minimum number of logs for outlier detection")
	flag.IntVar(&percentile, "percentile", 90, "the percentile for outlier detection")
	flag.Parse()
}


func main() {
	os.Exit(run())
}

func run() int {

	// Check if required flags are set
	if (target == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "") {
		fmt.Println("Please specify target")
		flag.PrintDefaults()
		os.Exit(checkError)
	}

	// Setup Logger and Log object
	data := log.NewLog()
	logger := log.NewLogger()
	path := strings.Split(os.Args[0], "/")
	data.SetProbeName(path[len(path)-1])
	data.SetWarnLimit(warnLimit)
	data.SetTarget(target)

	// Ping target system
	// the ping library has it's own timing capabilities, those are used instead of Time.TargetFunction,

	time, state := tests.Test_ping(data.Host())
	if (state == checkOK){
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

	//Insert Data into Database and get Outliers in return

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

