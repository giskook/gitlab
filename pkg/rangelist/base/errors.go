package base

import "fmt"

type ErrorCode int

const (
	ErrorRangeIsUnavailable ErrorCode = 400000
	ErrorRangeListIsNil     ErrorCode = 500000
)

var errDesc = map[ErrorCode]string{
	ErrorRangeIsUnavailable: "range is unavailable",
	ErrorRangeListIsNil:     "range list is nil",
}

func (e ErrorCode) Desc() string {
	return errDesc[e]
}

func (e ErrorCode) Error() error {
	return fmt.Errorf("error code is: %v desc: %v", e, e.Desc())
}
