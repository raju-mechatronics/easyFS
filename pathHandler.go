package gofs

import (
	"os"
	"path/filepath"
	"strings"
)

// PathHandler represents a file or directory path.
type PathHandler string

// PathInfo represents information about a file or directory.
type PathInfo os.FileInfo

// String returns the string representation of the PathHandler.
func (p PathHandler) String() string {
	return string(p)
}

// IsValidPath checks if the path contains valid directory names.
// Returns true if all directory names in the path are valid, otherwise false.
// Example:
//
//	path := PathHandler("/path/to/directory")
//	isValid := path.IsValidPath()
//	fmt.Println(isValid) // Output: true or false
func (p PathHandler) IsValidPath() bool {
	names := strings.Split(p.String(), string(os.PathSeparator))
	for _, name := range names {
		if !IsValidDirName(name) {
			return false
		}
	}
	return true
}

// Exists checks if the path exists.
// Returns true if the path exists, otherwise false.
// Example:
//
//	path := PathHandler("/path/to/file")
//	exists := path.Exists()
//	fmt.Println(exists) // Output: true or false
func (p PathHandler) Exists() bool {
	_, err := os.Stat(p.String())
	return err == nil
}

// IsDir checks if the path represents a directory.
// Returns true if the path is a directory, otherwise false.
// Example:
//
//	path := PathHandler("/path/to/directory")
//	isDir := path.IsDir()
//	fmt.Println(isDir) // Output: true or false
func (p PathHandler) IsDir() bool {
	info, err := os.Stat(p.String())
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile checks if the path represents a file.
// Returns true if the path is a file, otherwise false.
// Example:
//
//	path := PathHandler("/path/to/file")
//	isFile := path.IsFile()
//	fmt.Println(isFile) // Output: true or false
func (p PathHandler) IsFile() bool {
	info, err := os.Stat(p.String())
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsSymlink checks if the path represents a symbolic link.
// Returns true if the path is a symbolic link, otherwise false.
// Example:
//
//	path := PathHandler("/path/to/symlink")
//	isSymlink := path.IsSymlink()
//	fmt.Println(isSymlink) // Output: true or false
func (p PathHandler) IsSymlink() bool {
	info, err := os.Lstat(p.String())
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

// Stat returns information about the path.
// Example:
//
//	path := PathHandler("/path/to/file")
//	info, err := path.Stat()
//	if err == nil {
//	    fmt.Println(info.Name()) // Output: file
//	}
func (p PathHandler) Stat() (PathInfo, error) {
	return os.Stat(p.String())
}

// IsAbs reports whether the path is absolute.
func (p PathHandler) IsAbs() bool {
	path := string(p)
	return filepath.IsAbs(path)
}

// Resolve resolves the path to its absolute form.
// Returns the resolved path and any error encountered.
func (p *PathHandler) Resolve() (string, error) {
	rPath, err := filepath.Abs(p.String())
	if err == nil {
		*p = PathHandler(rPath)
	}
	return rPath, err
}

// Abs returns the absolute representation of the path.
// Returns the absolute path and any error encountered.
func (p PathHandler) Abs() (string, error) {
	return filepath.Abs(p.String())
}

// Name returns the base name of the file or directory.
func (p PathHandler) Name() string {
	return filepath.Base(p.String())
}

// Parent returns the parent directory of the path.
func (p PathHandler) Parent() Dir {
	parent := filepath.Dir(p.String())
	return Dir{PathHandler(parent)}
}

// Ext returns the file extension of the file.
func (p PathHandler) Ext() string {
	return filepath.Ext(p.String())
}

// Join joins the current path with the given elements.
func (p PathHandler) Join(elem ...string) PathHandler {
	newPath := Join(p.String(), elem...)
	return PathHandler(newPath)
}

// IsRel reports whether the path is relative.
func (p PathHandler) IsRel() bool {
	return !p.IsAbs()
}

// IsSameDir checks if the path is in the same directory as the other path.
func (p PathHandler) IsSameDir(other PathHandler) bool {
	pabs, err := p.Abs()
	oabs, oerr := other.Abs()

	if err == nil || oerr == nil {
		return false
	}
	return pabs == oabs
}

// IsSiblingOf checks if the path is a sibling of the other path.
func (p PathHandler) IsSiblingOf(other PathHandler) bool {
	return p.Parent() == other.Parent()
}

// IsDescendantOf checks if the path is a descendant of the other path.
func (p PathHandler) IsDescendantOf(other PathHandler) bool {
	f, e1 := p.Abs()
	o, e2 := other.Abs()
	if e1 != nil || e2 != nil {
		return false
	}
	return strings.HasPrefix(f, o)
}

// DeletePath deletes the file or directory at the path.
// If force is true, it deletes the path and any children.
// Returns any error encountered.
func (p PathHandler) DeletePath(force bool) error {
	if force {
		return os.RemoveAll(p.String())
	} else {
		return os.Remove(p.String())
	}
}

// File returns a File instance for the path.
func (p PathHandler) File() File {
	return File{p}
}

// Dir returns a Dir instance for the path.
func (p PathHandler) Dir() Dir {
	d := Dir{p}
	return d
}

// Rename renames the directory represented by the path.
// newName is the new name for the directory.
// Returns any error encountered.
func (p *PathHandler) Rename(newName string) error {
	// Rename the directory
	err := os.Rename(p.String(), string(Join(p.Parent().String(), newName)))
	if err == nil {
		newHandler := Join(p.Parent().String(), newName)
		*p = newHandler
	}
	return err
}

// Move moves the directory represented by the path to a new location.
// newPath is the new path for the directory.
// Returns any error encountered.
func (p *PathHandler) Move(newPath PathHandler) error {
	// Get the name
	name := p.Name()
	// Move the directory
	dest := Join(newPath.String(), name)
	err := os.Rename(p.String(), dest.String())
	if err == nil {
		*p = dest
	}
	return err
}

// SetPerm sets the permission bits for the path.
// perm is the permission bits to set.
// Returns any error encountered.
func (p PathHandler) SetPerm(perm os.FileMode) error {
	return os.Chmod(p.String(), perm)
}
