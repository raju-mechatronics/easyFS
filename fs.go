package gofs

type FS struct {
	path string
}

func File(path string) IFile {
	return nil
}

func Dir(path string) Directory {
	return nil
}

func FSys(path string) FileSystem {
	return nil
}

type FileSystem interface {
	//common functionality for both file and directory
	Copy(source, destination string) error
	Move(source, destination string) error
	Delete(path string) error
	Rename(path, newName string) error
	Exits(path string) bool
	AbsPath(path string) string
	ParentPath(path string) string
	Name(path string) string
	IsValidPath(path string) bool
	OnSameDir(path1, path2 string) bool
}

type IFile interface {
	Name() string
	Parent() Directory
	Path() string
	Ext() string
	Size() int64
	LastModified() int64
	LastAccessed() int64
	LastChanged() int64
	FullPath() string
	FileSystem() FileSystem
	Delete() error
	Rename(newName string) error
	Move(newParent Directory) error
	Copy(newParent Directory) error
	IsRoot() bool
	IsChildOf(parent Directory) bool
	IsSameAs(other IFile) bool
	IsDescendantOf(ancestor Directory) bool
	IsAncestorOf(descendant Directory) bool
	IsParentOf(child Directory) bool
	IsSiblingOf(sibling IFile) bool
	IsRelativeOf(relative Directory) bool
	IsRelativeOfPath(path string) bool
	Exists() bool
}

type Directory interface {
	Name() string
	Parent() Directory
	Child(name string) Directory
	Children() []Directory
	Files() []IFile
	Path() string
	FullPath() string
	FileSystem() FileSystem
	CreateDirectory(name string) Directory
	CreateFile(name string) IFile
	Delete() error
	DeleteDirectory(name string) error
	DeleteFile(name string) error
	Rename(newName string) error
	Move(newParent Directory) error
	Copy(newParent Directory) error
	IsRoot() bool
	IsChildOf(parent Directory) bool
	IsSameAs(other Directory) bool
	IsDescendantOf(ancestor Directory) bool
	IsAncestorOf(descendant Directory) bool
	IsParentOf(child Directory) bool
	IsSiblingOf(sibling Directory) bool
	IsRelativeOf(relative Directory) bool
	IsRelativeOfPath(path string) bool
	Exists() bool
	DirectoryExists(name string) bool
	FileExists(name string) bool
	DirectoryExistsPath(path string) bool
	FileExistsPath(path string) bool
	DirectoryPath(path string) Directory
	FilePath(path string) IFile
	DirectoryPathCreate(path string) Directory
	FilePathCreate(path string) IFile
	DirectoryPathCreateAll(path string) Directory
	FilePathCreateAll(path string) IFile
	DirectoryPathCreateAllIfNotExists(path string) Directory
	FilePathCreateAllIfNotExists(path string) IFile
	DirectoryPathCreateIfNotExists(path string) Directory
	FilePathCreateIfNotExists(path string) IFile
	DirectoryPathCreateIfNotExistsAll(path string) Directory
	FilePathCreateIfNotExistsAll(path string) IFile
	DirectoryPathCreateAllIfNotExistsAll(path string) Directory
	FilePathCreateAllIfNotExistsAll(path string) IFile
}
