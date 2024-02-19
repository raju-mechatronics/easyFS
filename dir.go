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
	if d.Exists() && d.IsDir() {
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
	el, err := d.All()
	if err != nil {
		return nil, err
	}
	files := []File{}
	for _, e := range el {
		if e.IsDir() {
			continue
		}
		files = append(files, e.File())
	}
	return files, nil
}

func (d *Dir) Dirs() ([]Dir, error) {
	//get all dirs in dir
	el, err := d.All()
	if err != nil {
		return nil, err
	}
	dirs := []Dir{}
	for _, e := range el {
		if e.IsFile() {
			continue
		}
		dirs = append(dirs, e.Dir())
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
	filePath := Join(d.String(), name)
	file := filePath.File()
	if file.Exists() && file.IsFile() {
		return file.Delete()
	} else {
		return os.ErrNotExist
	}
}

func (d *Dir) DeleteSubDir(name string, recursive bool) error {
	dirPath := Join(d.String(), name)
	dir := dirPath.Dir()
	if dir.Exists() && dir.IsDir() {
		return dir.Delete(recursive)
	} else {
		return os.ErrNotExist
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

func (d *Dir) Move(newPath PathHandler, recursive bool) error {

}

func (d *Dir) Copy(recursive bool) error {
	//copy the dir

}

func (d *Dir) HasDir(name string) bool {
	//check if the name exists in d and if it is a dir
	path := Join(d.String(), name)
	if path.Exists() && path.IsDir() {
		return true
	} else {
		return false
	}
}

func (d *Dir) HasFile(name string) bool {
	//check if the name exists in d and if it is a file
	path := Join(d.String(), name)
	if path.Exists() && path.IsFile() {
		return true
	} else {
		return false
	}
}

func (d *Dir) Find(match string, recursive bool, quantity int) []PathHandler {

}

func (d *Dir) FindFile(match string, recursive bool, quantity int) []File {

}

func (d *Dir) FindDir(match string, recursive bool, quantity int) []Dir {

}

func (d *Dir) CreateDir(name string) Dir {
	//create the dir inside d
	path := PathHandler(filepath.Join(d.String(), name))
	dir := path.Dir()
	dir.CreateIfNotExist()
	return dir
}

func (d *Dir) CreateFile(name string, overwrite bool) (File, error) {
	// create the file inside d
	path := PathHandler(filepath.Join(d.String(), name))
	file := path.File()
	err := file.Create(overwrite)
	if err != nil {
		return File{}, err
	}
	return file, nil
}

func (d *Dir) CreateFileWithData(name string, data []byte, overwrite bool) File {
	// create the file inside d
	file := NewFile(PathHandler(filepath.Join(d.String(), name)))
	err := file.Create(overwrite)
	if err != nil {
		return File{}
	}
}

func (d *Dir) CreateFileWithString(name string, data string, overwrite bool) File {

}

func getTree(p Dir) DirStructure {
	allEntry, err := p.All()
	if err != nil {
		return DirStructure{}
	}
	tree := DirStructure{}
	for _, entry := range allEntry {
		if entry.IsDir() {
			tree.Dirs[entry.String()] = getTree(entry.Dir())
		} else {
			tree.Files = append(tree.Files, entry.File())
		}
	}
	return tree
}

func (d *Dir) GetTree() DirStructure {
	return getTree(*d)
}

func (d *Dir) GetAllPathExists() string {

}

func (d *Dir) Clear(force bool) error {
	// clear everything inside the dir but not the dir itself
	if d.IsEmpty() {
		return nil
	} else {
		all_entries, err := d.All()
		if err != nil {
			return err
		}
		for _, entry := range all_entries {
			entry.DeletePath(force)
		}
	}
	return nil
}

func (d *Dir) IsEmpty() bool {
	// check if the dir is empty
	// if it is empty return true
	// else return false
	entries, err := d.All()
	if err != nil {
		return false
	}
	return len(entries) == 0
}
