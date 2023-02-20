package error

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestErrorWrap(t *testing.T) {
	// use %w to wrap error
	err := fmt.Errorf("error: %w", io.EOF)
	t.Log(err)

	// use errors.Is to check error
	if errors.Is(err, io.EOF) {
		t.Log("err is EOF")
	}

	// use errors.As to check error
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		t.Log("err is PathError")
	}
}
