package gofs

type Directory interface {
	Name() string
	Parent() Directory
	Child(name string) Directory
	Children() []Directory
	Files() []FilePrototype
	Path() string
	FullPath() string
	FileSystem() FileSystemPrototype
	CreateDirectory(name string) Directory
	CreateFile(name string) FilePrototype
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
	FilePath(path string) FilePrototype
	DirectoryPathCreate(path string) Directory
	FilePathCreate(path string) FilePrototype
	DirectoryPathCreateAll(path string) Directory
	FilePathCreateAll(path string) FilePrototype
	DirectoryPathCreateAllIfNotExists(path string) Directory
	FilePathCreateAllIfNotExists(path string) FilePrototype
	DirectoryPathCreateIfNotExists(path string) Directory
	FilePathCreateIfNotExists(path string) FilePrototype
	DirectoryPathCreateIfNotExistsAll(path string) Directory
	FilePathCreateIfNotExistsAll(path string) FilePrototype
	DirectoryPathCreateAllIfNotExistsAll(path string) Directory
	FilePathCreateAllIfNotExistsAll(path string) FilePrototype
}

// this is only for directories. FS will implement this too
type DirectoryProto interface {
	GetFiles() []FileProto
	GetFilesFiltered() []FileProto
	GetFile(name FS) FileProto
	GetDirectories() []DirectoryProto
	GetDirectory(name FS) DirectoryProto
	GetDirectoriesFiltered() []DirectoryProto
	GetAll() []FS
	GetRecursiveFiles() []FileProto
	GetRecursiveDirectories() []DirectoryProto
	GetAllRecursive() []FileProto
	GetFilesRecursiveFiltered(filter func(FileProto) bool) []FileProto
	GetDirectoriesRecursiveFiltered(filter func(DirectoryProto) bool) []DirectoryProto
	GetRecursiveFiltered(filter func(FS) bool) []FS
	CreateRecursiveFolder(path FS) error
	CreateFolder(path FS) error
	CreateFile(path FS) error
	CreateFileWithData(path FS, data []byte) error
	CreateFileWithString(path FS, data string) error
	CreateFileWithLines(path FS, data []string) error
}