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

func RunFSOperationsProbe(target string, user, pass string, warnLimit int, percentile int) int {
	// Check if required flags are set
	if target == "" || user == "" || pass == "" || len(strings.Split(target, ":")) != 2 || strings.Split(target, ":")[0] == "" || strings.Split(target, ":")[1] == "" {
		flag.PrintDefaults()
		os.Exit(nagios.CheckError)
	}

	// Setup Logger and Log object
	data, logger := log.CreateSystemLog(target, warnLimit)

	// Start API Session
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

	// Base directory of all tests
	root := "/home/fsoperations/"

	// Initialize testing
	tests.InitializeTests(session, root)

	// Time and test file enumeration
	res := ctx.RunIOPTest(tests.Test_ls, root, "List files")
	data.AddMetric("ls", res)

	// Time and test to create a directory
	res = ctx.RunIOPTest(tests.Test_mkdir, root, "Make directory")
	data.AddMetric("mkdir", res)

	// Time and test operation "directory exists"
	res = ctx.RunIOPTest(tests.Test_direxists, root, "Directory exists")
	data.AddMetric("direxists", res)

	// Time and test to remova a directory
	res = ctx.RunIOPTest(tests.Test_rmdir, root, "Remove directory")
	data.AddMetric("rmdir", res)

	// Time and Test to upload a file
	res = ctx.RunIOPTest(tests.Test_upload, root, "Upload file")
	data.AddMetric("upload", res)

	// Time and Test operation "file exists"
	res = ctx.RunIOPTest(tests.Test_fileexists, root, "File exists")
	data.AddMetric("fileexists", res)

	// Time and Test to download a file
	res = ctx.RunIOPTest(tests.Test_download, root, "Download file")
	data.AddMetric("download", res)

	// Time and test to move a file
	res = ctx.RunIOPTest(tests.Test_mvfile, root, "Move files")
	data.AddMetric("mvfile", res)

	// Time and test to remove a file
	res = ctx.RunIOPTest(tests.Test_rmfile, root, "Remove files")
	data.AddMetric("rmfile", res)

	// Insert Data into database and get outliers in return
	outliers := logger.InsertLog(data, percentile)

	ctx.EndTests(outliers)

	// Return CheckOK if there are no outliers
	if outliers != nil {
		return nagios.CheckWarning
	}
	return nagios.CheckOK
}
