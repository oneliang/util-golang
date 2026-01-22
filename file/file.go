package file

import (
	"bufio"
	"os"
	"path/filepath"
)

func CreateFileWithDirectory(fullFilename string) error {
	directory := filepath.Dir(fullFilename)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	file, err := os.Create(fullFilename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	return nil
}

func WriteFileWithDirectory(fullFilename string, data []byte, permission os.FileMode) error {
	directory := filepath.Dir(fullFilename)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}
	if err := os.WriteFile(fullFilename, data, permission); err != nil {
		return err
	}
	return nil
}

func ReadFileContentEachLine(filename string, readLineProcessor func(content string) bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readNext := readLineProcessor(scanner.Text())
		if readNext {
			continue
		} else {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
