package gofs

type Dir struct {
	PathHandler
}

type DirStructure struct {
	Dirs  map[string]DirStructure
	Files []File
}

func (d *Dir) CreateIfNotExist() error {

}

func (d *Dir) Files() []File {

}

func (d *Dir) Dirs() []Dir {

}

func (d *Dir) Delete(recursive bool) error {

}

func (d *Dir) DeleteFile(name string) error {

}

func (d *Dir) DeleteDir(name string, recursive bool) error {

}

func (d *Dir) Rename(newName string) error {

}

func (d *Dir) Move(recursive bool) error {

}

func (d *Dir) Copy(recursive bool) error {

}

func (d *Dir) HasDir(name string) bool {

}

func (d *Dir) HasFile(name string) bool {

}

func (d *Dir) Find(match string, recursive bool, quantity int) []PathHandler {

}

func (d *Dir) FindFile(match string, recursive bool, quantity int) []File {

}

func (d *Dir) FindDir(match string, recursive bool, quantity int) []Dir {

}

func (d *Dir) CreateDir(name string) Dir {

}

func (d *Dir) CreateFile(name string, overwrite bool) File {

}

func (d *Dir) CreateFileWithData(name string, data []byte, overwrite bool) File {

}

func (d *Dir) CreateFileWithString(name string, data string, overwrite bool) File {

}

func (d *Dir) GetTree() DirStructure {

}

func (d *Dir) GetAllPathExists() string {

}

func (d *Dir) Clear(force bool) {

}
