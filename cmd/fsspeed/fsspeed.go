package main

/*
Usage example: >$ ./fsspeed -target=reva:443 -user=username -pass=password -warnlimit=1000 -percentile=95
*/

import (
	"flag"
	"os"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/probes"
)

var (
	target, user, pass         string
	warnLimit, percentile, res int
)

func init() {
	flag.StringVar(&target, "target", "", "[required] the target iop")
	flag.StringVar(&user, "user", "", "[required] the username")
	flag.StringVar(&pass, "pass", "", "[required] the user password")
	flag.IntVar(&warnLimit, "warnlimit", 100, "minimum number of logs for outlier detection")
	flag.IntVar(&percentile, "percentile", 90, "the percentile for outlier detection")
	flag.Parse()
}

func main() {
	os.Exit(probes.RunFSSpeedProbe(target, user, pass, warnLimit, percentile))
}
