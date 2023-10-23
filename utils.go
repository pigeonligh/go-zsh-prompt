package gozshprompt

import (
	"fmt"
	"os"
)

func pipe() (parentRead *os.File, parentWrite *os.File, childRead *os.File, childWrite *os.File, err error) {
	parentRead, childWrite, err = os.Pipe()
	if err != nil {
		err = fmt.Errorf("failed to create parent pipe: %w", err)
		return nil, nil, nil, nil, err
	}
	childRead, parentWrite, err = os.Pipe()
	if err != nil {
		err = fmt.Errorf("failed to create child pipe: %w", err)
		parentRead.Close()
		childWrite.Close()
		return nil, nil, nil, nil, err
	}
	return
}
