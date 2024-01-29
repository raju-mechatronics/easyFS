package gofs


type FilePrototype interface {
	Name() string
	Parent() Directory
	Path() string
	Ext() string
	Size() int64
	LastModified() int64
	LastAccessed() int64
	LastChanged() int64
	FullPath() string
	FileSystem() FileSystemPrototype
	Delete() error
	Rename(newName string) error
	Move(newParent Directory) error
	Copy(newParent Directory) error
	IsRoot() bool
	IsChildOf(parent Directory) bool
	IsSameAs(other FilePrototype) bool
	IsDescendantOf(ancestor Directory) bool
	IsAncestorOf(descendant Directory) bool
	IsParentOf(child Directory) bool
	IsSiblingOf(sibling FilePrototype) bool
	IsRelativeOf(relative Directory) bool
	IsRelativeOfPath(path string) bool
	Exists() bool
}

// this is only for files. FS will implement this too
type FileProto interface {
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
