package gofs

import "os"

type File struct {
	PathHandler
}

func NewFile(path PathHandler) File {
	return File{path}
}

func (f *File) Size() (int64, error) {
	// get the file size
	stat, err := os.Stat(f.String())
	return stat.Size(), err
}

func (f *File) GetMetaData() {

}

func (f *File) Delete() error {
	err := os.Remove(f.String())
	return err
}

func (f *File) Rename(newName string) error {
	// rename
	err := os.Rename(f.String(), newName)
	return err
}

func (f *File) Move(newPath PathHandler) {

}

func (f *File) Copy(newPath PathHandler) {

}

func (f *File) Create(overwrite bool) error {
	if file, err := f.IsFile(); file && err != nil {
		if overwrite {
			f.Delete()
		} else {
			return nil
		}
	}
	_, err := os.Create(f.String())
	return err
}

func (f *File) CreateIfNotExists() error {
	return f.Create(false)
}

func (f *File) Read() ([]byte, error) {

}

func (f *File) ReadAll() ([]byte, error) {

}

func (f *File) ReadString() (string, error) {

}

func (f *File) ReadLines() ([]string, error) {

}

func (f *File) Write(data []byte) error {

}

func (f *File) WriteString(data string) error {

}

func (f *File) WriteLines(data []string) error {

}

func (f *File) Append(data []byte) error {

}

func (f *File) AppendString(data string, newLine bool) error {

}

func (f *File) AppendLine(data string) error {

}

func (f *File) ReadStream() (any, error) {

}

func (f *File) WriteStream() (any, error) {

}
