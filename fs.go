package gofs

type FS string

func (F FS) String() string {
	return string(F)
}

// FS will implement the FileSystemPrototype interface for any type of file or directory
type FileSystemPrototype interface {
	Name() string
	IsValidPath() bool
	Abs() FileSystemPrototype
	Rel() FileSystemPrototype
	Parent() FileSystemPrototype
	IsDir() bool
	IsFile() bool
	Exists() bool
	IsChildOf(path FS) bool
	IsAbs() bool
	IsRel() bool
	IsSame(path FS) bool
	IsDescendantOf(path FS) bool
	IsSiblingOf(path FS) bool
	Stat() any        /* Not implemented yet */
	Permissions() any /* Not implemented yet */
	Owner() any       /* Not implemented yet */
	IsReadable() bool
	IsWritable() bool
	IsExecutable() bool
	IsHidden() bool
	Copy(newPath FS) error
	Move(newPath FS) error
	Rename(newName string) error
	Delete() error
	Size() int64
}

// this is only for files. FS will implement this too
type FileProto interface {
	Ext() string
	Read() ([]byte, error)
	ReadAll() ([]byte, error)
	ReadString() (string, error)
	ReadLines() ([]string, error)
	Write(data []byte) error
	WriteAll(data []byte) error
	WriteString(data string) error
	WriteLines(data []string) error
	Append(data []byte) error
	AppendAll(data []byte) error
	AppendString(data string) error
	AppendLines(data []string) error
	ReadStream() (any, error)
	WriteStream() (any, error)
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
