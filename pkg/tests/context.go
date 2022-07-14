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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/nagios"
	"github.com/cs3org/reva/pkg/sdk"
	"github.com/cs3org/reva/pkg/sdk/action"
)

type TestContext struct {
	session *sdk.Session

	messages []string
}
type TestIOPFunction = func(*sdk.Session, string) (int, error)
type TestNetworkFunction = func(string) (int, int, error)

func (ctx *TestContext) BeginTests() {
	ctx.messages = []string{}
}

func (ctx *TestContext) EndTests(outliers map[string]int) {
	if len(ctx.messages) > 0 || len(outliers) > 0 {
		ctx.dumpMessages(outliers)
	} else {
		fmt.Println("All tests succeeded w/o warnings")
	}
	ctx.cleanup()
}

func (ctx *TestContext) RunIOPTest(f TestIOPFunction, root string, testName string) int {
	res, dur, err := ctx.timeIOPFunction(f, ctx.session, root)
	switch res {
	case nagios.CheckWarning:
		ctx.messages = append(ctx.messages, fmt.Sprintf("%v[%vms] - WARNING: %v", testName, dur, err))

	case nagios.CheckError:
		ctx.messages = append(ctx.messages, fmt.Sprintf("%v[%vms] - FAILED: %v", testName, dur, err))
	}

	if res == nagios.CheckError {
		ctx.dumpMessages(nil)
		ctx.cleanup()
		os.Exit(res)
	}

	return res
}

func (ctx *TestContext) RunNetworkTest(f TestNetworkFunction, target string, testName string) (int, int) {
	res, val, err := f(target)
	switch res {
	case nagios.CheckWarning:
		ctx.messages = append(ctx.messages, fmt.Sprintf("%v - WARNING: %v", testName, err))

	case nagios.CheckError:
		ctx.messages = append(ctx.messages, fmt.Sprintf("%v - FAILED: %v", testName, err))
	}

	if res == nagios.CheckError {
		ctx.dumpMessages(nil)
		ctx.cleanup()
		os.Exit(res)
	}

	return res, val
}

func (ctx *TestContext) cleanup() {
	if ctx.session.IsValid() {
		// Physically remove all generated files from the recycle bin
		if recycleAct, err := action.NewRecycleOperationsAction(ctx.session); err == nil {
			_ = recycleAct.Purge()
		}
	}
}

func (ctx *TestContext) dumpMessages(outliers map[string]int) {
	msgs := strings.Join(ctx.messages, "; ")
	if len(outliers) > 0 {
		warnings := []string{}
		for test, dur := range outliers {
			warnings = append(warnings, fmt.Sprintf("test %v took %vms", test, dur))
		}
		if len(msgs) > 0 {
			msgs += "; "
		}
		msgs += "TIME WARNINGS: "
		msgs += strings.Join(warnings, "; ")
	}
	fmt.Println(msgs)
}

func (ctx *TestContext) timeIOPFunction(f TestIOPFunction, session *sdk.Session, root string) (int, int64, error) {
	start := time.Now()
	state, err := f(session, root)
	return state, time.Since(start).Milliseconds(), err
}

func NewTestContext(session *sdk.Session) (*TestContext, error) {
	ctx := &TestContext{session: session}
	return ctx, nil
}
