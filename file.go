package gofs

import (
	"io"
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

func (f *File) Delete() error {
	err := os.Remove(f.String())
	return err
}

func (f *File) Copy(destDir Dir) (File, error) {
	// copy the file to the new path
	var srcFile *os.File
	if f.Exists() && f.IsFile() {
		var err error
		srcFile, err = os.Open(f.String())
		if err != nil {
			return File{}, err
		}
	} else {
		return File{}, os.ErrNotExist
	}
	err := destDir.CreateIfNotExist()
	if err != nil {
		return File{}, err
	}
	destFilePath := Join(destDir.String(), f.Name())
	destFile, err := os.Create(destFilePath.String())
	if err != nil {
		return File{}, err
	}
	_, err = io.Copy(destFile, srcFile)
	return destFilePath.File(), nil
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

func (f *File) ChunkReader(size int64) (func() ([]byte, error, bool), func() error, error) {
	if f.Exists() && f.IsFile() {
		file, err := os.Open(f.String())
		if err != nil {
			return nil, nil, err
		}
		return func() ([]byte, error, bool) {
				data := make([]byte, size)
				n, err := file.Read(data)
				if err != nil {
					return data, err, true
				}
				return data[:n], err, false
			}, func() error {
				return file.Close()
			}, nil
	}
	return nil, nil, os.ErrNotExist
}

func (f *File) ReadString() (string, error) {
	if f.Exists() && f.IsFile() {
		data, err := os.ReadFile(f.String())
		return string(data), err
	}
	return "", os.ErrNotExist
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
