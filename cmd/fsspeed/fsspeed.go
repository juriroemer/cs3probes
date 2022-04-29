package main

/*
Usage example: >$ ./fsspeed -target=reva:443 -user=username -pass=password -warnlimit=1000 -percentile=95
*/

import (
	"flag"
	"fmt"
	"os"
	"strings"

	iop "github.com/cs3org/reva/pkg/sdk"
	log "github.com/juriroemer/cs3probes/pkg/logging"
	timing "github.com/juriroemer/cs3probes/pkg/timing"
	tests "github.com/juriroemer/cs3probes/pkg/tests"
)

// Setup variables

var (
	target, user, pass string
	warnLimit, percentile, res		int
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

	flag.StringVar(&target, "target", "", "[required] the target iop")
	flag.StringVar(&user, "user", "", "[required] the username")
	flag.StringVar(&pass, "pass", "", "[required] the user password")
	flag.IntVar(&warnLimit, "warnlimit", 100, "minimum number of logs for outlier detection")
	flag.IntVar(&percentile, "percentile", 90, "the percentile for outlier detection")
	flag.Parse()
}


func main() {
	os.Exit(run())
}

func run() int {

	// Check if required flags are set

	if (target == "" || user == "" || pass == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "") {
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

	// start API-Session

	session := iop.MustNewSession()
	session.Initiate(target, false)
	err := session.BasicLogin(user, pass)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return checkError
	}

	// Test to upload 10 small 10kb files

	res, err = timing.TimeIopFunction(tests.Test_sUpload, session)

	if err != nil {
		fmt.Printf("Test_sUpload failed\n")
		return checkError
	}

	data.AddMetric("sUpload", res)


	// Test to upload 1 bigger 100kb file

	res, err = timing.TimeIopFunction(tests.Test_bUpload, session)

	if err != nil {
		fmt.Printf("Test_bUpload failed\n")
		return checkError
	}
	data.AddMetric("bUpload", res)

	// Test to move 10 small 10kb files

	res, err = timing.TimeIopFunction(tests.Test_sMove, session)

	if err != nil {
		fmt.Printf("Test_sMove failed\n")
		return checkError
	}

	data.AddMetric("sMove", res)

	// Test to move 1 bigger 100kb file

	res, err = timing.TimeIopFunction(tests.Test_bMove, session)

	if err != nil {
		fmt.Printf("Test_bMove failed\n")
		return checkError
	}
	data.AddMetric("bMove", res)

	// Test to remove 10 small 10kb files

	res, err = timing.TimeIopFunction(tests.Test_sRemove, session)

	if err != nil {
		fmt.Printf("Test_sRemove failed\n")
		return checkError
	}

	data.AddMetric("sRemove", res)

	// Test to remove 1 bigger 100kb file

	res, err = timing.TimeIopFunction(tests.Test_bRemove, session)

	if err != nil {
		fmt.Printf("Test_bRemove failed\n")
		return checkError
	}

	data.AddMetric("bRemove", res)


	//Insert Data into Database and get outliers in return

	outliers := logger.InsertLog(data, percentile)

	// Return checkOK, if there are no outliers

	if outliers == nil {
		fmt.Printf("Probe %s ended successfully\n", data.Probe())
		return checkOK
	}

	// Else, write warnings for each outlier to stdout and return checkWarn
	fmt.Printf("Probe %s ended with %d Warnings\n", data.Probe(), len(outliers))
	for test, time := range outliers {
		fmt.Printf("WARNING: Test %s took %d ms\n", test, time)
	}

	return checkWarning
}
