package main

/*
Usage example: >$ ./fsoperations -target=reva:443 -user=username -pass=password -warnlimit=1000 -percentile=95
*/

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/Daniel-WWU-IT/cs3probes/pkg/logging"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/tests"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/timing"
	iop "github.com/cs3org/reva/pkg/sdk"
)

// Setup variables

var (
	target, user, pass         string
	warnLimit, percentile, res int
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

	if target == "" || user == "" || pass == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
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

	// Start API Session

	session := iop.MustNewSession()
	session.Initiate(target, false)
	err := session.BasicLogin(user, pass)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return checkError
	}

	// Time and test file enumeration

	res, err = timing.TimeIopFunction(tests.Test_ls, session)

	if err != nil {
		fmt.Printf("Test_ls failed\n")
		return checkError
	}

	data.AddMetric("ls", res)

	// Time and test to create a directory

	res, err = timing.TimeIopFunction(tests.Test_mkdir, session)

	if err != nil {
		fmt.Printf("Test_mkdir failed\n")
		return checkError
	}

	data.AddMetric("mkdir", res)

	// Time and test operation "directory exists"

	res, err = timing.TimeIopFunction(tests.Test_direxists, session)

	if err != nil {
		fmt.Println("Test_direxists failed\n")
		return checkError
	}

	data.AddMetric("direxists", res)

	// Time and test to remova a directory

	res, err = timing.TimeIopFunction(tests.Test_rmdir, session)

	if err != nil {
		fmt.Println("Test_rmdir failed\n")
		return checkError
	}

	data.AddMetric("rmdir", res)

	// Time and Test to upload a file

	res, err = timing.TimeIopFunction(tests.Test_upload, session)

	if err != nil {
		fmt.Println("Test_upload failed\n")
		return checkError
	}

	data.AddMetric("upload", res)

	// Time and Test operation "file exists"

	res, err = timing.TimeIopFunction(tests.Test_fileexists, session)

	if err != nil {
		fmt.Println("Test_fileexists failed\n")
		return checkError
	}

	data.AddMetric("fileexists", res)

	// Time and test to move a file

	res, err = timing.TimeIopFunction(tests.Test_mvfile, session)
	if err != nil {
		fmt.Println("Test_mvfile failed\n")
		return checkError
	}

	data.AddMetric("mvfile", res)

	// Time and test to remove a file

	res, err = timing.TimeIopFunction(tests.Test_rmfile, session)

	if err != nil {
		fmt.Println("Test_rmfile failed\n")
		return checkError
	}
	data.AddMetric("rmfile", res)

	// Insert Data into database and get outliers in return

	outliers := logger.InsertLog(data, percentile)

	// Return checkOK, if there are no outliers

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
