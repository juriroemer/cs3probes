package main

/*
Usage example: >$ ./network -target=reva:443
*/

import (
	"flag"
	"os"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/probes"
)

var (
	target                string
	insecure              bool
	warnLimit, percentile int
)

func init() {
	flag.StringVar(&target, "target", "", "[required] the target, [host]:[port]")
	flag.IntVar(&warnLimit, "warnlimit", 100, "minimum number of logs for outlier detection")
	flag.IntVar(&percentile, "percentile", 90, "the percentile for outlier detection")
	flag.Parse()
}

func main() {
	os.Exit(probes.RunNetworkProbe(target, warnLimit, percentile))
}
