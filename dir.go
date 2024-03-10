package easyFS

import (
	"os"
	"path/filepath"
)

// Dir represents a directory in the file system.
type Dir struct {
	PathHandler
}

func NewDir(path PathHandler) Dir {
	return Dir{path}
}

// DirStructure represents the structure of a directory.
type DirStructure struct {
	Dirs  map[string]DirStructure
	Files []File
}

// CreateIfNotExist creates the directory if it does not exist already.
//
// Returns:
//   - error: Any error encountered during the creation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	err := dir.CreateIfNotExist()
func (d Dir) CreateIfNotExist() error {
	if d.Exists() && d.IsDir() {
		return nil
	}
	return os.MkdirAll(d.String(), os.ModePerm)
}

// All returns all paths within the directory.
//
// Returns:
//   - []PathHandler: List of paths within the directory.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	paths, err := dir.All()
func (d Dir) All() ([]PathHandler, error) {
	entries, err := os.ReadDir(d.String())
	if err != nil {
		return nil, err
	}
	var paths []PathHandler
	for _, entry := range entries {
		paths = append(paths, PathHandler(filepath.Join(d.String(), entry.Name())))
	}
	return paths, nil
}

// Files returns all files within the directory.
//
// Returns:
//   - []File: List of files within the directory.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	files, err := dir.Files()
func (d Dir) Files() ([]File, error) {
	entries, err := d.All()
	if err != nil {
		return nil, err
	}
	var files []File
	for _, entry := range entries {
		if entry.IsFile() {
			files = append(files, entry.File())
		}
	}
	return files, nil
}

// Dirs returns all subdirectories within the directory.
//
// Returns:
//   - []Dir: List of subdirectories within the directory.
//   - error: Any error encountered during the operation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	subdirs, err := dir.Dirs()
func (d Dir) Dirs() ([]Dir, error) {
	entries, err := d.All()
	if err != nil {
		return nil, err
	}
	var dirs []Dir
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Dir())
		}
	}
	return dirs, nil
}

// Delete deletes the directory.
//
// Args:
//   - recursive: If true, deletes the directory recursively along with its contents.
//
// Returns:
//   - error: Any error encountered during the deletion.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	err := dir.Delete(true)
func (d Dir) Delete(recursive bool) error {
	if recursive {
		return os.RemoveAll(d.String())
	}
	return os.Remove(d.String())
}

// DeleteSubFile deletes a file within the directory.
//
// Args:
//   - name: Name of the file to delete.
//
// Returns:
//   - error: Any error encountered during the deletion.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	err := dir.DeleteSubFile("file.txt")
func (d Dir) DeleteSubFile(name string) error {
	file := Join(d.String(), name).File()
	return file.Delete()
}

// DeleteSubDir deletes a subdirectory within the directory.
//
// Args:
//   - name: Name of the subdirectory to delete.
//   - recursive: If true, deletes the subdirectory recursively along with its contents.
//
// Returns:
//   - error: Any error encountered during the deletion.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	err := dir.DeleteSubDir("subdir", true)
func (d Dir) DeleteSubDir(name string, recursive bool) error {
	dir := Join(d.String(), name).Dir()
	return dir.Delete(recursive)
}

// Copy copies the directory and its contents to the specified destination.
//
// Args:
//   - dest: Destination path to copy the directory.
//
// Returns:
//   - error: Any error encountered during the copy operation.
//
// Example:
//
//	dir := Dir{"/path/to/source"}
//	err := dir.Copy("/path/to/destination")
func (d Dir) Copy(dest PathHandler) error {
	all, err := d.All()
	if err != nil {
		return err
	}
	destDir := dest.Dir()
	if err := destDir.CreateIfNotExist(); err != nil {
		return err
	}
	for _, entry := range all {
		if entry.IsDir() {
			if err := entry.Dir().Copy(destDir.PathHandler); err != nil {
				return err
			}
		} else {
			if _, err := entry.File().Copy(destDir); err != nil {
				return err
			}
		}
	}
	return nil
}

