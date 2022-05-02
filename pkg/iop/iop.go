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

package iop

import (
	"github.com/cs3org/reva/pkg/sdk"
	"github.com/pkg/errors"
)

// CreateSession creates a new IOP session.
func CreateSession(target string, username, password string) (*sdk.Session, error) {
	session := sdk.MustNewSession()
	if err := session.Initiate(target, false); err != nil {
		return nil, errors.Wrap(err, "unable to initiate session")
	}
	err := session.BasicLogin(username, password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login")
	}
	return session, nil
}
