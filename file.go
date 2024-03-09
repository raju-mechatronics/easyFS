package gofs

import (
	"bufio"
	"io"
	"os"
)

// File represents a file with additional functionalities.
type File struct {
	PathHandler
}

// NewFile creates a new File instance with the given path.
//
// Example:
//
//	path := PathHandler("/path/to/file.txt") or "path/to/file.txt"
//	file := NewFile(path)
func NewFile(path PathHandler) File {
	return File{path}
}

// Size returns the size of the file in bytes.
func (f File) Size() (int64, error) {
	stat, err := os.Stat(f.String())
	return stat.Size(), err
}

// Delete deletes the file.
func (f File) Delete() error {
	return os.Remove(f.String())
}

// Copy copies the file to the specified destination directory.
// Copy the file to the new path
// Example:
//
//	src := NewFile("/path/to/source.txt")
//	dest := NewDir("/path/to/destination")
//	copiedFile, err := src.Copy(dest)
func (f File) Copy(destDir Dir) (File, error) {
	if f.Exists() && f.IsFile() {
		srcFile, err := os.Open(f.String())
		defer func() {
			err := srcFile.Close()
			if err != nil {
				panic(err)
			}
		}()
		if err != nil {
			return File{}, err
		}

		err = destDir.CreateIfNotExist()
		if err != nil {
			return File{}, err
		}

		destFilePath := destDir.Join(f.Name())
		destFile, err := os.Create(destFilePath.String())
		if err != nil {
			return File{}, err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return File{}, err
		}

		return destFilePath.File(), nil
	}
	return File{}, os.ErrNotExist
}

// Create creates the file. If the file already exists, it can overwrite it based on the 'overwrite' parameter.
//
// Args:
//   - overwrite: If true, overwrites the existing file. If false, does nothing if the file already exists.
//
// Returns:
//   - error: Any error encountered during the creation or overwrite.
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

// CreateIfNotExists creates the file if it does not exist already.
func (f File) CreateIfNotExists() error {
	return f.Create(false)
}

// Read reads the entire file and returns its content as a byte slice.
//
// Returns:
//   - []byte: The content of the file as a byte slice.
//   - error: Any error encountered during the read operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	data, err := file.Read()
func (f File) Read() ([]byte, error) {
	if f.Exists() && f.IsFile() {
		return os.ReadFile(f.String())
	}
	return nil, os.ErrNotExist
}

// ChunkReader returns a function to read the file in chunks of specified size.
//
// Args:
//   - size: The size of each chunk.
//
// Returns:
//   - func() ([]byte, error, bool): A function to read the file chunk by chunk by every call.
//   - func() error: A function to close the file.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	readChunk, closer, err := file.ChunkReader(1024)
//
//	for dataChunk := readChunk(); dataChunk != nil; dataChunk = readChunk() {
//	  // Do something with the data chunk
//	}
//	err = closer()
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

// ReadString reads the entire file and returns its content as a string.
//
// Returns:
//   - string: The content of the file as a string.
//   - error: Any error encountered during the read operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	data, err := file.ReadString()
func (f File) ReadString() (string, error) {
	data, err := f.Read()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// IterateLine returns a function to iterate through each line of the file.
//
// Returns:
//   - func() (string, error): A function to iterate through each line of the file.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	iterator, err := file.IterateLine()
func (f File) IterateLine() (func() (string, error), error) {
	if f.Exists() && f.IsFile() {
		file, err := os.Open(f.String())
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

// Write writes the given data to the file.
//
// Args:
//   - data: The data to write to the file.
//
// Returns:
//   - error: Any error encountered during the write operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	err := file.Write([]byte("Hello, World!"))
func (f File) Write(data []byte) error {
	return os.WriteFile(f.String(), data, 0644)
}

// WriteString writes the given string data to the file.
//
// Args:
//   - data: The string data to write to the file.
//
// Returns:
//   - error: Any error encountered during the write operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	err := file.WriteString("Hello, World!")
func (f File) WriteString(data string) error {
	return f.Write([]byte(data))
}

// AppendString appends the given string data to the file, optionally adding a newline.
//
// Args:
//   - data: The string data to append to the file.
//   - newLine: If true, adds a newline character to the data before appending.
//
// Returns:
//   - error: Any error encountered during the append operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	err := file.AppendString("Hello, World!", true)
func (f File) AppendString(data string, newLine bool) error {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_WRONLY, 0644)
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

// AppendIterative returns a function to iteratively append data to the file.
//
// Returns:
//   - func(data []byte) error: A function to iteratively append data to the file.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	file := NewFile(PathHandler("/path/to/file.txt"))
//	appendNext, err := file.AppendIterative()
//
// appendNext([]byte("Hello, World!")
func (f File) AppendIterative() (func(data []byte) error, error) {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return func(data []byte) error {
		// Write the data to the file
		_, err := file.Write(data)
		if err != nil {
			return err
		}
		return nil
	}, nil
}

// AppendStringIterative returns a function to iteratively append string data to the file.
func (f *File) AppendStringIterative() (func(data string) error, error) {
	// Open the file in append mode
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return func(data string) error {
		// Write the data to the file
		_, err := file.WriteString(data)
		if err != nil {
			return err
		}
		return nil
	}, nil
}