// HasDir checks if the directory contains a subdirectory with the given name.
//
// Args:
//   - name: Name of the subdirectory.
//
// Returns:
//   - bool: True if the directory contains the subdirectory, false otherwise.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	exists := dir.HasDir("subdir")
func (d Dir) HasDir(name string) bool {
	path := Join(d.String(), name)
	return path.Exists() && path.IsDir()
}

// HasFile checks if the directory contains a file with the given name.
//
// Args:
//   - name: Name of the file.
//
// Returns:
//   - bool: True if the directory contains the file, false otherwise.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	exists := dir.HasFile("file.txt")
func (d Dir) HasFile(name string) bool {
	path := Join(d.String(), name)
	return path.Exists() && path.IsFile()
}

// Find finds paths matching the specified pattern within the directory.
//
// Args:
//   - match: Pattern to match against.
//   - recursive: If true, searches recursively within subdirectories.
//   - quantity: Maximum number of matching paths to return (0 for all).
//
// Returns:
//   - []PathHandler: List of matching paths.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	paths := dir.Find("*.txt", true, 5)
func (d Dir) Find(match string, recursive bool, quantity int) []PathHandler {
	if quantity == 0 {
		return []PathHandler{}
	} else {
		paths := []PathHandler{}
		d.forEach(recursive, func(p PathHandler) {
			matched, err := filepath.Match(match, p.String())
			if matched {
				paths = append(paths, p)
				quantity--
			}
			if quantity == 0 || err != nil {
				return
			}
		})
		return paths
	}
}

// FindFile finds files matching the specified pattern within the directory.
//
// Args:
//   - match: Pattern to match against.
//   - recursive: If true, searches recursively within subdirectories.
//   - quantity: Maximum number of matching files to return (0 for all).
//
// Returns:
//   - []File: List of matching files.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	files := dir.FindFile("*.txt", true, 5)
func (d Dir) FindFile(match string, recursive bool, quantity int) []File {
	// find the file in the dir
	if quantity == 0 {
		return []File{}
	} else {
		files := []File{}
		d.forEach(recursive, func(p PathHandler) {
			matched, err := filepath.Match(match, p.String())
			if matched && p.IsFile() {
				files = append(files, p.File())
				quantity--
			}
			if quantity == 0 || err != nil {
				return
			}
		})
		return files
	}

}

// FindDir finds subdirectories matching the specified pattern within the directory.
//
// Args:
//   - match: Pattern to match against.
//   - recursive: If true, searches recursively within subdirectories.
//   - quantity: Maximum number of matching directories to return (0 for all).
//
// Returns:
//   - []Dir: List of matching subdirectories.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	subdirs := dir.FindDir("*", true, 5)
func (d Dir) FindDir(match string, recursive bool, quantity int) []Dir {
	// find the dir in the dir
	if quantity == 0 {
		return []Dir{}
	} else {
		dirs := []Dir{}
		d.forEach(recursive, func(p PathHandler) {
			matched, err := filepath.Match(match, p.String())
			if matched && p.IsDir() {
				dirs = append(dirs, p.Dir())
				quantity--
			}
			if quantity == 0 || err != nil {
				return
			}
		})
		return dirs
	}
}

// CreateSubdir creates a subdirectory within the directory.
//
// Args:
//   - name: Name of the subdirectory to create.
//
// Returns:
//   - Dir: Created subdirectory.
//   - error: Any error encountered during the creation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	subdir, err := dir.CreateSubdir("subdir")
func (d Dir) CreateSubdir(name string) (Dir, error) {
	//create the dir inside d
	path := PathHandler(filepath.Join(d.String(), name))
	dir := path.Dir()
	err := dir.CreateIfNotExist()
	if err != nil {
		return Dir{}, err
	}
	return dir, nil
}

// CreateFile creates a file within the directory.
//
// Args:
//   - name: Name of the file to create.
//   - overwrite: If true, overwrites the file if it already exists.
//
// Returns:
//   - File: Created file.
//   - error: Any error encountered during the creation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	file, err := dir.CreateFile("file.txt", true)
func (d Dir) CreateFile(name string, overwrite bool) (File, error) {
	// create the file inside d
	path := PathHandler(filepath.Join(d.String(), name))
	file := path.File()
	err := file.Create(overwrite)
	if err != nil {
		return File{}, err
	}
	return file, nil
}

