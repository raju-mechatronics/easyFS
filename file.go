package gofs

type File struct {
	PathHandler
}

func NewFile(path PathHandler) File {
	return File{path}
}

func (f *File) Size() {

}

func (f *File) GetMetaData() {

}

func (f *File) Delete() {

}

func (f *File) Rename(newName string) {

}

func (f *File) Move(newPath PathHandler) {

}

func (f *File) Copy(newPath PathHandler) {

}

func (f *File) Create(overwrite bool) {

}

func (f *File) CreateIfNotExists() {

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
