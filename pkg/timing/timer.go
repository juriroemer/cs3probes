package timing

import (
	"time"
	"errors"

	iop "github.com/cs3org/reva/pkg/sdk"
)

const (
	checkOK      = 0
	checkWarning = 1
	checkError   = 2
	checkUnknown = 3
)

// Takes a Test-function and a CS3APIs session object
// Times execution time of provided function, returns it in milliseconds
func TimeIopFunction(f func(*iop.Session) int, session *iop.Session) (int, error){
	start := time.Now()
	state := f(session)
	t := int(time.Since(start).Milliseconds())
	if state != checkOK{
		return state, errors.New("Function Error")
	}
	return t, nil
}
