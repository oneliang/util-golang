package file

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Zip(zipFilename string, fromDirectoryOrFile string) error {
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer func() { _ = zipFile.Close() }()

	zipWriter := zip.NewWriter(zipFile)
	defer func() { _ = zipWriter.Close() }()
	fileInfo, err := os.Stat(fromDirectoryOrFile)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() { //directory
		fromDirectoryAbsolutePath, err := filepath.Abs(fromDirectoryOrFile)
		if err != nil {
			return err
		}
		err = filepath.Walk(fromDirectoryAbsolutePath, func(path string, info fs.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if !info.IsDir() {
				zipEntryName := path[len(fromDirectoryAbsolutePath)+1:]
				if err = addZipFileEntry(zipWriter, zipEntryName, path); err != nil {
					return err
				}
			} else {
				//directory
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else { //file
		zipEntryName := filepath.Base(fromDirectoryOrFile)
		if err = addZipFileEntry(zipWriter, zipEntryName, fromDirectoryOrFile); err != nil {
			return err
		}
	}
	return nil
}

func addZipFileEntry(zipWriter *zip.Writer, zipEntryName string, fromFilePath string) error {
	ioWriter, err := zipWriter.Create(zipEntryName)
	if err != nil {
		return err
	}
	file, err := os.Open(fromFilePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if _, err := io.Copy(ioWriter, file); err != nil {
		return err
	}
	return nil
}

func UnZip(zipFile string, outputFilePath string) error {
	// 第一步，打开 zip 文件
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func() { _ = zipReader.Close() }()

	for _, zipEntry := range zipReader.File {
		filePath := zipEntry.Name
		if zipEntry.FileInfo().IsDir() {
			_ = os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipEntry.Mode())
		if err != nil {
			return err
		}
		file, err := zipEntry.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(destinationFile, file); err != nil {
			return err
		}
		defer func() { _ = destinationFile.Close() }()
		defer func() { _ = file.Close() }()
	}
	return nil
}
