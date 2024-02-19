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
	// move the file to the new path

}

func (f *File) Copy(newPath PathHandler) {
	// copy the file to the new path

}

func (f *File) Create(overwrite bool) error {
	if f.IsFile() || f.Exists() {
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
	// read the file
	file, err := os.Open(f.String())
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data := make([]byte, f.Size())
	_, err = file.Read(data)
	return data, err
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
