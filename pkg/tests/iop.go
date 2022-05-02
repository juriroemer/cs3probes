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
	"crypto/rand"
	"fmt"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/nagios"
	"github.com/cs3org/reva/pkg/sdk"
	"github.com/cs3org/reva/pkg/sdk/action"
	"github.com/pkg/errors"
)

// Tests path enumeration
func Test_ls(session *sdk.Session) (int, error) {
	enum := action.MustNewEnumFilesAction(session)
	_, err := enum.ListAll("/home/", true)
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to enumerate home files")
	}
	return nagios.CheckOK, nil
}

// Tests make directory
func Test_mkdir(session *sdk.Session) (int, error) {
	mkdir := action.MustNewFileOperationsAction(session)
	err := mkdir.MakePath("/home/fsoperations/testdir")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to create test directory")
	}
	return nagios.CheckOK, nil
}

// Tests if a directory exists
func Test_direxists(session *sdk.Session) (int, error) {
	direxists := action.MustNewFileOperationsAction(session)
	if direxists.DirExists("/home/fsoperations/testdir") {
		return nagios.CheckOK, nil
	}
	return nagios.CheckError, errors.Errorf("test directory doesn't exist")
}

// Tests to delete a directory
func Test_rmdir(session *sdk.Session) (int, error) {
	rmdir := action.MustNewFileOperationsAction(session)
	err := rmdir.Remove("/home/fsoperations/testdir")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to remove test directory")
	}
	if rmdir.DirExists("/home/fsoperations/testdir") {
		return nagios.CheckError, errors.Errorf("test directory still exists")
	}
	return nagios.CheckOK, nil
}

// Tests to upload a file
func Test_upload(session *sdk.Session) (int, error) {
	upload := action.MustNewUploadAction(session)
	upload.EnableTUS = true
	_, err := upload.UploadBytes([]byte("Hello World\n"), "/home/fsoperations/test.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to upload test file")
	}
	return nagios.CheckOK, nil
}

// Tests if a file exists
func Test_fileexists(session *sdk.Session) (int, error) {
	fileexists := action.MustNewFileOperationsAction(session)
	if fileexists.FileExists("/home/fsoperations/test.txt") {
		return nagios.CheckOK, nil
	}
	return nagios.CheckError, errors.Errorf("test file doesn't exist")
}

// Tests to move file to different location
func Test_mvfile(session *sdk.Session) (int, error) {
	mv := action.MustNewFileOperationsAction(session)
	err := mv.Move("/home/fsoperations/test.txt", "/home/fsoperations/testmoved.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to move test file")
	}
	return nagios.CheckOK, nil
}

// Tests to delete a file
func Test_rmfile(session *sdk.Session) (int, error) {
	rmfile := action.MustNewFileOperationsAction(session)
	err := rmfile.Remove("/home/fsoperations/testmoved.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to remove test file")
	}
	if rmfile.FileExists("/home/fsoperations/testmoved.txt") {
		return nagios.CheckError, errors.Errorf("test file still exists")
	}
	return nagios.CheckOK, nil
}

// Tests to upload 10 small randomly generated 10kb files
func Test_sUpload(session *sdk.Session) (int, error) {
	for i := 0; i < 10; i++ {
		upload := action.MustNewUploadAction(session)
		upload.EnableTUS = true
		data := generateData(10 * 1024)
		targetFile := "/home/writespeed/small" + fmt.Sprint(i) + ".txt"
		_, err := upload.UploadBytes(data, targetFile)
		if err != nil {
			return nagios.CheckError, errors.Wrapf(err, "unable to upload test file %v", targetFile)
		}
	}

	return nagios.CheckOK, nil
}

// Tests upload of 1 bigger randomly generated 100kb file
func Test_bUpload(session *sdk.Session) (int, error) {
	upload := action.MustNewUploadAction(session)
	upload.EnableTUS = true
	data := generateData(100 * 1024)
	_, err := upload.UploadBytes(data, "/home/writespeed/big.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to upload test file")
	}

	return nagios.CheckOK, nil
}

// Tests to move 10 small randomly generated 10kb files
func Test_sMove(session *sdk.Session) (int, error) {
	sMove := action.MustNewFileOperationsAction(session)
	for i := 0; i < 10; i++ {
		sourceFile := "/home/writespeed/small" + fmt.Sprint(i) + ".txt"
		targetFile := "/home/writespeed/smallmoved" + fmt.Sprint(i) + ".txt"
		err := sMove.Move(sourceFile, targetFile)
		if err != nil {
			return nagios.CheckError, errors.Wrapf(err, "unable to move file %v", sourceFile)
		}
	}
	return nagios.CheckOK, nil
}

// Tests to move 1 bigger randomly generated 100kb files
func Test_bMove(session *sdk.Session) (int, error) {
	bMove := action.MustNewFileOperationsAction(session)
	err := bMove.Move("/home/writespeed/big.txt", "/home/writespeed/bigmoved.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to move test file")
	}
	return nagios.CheckOK, nil
}

// Tests to remove 10 small randomly generated 10kb files
func Test_sRemove(session *sdk.Session) (int, error) {
	sRemove := action.MustNewFileOperationsAction(session)
	for i := 0; i < 10; i++ {
		targetFile := "/home/writespeed/smallmoved" + fmt.Sprint(i) + ".txt"
		err := sRemove.Remove(targetFile)
		if err != nil {
			return nagios.CheckError, errors.Wrapf(err, "unable to remove test file %v", targetFile)
		}
	}
	return nagios.CheckOK, nil
}

// Tests to remove 1 bigger randomly generated 100kb file
func Test_bRemove(session *sdk.Session) (int, error) {
	bRemove := action.MustNewFileOperationsAction(session)
	err := bRemove.Remove("/home/writespeed/bigmoved.txt")
	if err != nil {
		return nagios.CheckError, errors.Wrap(err, "unable to remove test file")
	}
	return nagios.CheckOK, nil
}

// generates random data of given size
func generateData(size int) []byte {
	data := make([]byte, size)
	_, _ = rand.Read(data)
	return data
}
