package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name         string
		fromPath     string
		offset       int64
		limit        int64
		err          error
		errEqualFlag bool
	}{
		{"Offset exceeds file size", "testdata/input.txt", 7000, 0, ErrOffsetExceedsFileSize, true},
		{"Unsupported file", "testdata", 0, 0, ErrUnsupportedFile, true},
		{"File not found", "testdata/input1.txt", 0, 0, ErrUnsupportedFile, false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			out, outErr := os.CreateTemp("", "out")
			if outErr != nil {
				return
			}
			defer os.Remove(out.Name())

			err := Copy(tc.fromPath, out.Name(), tc.offset, tc.limit)
			if tc.errEqualFlag {
				require.EqualError(t, tc.err, err.Error())
			} else {
				require.Error(t, err)
			}
		})
	}
}
