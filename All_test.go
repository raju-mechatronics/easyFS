package gofs

import (
	"fmt"
	"testing"
)

func TestPathHandler(t *testing.T) {
	p := PathHandler("testpath")
	p.DeletePath(true)
	// p.exists() false
	if p.Exists() {
		t.Error("Path should not exist")
	}
	p.Dir().CreateIfNotExist()
	// p.exists() true
	if !p.Exists() {
		t.Error("Path should exist")
	}
	//p.isDir() true
	if !p.IsDir() {
		t.Error("Path should be a directory")
	}
	//p.isFile() false
	if p.IsFile() {
		t.Error("Path should not be a file")
	}
	dir := p.Dir()
	file, err := dir.CreateFile("testfile.ext", false)
	if err != nil {
		t.Error(err)
	}
	//p.isFile() true
	if !file.IsFile() {
		t.Error("Path should be a file")
	}

	//p.isSymlink() false
	if file.IsSymlink() {
		t.Error("Path should not be a symlink")
	}
	//p.isSymlink() false
	if dir.IsSymlink() {
		t.Error("Path should not be a symlink")
	}

	//p.isAbsolute() false
	if p.IsAbs() {
		t.Error("Path should not be absolute got:", p.String())
	}
	absPath, err := file.Abs()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("if abs then passed else did not passed", absPath)
	// p.resolve()
	file.Resolve()
	// p.isAbsolute() true
	if !file.IsAbs() {
		t.Error("Path should be absolute got:", file.String())
	}

	// p.stat()
	info, err := file.Stat()
	if err != nil {
		t.Error(err)
	}
	// p.stat() != nil
	if info == nil {
		t.Error("Stat should not be nil")
	}
	fmt.Println("Stat", info)
	//Name, Ext, Parent test
	if file.Name() != "testfile.ext" {
		t.Error("Name should be testfile.ext got:", file.Name())
	}
	if file.Ext() != ".ext" {
		t.Error("Ext should be .ext got:", file.Ext())
	}
	if file.Parent().Name() != dir.String() {
		t.Error("Parent should be", dir.String(), "got:", file.Parent().Name())
	}

	err = dir.Delete(true)
	if err != nil {
		t.Error(err)
	}
	// p.exists() false
	if p.Exists() {
		t.Error("Path should not exist")
	}
}
