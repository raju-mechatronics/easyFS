package gofs

import (
	"log"
	"os"
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
	//srcFile, err := os.Open(f.String())
	//srcFile.Wri

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
	if f.Exists() && f.IsFile() {
		data, err := os.ReadFile(f.String())
		return data, err
	}
	return nil, os.ErrNotExist
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

type Reader struct {
	file *os.File
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read string from file without ioutil
	data, err := file.Read()
	file.ReadFrom()

}
