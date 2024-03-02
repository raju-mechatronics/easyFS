# EasyFS - Go File System Library

EasyFS is a powerful Go library for handling file system operations in various operating systems. Whether you're working on Windows, macOS, or Linux, EasyFS provides a unified interface for managing files and directories with ease.

## Motivation

Managing file systems can be challenging, especially when dealing with differences across various operating systems. I built EasyFS to provide a simple and consistent way to handle file system operations in Go, regardless of the underlying platform. By abstracting away the complexities of interacting with the file system, EasyFS allows developers to focus on building robust applications without worrying about platform-specific intricacies.

## Features
EasyFS provides the following features:

 * Convenient abstractions for working with files and directories.
 * Simplified methods for creating, deleting, and manipulating files and directories.
 * Easy-to-understand API that reduces boilerplate code and improves code readability.

## Installation

To use EasyFS in your Go projects, simply import it using Go modules:

```bash
go get github.com/raju-mechatronics/EasyFS
```

## Usage

Here's a quick example demonstrating how to use EasyFS to perform common file system operations:

```go
package main

import (
	"fmt"
	"github.com/raju-mechatronics/EasyFS"
)

func main() {
	// Create a directory
	dir := EasyFS.NewDir("/path/to/directory")
	err := dir.CreateIfNotExist()
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create a file inside the directory with overwrite enable
	file, err := dir.CreateFile("example.txt", true)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
    // write something on file
    file.WriteString("Hello World")
    fmt.Println("Directory and file created successfully!")

    targetDir := NewDir("/path/to/Dir")
    targetDir.Copy("path/to/Target") //also can recive another dir or pathhandler
}

```

# EasyFS Documentation

## Package `EasyFS`

### Struct `PathHandler`

Represents a file path or directory path.

#### Methods

- `String() string`: Converts the path handler to a string.
- `IsValidPath() bool`: Checks if the path is valid.
- `Exists() bool`: Checks if the path exists.
- `IsDir() bool`: Checks if the path is a directory.
- `IsFile() bool`: Checks if the path is a file.
- `IsSymlink() bool`: Checks if the path is a symbolic link.
- `Stat() (PathInfo, error)`: Retrieves file information.
- `Lstat() (PathInfo, error)`: Retrieves file information without following symbolic links.
- `IsAbs() bool`: Checks if the path is absolute.
- `Resolve() (string, error)`: Resolves the path to an absolute path.
- `Abs() (string, error)`: Returns the absolute path.
- `Name() string`: Returns the base name of the path.
- `Parent() Dir`: Returns the parent directory of the path.
- `Ext() string`: Returns the extension of the file.
- `Join(elem ...string) string`: Joins path elements into a single path.
- `IsRel() bool`: Checks if the path is relative.
- `IsSameDir(other PathHandler) bool`: Checks if the path is in the same directory as another path.
- `IsSiblingOf(other PathHandler) bool`: Checks if the path is a sibling of another path.
- `IsDescendantOf(other PathHandler) bool`: Checks if the path is a descendant of another path.
- `DeletePath(force bool) error`: Deletes the path. If `force` is true, deletes recursively for directories.
- `File() File`: Converts the path handler to a file object.
- `Dir() Dir`: Converts the path handler to a directory object.
- `Rename(newName string) error`: Renames the file or directory.
- `Move(newPath PathHandler) error`: Moves the file or directory to a new path.
- `SetPerm(perm os.FileMode) error`: Sets the permission of the file or directory.

### Struct `Dir`

Represents a directory in the file system. Inherit the Path Handler. all the methods of PathHandler are available here.

#### Methods

