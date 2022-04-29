package tests

import (
	"fmt"
	"net"
	"time"
	"crypto/rand"

	iop "github.com/cs3org/reva/pkg/sdk"
	action "github.com/cs3org/reva/pkg/sdk/action"
	ping "github.com/go-ping/ping"
)



const (
	checkOK      = 0
	checkWarning = 1
	checkError   = 2
	checkUnknown = 3
)

// Tests path enumeration
func Test_ls(session *iop.Session) (int) {
	enum := action.MustNewEnumFilesAction(session)
	_, err := enum.ListAll("/home/", true)
	if err != nil {
		fmt.Printf("%v\n", err)
		return checkError
	}
	//for _, f := range files {
	//fmt.Println(f.Path)}
	return checkOK
}

// Tests make directory
func Test_mkdir(session *iop.Session) (int) {
	mkdir := action.MustNewFileOperationsAction(session)
	err := mkdir.MakePath("/home/fsoperations/testdir")
	if err != nil{
		fmt.Printf("%v\n", err)
		return checkError
	}
	return checkOK
}

// Tests if a directory exists
func Test_direxists(session *iop.Session) (int) {
	direxists := action.MustNewFileOperationsAction(session)
	if direxists.DirExists("/home/fsoperations/testdir") {
		return checkOK
	}
	return checkError
}

// Tests to delete a directory
func Test_rmdir(session *iop.Session) (int) {
	rmdir := action.MustNewFileOperationsAction(session)
	err := rmdir.Remove("/home/fsoperations/testdir")
	if err != nil {
		fmt.Printf("%v\n", err)
		return checkError
	}
	if rmdir.DirExists("/home/fsoperations/testdir") {
		return checkError
	}

	return checkOK
}

// Tests to upload a file
func Test_upload(session *iop.Session) int {
	upload := action.MustNewUploadAction(session)
	upload.EnableTUS = true
	_, err := upload.UploadBytes([]byte("Hello World\n"), "/home/fsoperations/test.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return checkError
	}
	return checkOK
}

// Tests if a file exists
func Test_fileexists(session *iop.Session) (int) {
	fileexists := action.MustNewFileOperationsAction(session)
	if fileexists.FileExists("/home/fsoperations/test.txt") {
		return checkOK
	}
	return checkError
}

// Tests to move file to different location
func Test_mvfile(session *iop.Session) (int) {
	mv := action.MustNewFileOperationsAction(session)
	err := mv.Move("/home/fsoperations/test.txt", "/home/fsoperations/testmoved.txt")
	if err != nil {
		return checkError
	}
	return checkOK
}

// Tests to delete a file
func Test_rmfile(session *iop.Session) (int) {
	rmfile := action.MustNewFileOperationsAction(session)
	err := rmfile.Remove("/home/fsoperations/testmoved.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return checkError
	}
	if rmfile.FileExists("/home/fsoperations/testmoved.txt") {
		return checkError
	}
	return checkOK
}

// Test to ping a target system, returns measured time and checkOk or 0 and checkError
// returns 0 and checkUnknown if it is blocked by a firewall
func Test_ping(target string) (int, int) {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return 0, checkError
	}
	pinger.Count = 3
	pinger.Timeout = time.Second
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return 0, checkError
	}
	stats := pinger.Statistics()
	fmt.Println(stats.AvgRtt)
	if (stats.AvgRtt == 0) {

		return 0, checkUnknown
	}
	fmt.Println(stats.AvgRtt == 0)
	return int(stats.AvgRtt / time.Millisecond), checkOK
}

// Tests to portscan the target system
func Test_portscan(target string) int {
	conn, _ := net.DialTimeout("tcp", target, time.Millisecond*300)
	if conn != nil {
		conn.Close()
		return checkOK
	}
	return checkError
}

// Tests to upload 10 small randomly generated 10kb files
func Test_sUpload(session *iop.Session) int {
	for i := 0; i<10; i++ {
		upload := action.MustNewUploadAction(session)
		upload.EnableTUS = true
		data, _ := generateData(10)
		_, err := upload.UploadBytes(data, "/home/writespeed/small" + fmt.Sprint(i) + ".txt")
		if err != nil {
			fmt.Printf("%v\n", err)
			return checkError
		}
	}

	return checkOK
}

// Tests upload of 1 bigger randomly generated 100kb file
func Test_bUpload(session *iop.Session) int {
	upload := action.MustNewUploadAction(session)
	upload.EnableTUS = true
		data, _ := generateData(100)
		_, err := upload.UploadBytes(data, "/home/writespeed/big.txt")
		if err != nil {
			fmt.Printf("%v\n", err)
			return checkError
		}

	return checkOK
}

// Tests to move 10 small randomly generated 10kb files
func Test_sMove(session *iop.Session) int {
	sMove := action.MustNewFileOperationsAction(session)
	for i := 0; i<10; i++ {
		err := sMove.Move("/home/writespeed/small" + fmt.Sprint(i) +".txt", "/home/writespeed/smallmoved" + fmt.Sprint(i) + ".txt")
		if err != nil {
			fmt.Printf("%v\n", err)
			return checkError
		}
	}
	return checkOK
}

// Tests to move 1 bigger randomly generated 100kb files
func Test_bMove(session *iop.Session) int {
	bMove := action.MustNewFileOperationsAction(session)
		err := bMove.Move("/home/writespeed/big.txt", "/home/writespeed/bigmoved.txt")
		if err != nil {
			fmt.Printf("%v\n", err)
			return checkError
		}
	return checkOK
}

// Tests to remove 10 small randomly generated 10kb files
func Test_sRemove(session *iop.Session) int {
	sRemove := action.MustNewFileOperationsAction(session)
	for i := 0; i<10; i++ {
		err := sRemove.Remove("/home/writespeed/smallmoved" + fmt.Sprint(i) + ".txt")
		if err != nil {
			fmt.Printf("%v\n", err)
			return checkError
		}
	}
	return checkOK
}

// Tests to remove 1 bigger randomly generated 100kb file
func Test_bRemove(session *iop.Session) int {
	bRemove := action.MustNewFileOperationsAction(session)
	err := bRemove.Remove("/home/writespeed/bigmoved.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return checkError
	}
	return checkOK
}


// generates random data of given size
func generateData(size int) ([]byte, error) {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("Failed to generate random data")
	}
	return data, err
}