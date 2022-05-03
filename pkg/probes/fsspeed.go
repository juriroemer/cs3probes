package probes

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/iop"
	log "github.com/Daniel-WWU-IT/cs3probes/pkg/logging"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/nagios"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/tests"
)

func RunFSSpeedProbe(target string, user, pass string, warnLimit int, percentile int) int {
	// Check if required flags are set
	if target == "" || user == "" || pass == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
		flag.PrintDefaults()
		os.Exit(nagios.CheckError)
	}

	// Setup Logger and Log object
	data, logger := log.CreateSystemLog(target, warnLimit)

	// Start API-Session
	session, err := iop.CreateSession(target, user, pass)
	if err != nil {
		fmt.Printf("Session creation failed: %v\n", err)
		return nagios.CheckError
	}

	// Create testing context
	ctx, err := tests.NewTestContext(session)
	if err != nil {
		fmt.Printf("Test context creation failed: %v\n", err)
		return nagios.CheckError
	}

	ctx.BeginTests()

	// Test to upload 10 small 10kb files
	res := ctx.RunIOPTest(tests.Test_sUpload, "Upload small files")
	data.AddMetric("sUpload", res)

	// Test to upload 1 bigger 100kb file
	res = ctx.RunIOPTest(tests.Test_bUpload, "Upload big file")
	data.AddMetric("bUpload", res)

	// Test to move 10 small 10kb files
	res = ctx.RunIOPTest(tests.Test_sMove, "Move small files")
	data.AddMetric("sMove", res)

	// Test to move 1 bigger 100kb file
	res = ctx.RunIOPTest(tests.Test_bMove, "Move big file")
	data.AddMetric("bMove", res)

	// Test to remove 10 small 10kb files
	res = ctx.RunIOPTest(tests.Test_sRemove, "Remove small files")
	data.AddMetric("sRemove", res)

	// Test to remove 1 bigger 100kb file
	res = ctx.RunIOPTest(tests.Test_bRemove, "Remove big file")
	data.AddMetric("bRemove", res)

	// Insert Data into Database and get outliers in return
	outliers := logger.InsertLog(data, percentile)

	ctx.EndTests(outliers)

	// Return CheckOK if there are no outliers
	if outliers != nil {
		return nagios.CheckWarning
	}
	return nagios.CheckOK
}
