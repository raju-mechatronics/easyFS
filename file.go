package gofs

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	PathHandler
}

func NewFile(path PathHandler) File {
	return File{path}
}

func (f File) Size() (int64, error) {
	// get the file size
	stat, err := os.Stat(f.String())
	return stat.Size(), err
}

func (f File) Delete() error {
	err := os.Remove(f.String())
	return err
}

func (f File) Copy(destDir Dir) (File, error) {
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
	// close the files
	srcFile.Close()
	destFile.Close()
	if err != nil {
		return File{}, err
	}
	return destFilePath.File(), nil
}

func (f File) Create(overwrite bool) error {
	if f.IsFile() || f.Exists() {
		if overwrite {
			f.Delete()
		} else {
			return nil
		}
	}
	file, err := os.Create(f.String())
	if err == nil {
		file.Close()
	}
	return err
}

func (f File) CreateIfNotExists() error {
	return f.Create(false)
}

func (f File) Read() ([]byte, error) {
	if f.Exists() && f.IsFile() {
		data, err := os.ReadFile(f.String())
		return data, err
	}
	return nil, os.ErrNotExist
}

func (f File) ChunkReader(size int64) (func() ([]byte, error, bool), func() error, error) {
	if f.Exists() && f.IsFile() {
		file, err := os.Open(f.String())
		if err != nil {
			return nil, nil, err
		}
		return func() ([]byte, error, bool) {
				data := make([]byte, size)
				n, err := file.Read(data)
				if err != nil {
					file.Close()
					return data, err, true
				}
				return data[:n], err, false
			}, func() error {
				return file.Close()
			}, nil
	}
	return nil, nil, os.ErrNotExist
}

func (f File) ReadString() (string, error) {
	if f.Exists() && f.IsFile() {
		data, err := os.ReadFile(f.String())
		return string(data), err
	}
	return "", os.ErrNotExist
}

func (f File) IterateLine() (func() (string, error), error) {
	if f.Exists() && f.IsFile() {
		file, err := os.Open(f.String())
		defer file.Close()
		if err != nil {
			return nil, err
		}
		reader := bufio.NewReader(file)
		return func() (string, error) {
			line, err := reader.ReadString('\n')
			if err != nil {
				return line, err
			}
			return line[:len(line)-1], nil
		}, nil
	}
	return nil, os.ErrNotExist
}

func (f File) Write(data []byte) error {
	err := os.WriteFile(f.String(), data, 0644)
	return err
}

func (f File) WriteString(data string) error {
	return f.Write([]byte(data))
}

func (f File) AppendString(data string, newLine bool) error {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	defer file.Close()

	// If newLine is true, add a newline character to the data
	if newLine {
		data = "\n" + data
	}

	// Write the data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func (f File) AppendIterative() (func(data []byte) error, error) {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return func(data []byte) error {
		// Write the data to the file
		_, err := file.Write(data)
		if err != nil {
			return err
		}
		return nil
	}, nil

}

func (f *File) AppendStringIterative() (func(data string) error, error) {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return func(data string) error {
		// Write the data to the file
		_, err := file.WriteString(data)
		if err != nil {
			return err
		}
		return nil
	}, nil
}
