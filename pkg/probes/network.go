package probes

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/Daniel-WWU-IT/cs3probes/pkg/logging"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/nagios"
	"github.com/Daniel-WWU-IT/cs3probes/pkg/tests"
)

func RunNetworkProbe(target string, warnLimit int, percentile int) int {
	// Check if required flags are set
	if target == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
		fmt.Println("Please specify target")
		flag.PrintDefaults()
		os.Exit(nagios.CheckError)
	}

	// Setup Logger and Log object
	data, logger := log.CreateSystemLog(target, warnLimit)

	// Create testing context
	ctx, err := tests.NewTestContext(nil)
	if err != nil {
		fmt.Printf("Test context creation failed: %v\n", err)
		return nagios.CheckError
	}

	ctx.BeginTests()

	// Ping target system
	// the ping library has it's own timing capabilities, those are used instead of Time.TargetFunction,
	_, time := ctx.RunNetworkTest(tests.Test_ping, data.Host(), "Ping")
	data.AddMetric("ping", time)

	// Test if port is available
	res, _ := ctx.RunNetworkTest(tests.Test_portscan, target, "Portscan")
	data.AddMetric("portscan", (res+1)%2)

	// Insert Data into Database and get Outliers in return
	outliers := logger.InsertLog(data, percentile)

	ctx.EndTests(outliers)

	// Return CheckOK if there are no outliers
	if outliers != nil {
		return nagios.CheckWarning
	}
	return nagios.CheckOK
}