- `CreateIfNotExist() error`: Creates the directory if it doesn't already exist.
- `All() ([]PathHandler, error)`: Retrieves all paths within the directory.
- `Files() ([]File, error)`: Retrieves all files within the directory.
- `Dirs() ([]Dir, error)`: Retrieves all subdirectories within the directory.
- `Delete(recursive bool) error`: Deletes the directory. If `recursive` is true, deletes all contents recursively.
- `DeleteSubFile(name string) error`: Deletes a file within the directory by name.
- `DeleteSubDir(name string, recursive bool) error`: Deletes a subdirectory within the directory by name. If `recursive` is true, deletes all contents recursively.
- `Copy(dest PathHandler) error`: Copies the directory to the specified destination.
- `HasDir(name string) bool`: Checks if a subdirectory exists within the directory.
- `HasFile(name string) bool`: Checks if a file exists within the directory.
- `Find(match string, recursive bool, quantity int) []PathHandler`: Finds paths matching a pattern within the directory.
- `FindFile(match string, recursive bool, quantity int) []File`: Finds files matching a pattern within the directory.
- `FindDir(match string, recursive bool, quantity int) []Dir`: Finds subdirectories matching a pattern within the directory.
- `CreateDir(name string) (Dir, error)`: Creates a subdirectory with the specified name.
- `CreateFile(name string, overwrite bool) (File, error)`: Creates a file within the directory with the specified name. If `overwrite` is true, overwrites the file if it already exists.
- `CreateFileWithData(name string, data []byte, overwrite bool) (File, error)`: Creates a file with the specified name and writes the given data to it.
- `CreateFileWithString(name string, data string, overwrite bool) File`: Creates a file with the specified name and writes the given data (string) to it.
- `GetTree() DirStructure`: Returns the directory structure as a tree.
- `GetAllPathExists() []PathHandler`: Returns all paths existing within the directory.
- `Clear(force bool) error`: Clears all contents within the directory. If `force` is true, deletes all contents recursively.
- `IsEmpty() bool`: Checks if the directory is empty.

### Struct `File`

Represents a file in the file system. Also Inherit the Path Handler. all the methods of PathHandler are available here.

#### Methods

- `Size() (int64, error)`: Retrieves the size of the file.
- `Delete() error`: Deletes the file.
- `Copy(destDir Dir) (File, error)`: Copies the file to the specified destination directory.
- `Create(overwrite bool) error`: Creates the file. If `overwrite` is true, overwrites the file if it already exists.
- `CreateIfNotExists() error`: Creates the file if it doesn't already exist.
- `Read() ([]byte, error)`: Reads the contents of the file.
- `ChunkReader(size int64) (func() ([]byte, error, bool), func() error, error)`: Reads the file in chunks of the specified size.
- `ReadString() (string, error)`: Reads the contents of the file as a string.
- `IterateLine() (func() (string, error), error)`: Iterates over each line of the file.
- `Write(data []byte) error`: Writes data to the file.
- `WriteString(data string) error`: Writes a string to the file.
- `AppendString(data string, newLine bool) error`: Appends a string to the file. If `newLine` is true, adds a newline character.
- `AppendIterative() (func(data []byte) error, error)`: Appends data to the file iteratively.
- `AppendStringIterative() (func(data string) error, error)`: Appends a string to the file iteratively.

## ⚠️ Attention

**Note**: EasyFS is still under active development and may not be fully tested or completed yet. While the library offers a range of functionalities for interacting with the file system in Go, there may be areas that require further refinement or optimization.

**Community Help**: We welcome contributions from the community to help improve EasyFS. If you encounter any bugs, issues, or have suggestions for enhancements, please don't hesitate to open an issue or submit a pull request on our [GitHub repository](https://github.com/your/repository).

**Testing**: Although we strive to ensure the stability and reliability of EasyFS, it's important to note that thorough testing is ongoing. As such, we recommend exercising caution when using the library in production environments or mission-critical projects.

Thank you for your understanding and support as we work towards making EasyFS a robust and reliable tool for file system operations in Go.


## Compatibility

EasyFS is compatible with the following operating systems:

- Windows
- macOS
- Linux

## Contributing

Contributions to EasyFS are welcome! If you encounter any issues, have feature requests, or would like to contribute code, please feel free to open an issue or submit a pull request.

## License

EasyFS is licensed under the MIT License. 
