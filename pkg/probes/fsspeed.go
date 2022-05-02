package probes

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/iop"
	log "github.com/Daniel-WWU-IT/cs3probes/pkg/logging"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/tests"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/timing"
)

func RunFSSpeedProbe(target string, user, pass string, warnLimit int, percentile int) int {

	// Check if required flags are set

	if target == "" || user == "" || pass == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
		flag.PrintDefaults()
		os.Exit(checkError)
	}

	// Setup Logger and Log object

	data, logger := log.CreateSystemLog(target, warnLimit)

	// start API-Session

	session, err := iop.CreateSession(target, user, pass)
	if err != nil {
		fmt.Printf("Session failed\n")
		return checkError
	}

	// Test to upload 10 small 10kb files

	res, err := timing.TimeIopFunction(tests.Test_sUpload, session)

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

	// Insert Data into Database and get outliers in return

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
