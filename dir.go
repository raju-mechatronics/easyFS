package gofs

import (
	"os"
	"path/filepath"
)

type Dir struct {
	PathHandler
}

type DirStructure struct {
	Dirs  map[string]DirStructure
	Files []File
}

func (d *Dir) CreateIfNotExist() error {

	if !d.Exists() {
		return nil
	} else {
		//create dir
		err := os.Mkdir(d.String(), 0777)
		return err
	}
}

func (d *Dir) All() ([]PathHandler, error) {
	//read the dir
	el, err := os.ReadDir(d.String())
	if err != nil {
		return nil, err
	}
	//convert to pathhandler
	var paths []PathHandler
	for _, e := range el {
		//get the file path
		path := d.String() + "/" + e.Name()
		paths = append(paths, PathHandler(path))
	}
	return paths, nil
}
func (d *Dir) Files() ([]File, error) {
	//get all files in dir
	el, err := os.ReadDir(d.String())
	if err != nil {
		return nil, err
	}
	files := []File{}
	for _, e := range el {
		//get the file path
		path := d.String() + "/" + e.Name()
		//check if it is a file
		if e.IsDir() {
			continue
		}
		//create file
		file := NewFile(PathHandler(path))
		files = append(files, file)
	}
	return files, nil
}

func (d *Dir) Dirs() ([]Dir, error) {
	//get all dirs in dir
	el, err := os.ReadDir(d.String())
	if err != nil {
		return nil, err
	}
	dirs := []Dir{}
	for _, e := range el {
		//get the file path
		path := d.String() + "/" + e.Name()
		//check if it is a dir
		if !e.IsDir() {
			continue
		}
		//create dir
		dir := Dir{PathHandler(path)}
		dirs = append(dirs, dir)
	}
	return dirs, nil
}

// delete the dir d
func (d *Dir) Delete(recursive bool) error {
	//delete the dir
	if !recursive {
		err := os.Remove(d.String())
		return err
	} else {
		err := os.RemoveAll(d.String())
		return err
	}
}

func (d *Dir) DeleteSubFile(name string) error {
	isFile, err := d.IsFile()
	if err != nil && isFile {
		return os.Remove(d.String())
	}
	return err
}

func (d *Dir) DeleteSubDir(name string, recursive bool) error {
	//delete the dir
	if !recursive {
		err := os.Remove(filepath.Join(d.String(), name))
		return err
	} else {
		err := os.RemoveAll(filepath.Join(d.String(), name))
		return err
	}
}

// delete anything named name that inside the dir d
func (d *Dir) DeleteAnything(match string, force bool) error {
}

func (d *Dir) Rename(newName string) error {
	//rename the dir
	err := os.Rename(d.String(), newName)
	return err
}

func (d *Dir) Move(recursive bool) error {

}

func (d *Dir) Copy(recursive bool) error {

}

func (d *Dir) HasDir(name string) bool {
	//check if the name exists in d and if it is a dir
	stat, err := os.Stat(filepath.Join(d.String(), name))
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func (d *Dir) HasFile(name string) bool {
	//check if the name exists in d and if it is a file
	stat, err := os.Stat(filepath.Join(d.String(), name))
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func (d *Dir) Find(match string, recursive bool, quantity int) []PathHandler {

}

func (d *Dir) FindFile(match string, recursive bool, quantity int) []File {

}

func (d *Dir) FindDir(match string, recursive bool, quantity int) []Dir {

}

func (d *Dir) CreateDir(name string) Dir {
	//create the dir inside d
	dir := Dir{PathHandler(filepath.Join(d.String(), name))}
	dir.CreateIfNotExist()
	return dir
}

func (d *Dir) CreateFile(name string, overwrite bool) File {
	// create the file inside d
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
