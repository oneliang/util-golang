package file

import (
	"bufio"
	"os"
)

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