// CreateFileWithData creates a file with the given data within the directory.
//
// Args:
//   - name: Name of the file to create.
//   - data: Data to write to the file.
//   - overwrite: If true, overwrites the file if it already exists.
//
// Returns:
//   - File: Created file.
//   - error: Any error encountered during the creation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	file, err := dir.CreateFileWithData("file.txt", []byte("Hello"), true)
func (d Dir) CreateFileWithData(name string, data []byte, overwrite bool) (File, error) {
	file := Join(d.String(), name).File()
	err := file.Create(overwrite)
	if err != nil {
		return File{}, err
	}
	err = file.Write(data)
	if err != nil {
		return File{}, err
	}
	return file, nil
}

// CreateFileWithString creates a file with the given string data within the directory.
//
// Args:
//   - name: Name of the file to create.
//   - data: String data to write to the file.
//   - overwrite: If true, overwrites the file if it already exists.
//
// Returns:
//   - File: Created file.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	file := dir.CreateFileWithString("file.txt", "Hello", true)
func (d Dir) CreateFileWithString(name string, data string, overwrite bool) File {
	file := Join(d.String(), name).File()
	file.Create(overwrite)
	file.WriteString(data)
	return file
}

// GetTree returns the directory structure as a tree.
//
// Returns:
//   - DirStructure: Directory structure represented as a tree.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	tree := dir.GetTree()

func (d Dir) GetTree() DirStructure {
	return getTree(d)
}

func getTree(p Dir) DirStructure {
	allEntry, err := p.All()
	if err != nil {
		return DirStructure{}
	}
	tree := DirStructure{}
	for _, entry := range allEntry {
		if entry.IsDir() {
			tree.Dirs[entry.String()] = getTree(entry.Dir())
		} else {
			tree.Files = append(tree.Files, entry.File())
		}
	}
	return tree
}

// GetAllPathExists returns all existing paths within the directory.
//
// Returns:
//   - []PathHandler: List of existing paths.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	paths := dir.GetAllPathExists()
func (d Dir) GetAllPathExists() []PathHandler {
	all, err := d.All()
	if err != nil {
		return []PathHandler{}
	} else {
		paths := []PathHandler{}
		for _, entry := range all {
			paths = append(paths, entry)
			if entry.IsDir() {
				entryDir := entry.Dir()
				entryPaths := entryDir.GetAllPathExists()
				paths = append(paths, entryPaths...)
			}
		}
		return paths
	}
}

// Clear deletes all contents within the directory.
//
// Args:
//   - force: If true, deletes contents even if the directory is not empty.
//
// Returns:
//   - error: Any error encountered during the operation.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	err := dir.Clear(true)
func (d Dir) Clear(force bool) error {
	// clear everything inside the dir but not the dir itself
	if d.IsEmpty() {
		return nil
	} else {
		all_entries, err := d.All()
		if err != nil {
			return err
		}
		for _, entry := range all_entries {
			entry.DeletePath(force)
		}
	}
	return nil
}

// IsEmpty checks if the directory is empty.
//
// Returns:
//   - bool: True if the directory is empty, false otherwise.
//
// Example:
//
//	dir := Dir{"/path/to/directory"}
//	empty := dir.IsEmpty()
func (d Dir) IsEmpty() bool {
	entries, err := d.All()
	if err != nil {
		return false
	}
	return len(entries) == 0
}

// forEach iterates over each path within the directory.
//
// Args:
//   - recursive: If true, iterates recursively over subdirectories.
//   - handler: Handler function to execute for each path.
func (d Dir) forEach(recursive bool, handler func(PathHandler)) {
	entries, err := d.All()
	if err != nil {
		return
	}
	for _, entry := range entries {
		handler(entry)
		if entry.IsDir() && recursive {
			entryDir := entry.Dir()
			entryDir.forEach(recursive, handler)
		}
	}
}
