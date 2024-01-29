package gofs

type FS string

func (F FS) Name() {
	
}

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
	DeleteRecursive() error
	Size() int64
}
