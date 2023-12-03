package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	fromPath   = "testdata/input.txt"
	toPath     = "testdata/testcopy.txt"
	offsetTest = 0
	limitTest  = 0
)

func TestCopy(t *testing.T) {
	err := Copy(fromPath, toPath, offsetTest, limitTest)
	if err != nil {
		return
	}

	// get source file size
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return
	}

	defer srcFile.Close()

	srcStat, err := srcFile.Stat()
	if err != nil {
		return
	}

	// get destination file size
	destFile, err := os.Open(toPath)
	if err != nil {
		return
	}

	defer destFile.Close()

	destStat, err := srcFile.Stat()
	if err != nil {
		return
	}

	// src and dest expected to be equal size
	fmt.Printf("Source      file size: %d bytes\n", srcStat.Size())
	fmt.Printf("Destination file size: %d bytes\n", destStat.Size())
	require.Equal(t, srcStat.Size(), destStat.Size())

	// cleanup
	cmd := exec.Command("rm", "-f", toPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
