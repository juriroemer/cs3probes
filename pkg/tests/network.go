// Copyright 2018-2020 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package tests

import (
	"net"
	"time"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/nagios"
	"github.com/go-ping/ping"
	"github.com/pkg/errors"
)

// Test to ping a target system, returns measured time and probes.CheckOK or 0 and probes.CheckError
// returns 0 and checkUnknown if it is blocked by a firewall
func Test_ping(target string) (int, int, error) {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return nagios.CheckError, 0, errors.Wrap(err, "unable to create pinger")
	}
	pinger.Count = 3
	pinger.Timeout = time.Second
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return nagios.CheckError, 0, errors.Wrap(err, "unable to ping")
	}
	stats := pinger.Statistics()
	if stats.AvgRtt == 0 {
		return nagios.CheckUnknown, 0, errors.Errorf("ping average of 0")
	}
	return nagios.CheckOK, int(stats.AvgRtt / time.Millisecond), nil
}

// Tests to portscan the target system
func Test_portscan(target string) (int, int, error) {
	conn, err := net.DialTimeout("tcp", target, time.Millisecond*300)
	if conn != nil {
		conn.Close()
		return nagios.CheckOK, 0, nil
	}
	return nagios.CheckError, 0, errors.Wrap(err, "unable to dial out")
}
