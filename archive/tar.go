package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."
	tarFileName := filepath.Join(dir, "files.tar")

	f, err := os.OpenFile(tarFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	checkError(err)
	defer f.Close()

	filesToTar := []string{}
	entries, err := os.ReadDir(dir)
	checkError(err)
	for _, ent := range entries {
		if !ent.IsDir() && strings.HasSuffix(ent.Name(), ".txt") {
			filesToTar = append(filesToTar, filepath.Join(dir, ent.Name()))
		}
	}

	err = archiveFiles(f, filesToTar)
	checkError(err)

	err = readArchive(tarFileName)
	checkError(err)

}

func archiveFiles(f io.Writer, files []string) error {
	tw := tar.NewWriter(f)
	defer tw.Close()

	for _, filename := range files {
		if err := addToArchive(filename, tw); err != nil {
			return err
		}
	}

	return nil
}

func addToArchive(name string, tw *tar.Writer) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}

	hdr := &tar.Header{
		ModTime: fileinfo.ModTime(),
		Name:    name,
		Size:    fileinfo.Size(),
		Mode:    int64(fileinfo.Mode().Perm()),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}

	copied, err := io.Copy(tw, file)
	if err != nil {
		return err
	}

	if copied < fileinfo.Size() {
		return fmt.Errorf("Size of the copied file doesn't match with source file %s: $s", name, err)
	}

	return nil
}

func readArchive(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		size := hdr.Size
		contents := make([]byte, size)
		read, err := io.ReadFull(tr, contents)
		if int64(read) != size {
			return fmt.Errorf("Size of the opened file doesn't match with the file %s", hdr.Name)
		}

		fmt.Printf("Contents of the file %s:\n", hdr.Name)
		fmt.Fprintf(os.Stdout, "\n%s", contents)
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
