package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."
	zipFileName := filepath.Join(dir, "files.zip")

	f, err := os.OpenFile(zipFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	checkError(err)
	defer f.Close()

	filesToZip := []string{}
	entries, err := os.ReadDir(dir)
	checkError(err)
	for _, ent := range entries {
		if !ent.IsDir() && strings.HasSuffix(ent.Name(), ".txt") {
			filesToZip = append(filesToZip, filepath.Join(dir, ent.Name()))
		}
	}

	err = zipFiles(f, filesToZip)
	checkError(err)

	err = readZip(zipFileName)
	checkError(err)
}

func zipFiles(f io.Writer, filesToZip []string) error {
	zw := zip.NewWriter(f)
	defer zw.Close()

	for _, ftz := range filesToZip {
		if err := addToZip(ftz, zw); err != nil {
			return err
		}
	}

	return nil
}

func addToZip(name string, zw *zip.Writer) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	wr, err := zw.Create(name)
	if err != nil {
		return err
	}

	if _, err := io.Copy(wr, file); err != nil {
		return err
	}

	return nil
}

func readZip(name string) error {
	rc, err := zip.OpenReader(name)
	if err != nil {
		return err
	}
	defer rc.Close()

	for _, file := range rc.File {
		frc, err := file.Open()
		if err != nil {
			return err
		}
		defer frc.Close()

		fmt.Fprintf(os.Stdout, "Contents of the file: %s \n", file.Name)
		copied, err := io.Copy(os.Stdout, frc)
		if err != nil {
			return err
		}

		if uint64(copied) != file.UncompressedSize64 {
			return fmt.Errorf("Length of the file contents doesn't match with the file %s", file.Name)
		}
		fmt.Println()
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
