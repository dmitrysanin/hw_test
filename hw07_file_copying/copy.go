package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	fi, err := src.Stat()
	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	srcSize := fi.Size()
	if offset > 0 {
		if offset > srcSize {
			return ErrOffsetExceedsFileSize
		}
		_, err = src.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}
	defer src.Close()

	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	cpLimit := limit
	if cpLimit == 0 {
		cpLimit = srcSize
	}
	if cpLimit > srcSize-offset {
		cpLimit = srcSize - offset
	}

	bar := pb.Full.Start64(cpLimit)
	barSrc := bar.NewProxyReader(src)

	_, err = io.CopyN(dst, barSrc, cpLimit)

	bar.Finish()

	return err
}
