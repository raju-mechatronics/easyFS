package gofs

import (
	"os"
	"path/filepath"
	"strings"
)

// PathHandler represents a file path.
type PathHandler string

type PathInfo os.FileInfo

// String returns the string representation of the path.
func (p PathHandler) String() string {
	return string(p)
}

// IsValidPath checks if the path string pattern is valid.
// Example:
//
//	path := gofs.PathHandler("/path/to/directory")
//	isValid := path.IsValidPath() // true
func (p *PathHandler) IsValidPath() bool {
	names := strings.Split(p.String(), string(os.PathSeparator))
	for _, name := range names {
		if !IsValidDirName(name) {
			return false
		}
	}
	return true
}

// Exists checks if the path exists.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	exists := path.Exists() // true if the file exists, otherwise false
func (p *PathHandler) Exists() bool {
	_, err := os.Stat(p.String())
	return err == nil
}

// IsDir checks if the path is a directory.
// Example:
//
//	path := gofs.PathHandler("/path/to/directory")
//	isDir, err := path.IsDir() // true, nil if it's a directory, otherwise false, error
func (p *PathHandler) IsDir() bool {
	info, err := os.Stat(p.String())
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile checks if the path is a file.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	isFile, err := path.IsFile() // true, nil if it's a file, otherwise false, error
func (p *PathHandler) IsFile() bool {
	info, err := os.Stat(p.String())
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsSymlink checks if the path is a symlink.
// Example:
//
//	path := gofs.PathHandler("/path/to/symlink")
//	isSymlink, err := path.IsSymlink() // true, nil if it's a symlink, otherwise false, error
func (p *PathHandler) IsSymlink() (bool, error) {
	info, err := os.Lstat(p.String())
	if err != nil {
		return false, err
	}
	return info.Mode()&os.ModeSymlink != 0, nil
}

// Stat returns the FileInfo structure describing the file.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	fileInfo, err := path.Stat() // returns FileInfo, nil if successful, otherwise error
func (p *PathHandler) Stat() (PathInfo, error) {
	return os.Stat(p.String())
}

// Lstat returns the FileInfo structure describing the file, with symlink information if it is a symlink.
// Example:
//
//	path := gofs.PathHandler("/path/to/symlink")
//	fileInfo, err := path.Lstat() // returns FileInfo, nil if successful, otherwise error
func (p *PathHandler) Lstat() (PathInfo, error) {
	return os.Lstat(p.String())
}

// IsAbs checks if the path is absolute or relative.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	isAbsolute := path.IsAbs() // true if it's an absolute path, otherwise false
func (p *PathHandler) IsAbs() bool {
	path := string(*p)
	return filepath.IsAbs(path)
}

// Resolve resolves the path to an absolute path.
// Example:
//
//	path := gofs.PathHandler("./path/to/file.txt")
//	resolvedPath, err := path.Resolve() // returns the resolved absolute path, nil if successful, otherwise error
func (p *PathHandler) Resolve() (string, error) {
	return filepath.Abs(p.String())
}

// Abs returns the absolute path of the path.
// Example:
//
//	path := gofs.PathHandler("./path/to/file.txt")
//	absPath, err := path.Abs() // returns the absolute path, nil if successful, otherwise error
func (p *PathHandler) Abs() (string, error) {
	return filepath.Abs(p.String())
}

// Name returns the last element of the path.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	name := path.Name() // returns "file.txt"
func (p *PathHandler) Name() string {
	return filepath.Base(p.String())
}

// Parent returns the parent directory of the path.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	parent := path.Parent() // returns "/path/to"
func (p *PathHandler) Parent() string {
	return filepath.Dir(p.String())
}

// Ext returns the file extension of the path.
// Example:
//
//	path := gofs.PathHandler("/path/to/file.txt")
//	ext := path.Ext() // returns ".txt"
func (p *PathHandler) Ext() string {
	return filepath.Ext(p.String())
}

// Join joins the path with additional elements.
// Example:
//
//	path := gofs.PathHandler("/path/to")
//	newPath := path.Join("file.txt") // returns "/path/to/file.txt"
func (p *PathHandler) Join(elem ...string) string {
	return filepath.Join(append([]string{p.String()}, elem...)...)
}

// IsRel checks if the path is relative.
// Example:
//
//	path := gofs.PathHandler("./path/to/file.txt")
//	isRelative := path.IsRel() // true if it's a relative path, otherwise false
func (p *PathHandler) IsRel() bool {
	return !p.IsAbs()
}

// IsSameDir checks if the path is the same directory as another path.
// Example:
//
//	path1 := gofs.PathHandler("/path/to/directory1")
//	path2 := gofs.PathHandler("/path/to/directory2")
//	isSameDir := path1.IsSameDir(path2) // true if they're the same directory, otherwise false
func (p *PathHandler) IsSameDir(other PathHandler) bool {
	pabs, err := p.Abs()
	oabs, oerr := other.Abs()

	if err == nil || oerr == nil {
		return false
	}
	return pabs == oabs
}

// IsSiblingOf checks if the path is a sibling of another path.
// Example:
//
//	path1 := gofs.PathHandler("/path/to/directory1")
//	path2 := gofs.PathHandler("/path/to/directory2/subdir")
//	isSibling := path1.IsSiblingOf(path2) // true if they're siblings, otherwise false
func (p *PathHandler) IsSiblingOf(other PathHandler) bool {
	return p.Parent() == other.Parent()
}

// IsDescendantOf checks if the path is a descendant of another path.
// Example:
//
//	path1 := gofs.PathHandler("/path/to/directory1/subdir")
//	path2 := gofs.PathHandler("/path/to/directory2")
//	isDescendant := path1.IsDescendantOf(path2) // true if path1 is a descendant of path2, otherwise false
func (p *PathHandler) IsDescendantOf(other PathHandler) bool {
	f, e1 := p.Abs()
	o, e2 := other.Abs()
	if e1 != nil || e2 != nil {
		return false
	}
	return strings.HasPrefix(f, o)
}

// Delete the path. if force is true it will delete recursive
func (p *PathHandler) DeletePath(force bool) error {
	if force {
		return os.RemoveAll(p.String())
	} else {
		return os.Remove(p.String())
	}
}

// File
func (p *PathHandler) File() File {
	return File{*p}
}

// Dir
func (p *PathHandler) Dir() Dir {
	return Dir{*p}
}
