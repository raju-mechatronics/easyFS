package gofs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type File struct {
	PathHandler
}

func NewFile(path PathHandler) File {
	return File{path}
}

func (f *File) Size() (int64, error) {
	// get the file size
	stat, err := os.Stat(f.String())
	return stat.Size(), err
}

func (f *File) GetMetaData() {
	log.Fatal("Not implemented")
}

func (f *File) Delete() error {
	err := os.Remove(f.String())
	return err
}

func (f *File) Copy(destPath PathHandler) error {
	// copy the file to the new path
	destDir := destPath.Dir()
	err := destDir.CreateIfNotExist()
	if err != nil {
		return err
	}
	srcFile, err := os.Open(f.String())
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(filepath.Join(destPath.String(), f.Name()))
	if err != nil {
		return fmt.Errorf("creating destination file: %w", err)
	}
	defer dstFile.Close()

	// Allocate a buffer
	buf := make([]byte, 4096)
	for {
		n, err := syscall.Read(int(srcFile.Fd()), buf)
		if err != nil {
			return fmt.Errorf("reading from source file: %w", err)
		}
		if n == 0 {
			break // EOF
		}
		_, err = syscall.Write(int(dstFile.Fd()), buf[:n])
		if err != nil {
			return fmt.Errorf("writing to destination file: %w", err)
		}
	}

	return nil
}

func (f *File) Create(overwrite bool) error {
	if f.IsFile() || f.Exists() {
		if overwrite {
			f.Delete()
		} else {
			return nil
		}
	}
	_, err := os.Create(f.String())
	return err
}

func (f *File) CreateIfNotExists() error {
	return f.Create(false)
}

func (f *File) Read() ([]byte, error) {

}

func (f *File) ChunkReader() (func(size int32) ([]byte, error), error) {

}

func (f *File) ReadString() (string, error) {

}

func (f *File) ReadStringChunk() (func(size int32) (string, error), error) {

}

func (f *File) IterateLine() (func() (string, error), error) {

}

func (f *File) Write(data []byte) error {

}

func (f *File) WriteString(data string) error {

}

func (f *File) Append(data []byte) error {

}

func (f *File) AppendString(data string, newLine bool) error {

}

func (f *File) AppendIterative() (func(data []byte) error, error) {

}

func (f *File) AppendStringIterative() (func(data string) error, error) {
}
