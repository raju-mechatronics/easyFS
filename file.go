package gofs

type File struct {
	path string
}

func (F File) String() string {
	return F.path
}

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
