package main

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// open source file
	srcFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("error: no such file: %w", err)
		}
		return fmt.Errorf("error: cannot open file: %w", err)
	}

	// check it's length to omit unlimited files like /dev/urandom
	/*
		$ stat /dev/urandom
		  File: /dev/urandom
		  Size: 0         	Blocks: 0          IO Block: 4096   character special file
		...
	*/
	stat, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("error: cannot Stat() file: %w", err)
	}
	scrSize := stat.Size()

	// check if offset applicable
	if scrSize < offset {
		return fmt.Errorf("could not offet size: %w", ErrOffsetExceedsFileSize)
	}

	// limit < 0 == whole file is copied
	if limit > 0 {
		// new size to be copied
		if offset+limit > scrSize {
			// file size exceeded, ignore limit
			scrSize = scrSize - offset
		} else {
			// file size is not exceeded, limit should be used
			scrSize = limit
		}
	} else {
		fmt.Println("limit ignored:", limit)
	}

	// wipe or create destination file
	destFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("error: cannot open file: %w", err)
	}

	// close source file on return
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error: cannot close srcFile: %w", err))
		}
	}(srcFile)

	// close destination file on return
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error: cannot close destFile: %w", err))
		}
	}(destFile)

	// seek for and check offset
	_, err = srcFile.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("error: cannot seek for offset in srcFile: %w", err)
	}

	// create progress bar
	bar := pb.Full.Start64(scrSize)
	bar.SetRefreshRate(100 * time.Millisecond)

	// bar proxy to srcFile
	barReader := bar.NewProxyReader(srcFile)

	// copy from proxy
	_, err = io.CopyN(destFile, barReader, scrSize)
	if err != nil {
		return fmt.Errorf("error: cannot copyN: %w", err)
	}

	bar.Finish()

	return nil
}
