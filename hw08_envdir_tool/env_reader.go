package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Prepare results with map
	result := make(Environment)

	// read nested directories recursively
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	//read current directory's content
	for _, item := range readDir {
		// имя S не должно содержать =;
		name := strings.ReplaceAll(item.Name(), "=", "")
		// open file
		file, err := os.Open(dir + "/" + name)
		if err != nil {
			return nil, fmt.Errorf("file open error: %w", err)
		}
		// get file statistics
		stat, err := file.Stat()
		if err != nil {
			return nil, fmt.Errorf("file stat error: %w", err)
		}
		// empty file == env to be removed
		if stat.Size() == 0 {
			result[name] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
			// each file need to be processed
			continue
		}

		// file exists & it's size > 0, need to process it's content
		res, err := prepEnv(file)
		if err != nil {
			return nil, fmt.Errorf("env prep error: %w", err)
		}

		result[name] = *res

		err = file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("file close error: %w", err))
		}
	}

	return result, nil
}

func prepEnv(reader io.Reader) (*EnvValue, error) {
	rd := bufio.NewReader(reader)

	line, err := rd.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("read string error: %w", err)
	}
	// терминальные нули (0x00) заменяются на перевод строки (\n);
	line = strings.ReplaceAll(line, string([]byte{0x00}), string('\n'))
	// ...файл с именем S, первой строкой которого является T... == ignore all after first line
	line = strings.TrimRight(line, string('\n'))
	// пробелы и табуляция в конце T удаляются;
	line = strings.TrimRight(line, " ")

	return &EnvValue{
		Value:      line,
		NeedRemove: false,
	}, nil
}
