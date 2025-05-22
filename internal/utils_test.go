package internal

import (
	"errors"
	"testing"
)

func TestCheckErr(t *testing.T) {
	CheckErr(nil)
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic")
		}
	}()
	CheckErr(errors.New("boom"))
}
