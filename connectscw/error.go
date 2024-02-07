package connectscw

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SdkError(err error) error {
	connectErr := &connect.Error{}
	if !errors.As(err, &connectErr) {
		return err
	}

	switch connectErr.Code() {
	case connect.CodeInvalidArgument:
		return &scw.InvalidArgumentsError{
			Details: connectErr.Details(),
			RawBody: nil,
		}
	}

	if connectErr := (&connect.Error{}); errors.As(err, &connectErr) {
		fmt.Println("underlying error message:", connectErr.Message())
	}

	return err
}
