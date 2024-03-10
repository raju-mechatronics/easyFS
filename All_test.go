package easyFS

import (
	"fmt"
	"math/rand"
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
	resolve, err := file.Resolve()
	if err != nil {
		return
	}
	fmt.Println("if resolved then passed else did not passed", resolve)
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

	//test rename
	nP := PathHandler("testpath").Dir()
	nP.CreateIfNotExist()
	dir1, err1 := nP.CreateSubdir("testdir1")
	dir2, err2 := nP.CreateSubdir("testdir2")
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
	}
	dir1.Rename("testdir3")
	if dir1.String() == "testdir3" && !dir1.Exists() {
		t.Error("Rename failed")
	}
	//move dir1 to dir2
	dir1.Move(dir2.PathHandler)
	if dir1.Parent().String() != dir2.String() {
		t.Error("Move failed")
	}

	//change permission
	errPerm := dir2.SetPerm(0777)
	if errPerm != nil {
		t.Error(errPerm)
	}
	//remove nP
	err = nP.Delete(true)
	if err != nil {
		t.Error(err)
	}
}

func TestDir(t *testing.T) {
	dir := NewDir("testdir")
	//delete dir
	dir.Delete(true)
	// dir.exists() false
	if dir.Exists() {
		t.Error("Dir should not exist")
	}
	dir.CreateIfNotExist()
	// dir.exists() true
	if !dir.Exists() {
		t.Error("Dir should exist")
	}
	//dir.isDir() true
	if !dir.IsDir() {
		t.Error("Dir should be a directory")
	}
	// dir.createdir
	dir1, err1 := dir.CreateSubdir("testdir1")
	if err1 != nil {
		t.Error("failed create subdir:", err1)
	}
	dir2, err2 := dir.CreateSubdir("testdir2")
	if err2 != nil {
		t.Error("failed create subdir:", err2)
	}
	dir3, err3 := dir.CreateSubdir("testdir3")
	if err3 != nil {
		t.Error("failed create subdir:", err3)
	}
	file1, ferr1 := dir.CreateFile("testfile1.ext", false)
	if ferr1 != nil {
		t.Error("failed create file:", ferr1)
	}
	file2, ferr2 := dir.CreateFile("testfile2.ext", false)
	if ferr2 != nil {
		t.Error("failed create file:", ferr2)
	}
	file3, ferr3 := dir.CreateFile("testfile3.ext", false)
	if ferr3 != nil {
		t.Error("failed create file:", ferr3)
	}
	fmt.Println(dir3, file1, file2, file3)

	all, err := dir.All()
	if err != nil {
		t.Error("Testing All failed. got:", err)
	}
	if len(all) != 6 {
		t.Error("Testing All failed. got:", len(all))
	}
	// test file
	files, err := dir.Files()
	if err != nil {
		t.Error("Testing Files failed. got:", err)
	}
	if len(files) != 3 {
		t.Error("Testing Files failed. got:", len(files))
	}
	// test dir
	dirs, err := dir.Dirs()
	if err != nil {
		t.Error("Testing Dirs failed. got:", err)
	}
	if len(dirs) != 3 {
		t.Error("Testing Dirs failed. got:", len(dirs))
	}
	// test delete
	err = dir1.Delete(true)
	if err != nil {
		t.Error("Testing delete failed. got:", err)
	}
	if dir1.Exists() {
		t.Error("Testing delete failed. got:", err)
	}
	// test delete subdir
	err = dir.DeleteSubDir("testdir2", true)
	if err != nil {
		t.Error("Testing delete subdir failed. got:", err)
	}
	if dir2.Exists() {
		t.Error("Testing delete subdir failed. got:", err)
	}
	// test delete file
	err = dir.DeleteSubFile("testfile3.ext")
	if err != nil {
		t.Error("Testing delete file failed. got:", err)
	}
	if file3.Exists() {
		t.Error("Testing delete file failed. got:", err)

	}

	// test hasdir
	if !dir.HasDir("testdir3") {
		t.Error("Testing hasdir failed. got:", err)
	}
	// test hasfile
	if !dir.HasFile("testfile2.ext") {
		t.Error("Testing hasfile failed. got:", err)
	}

}

// func to generate random string
func randString(n int) string {
	var letterRunes = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		if i%100 == 0 {
			b[i] = '\n'
			continue
		}
		randomInt := rand.Int() % len(letterRunes)
		b[i] = letterRunes[randomInt]
	}
	return string(b)
}

func TestFile(t *testing.T) {
	dir := NewDir("testdir")
	//delete dir
	dir.Delete(true)
	fileName := NewFile("testfile.ext")
	dir.CreateIfNotExist()
	file := dir.Join(fileName.Name()).File()
	fmt.Println("file:", file.String())
	err := file.CreateIfNotExists()
	if err != nil {
		t.Error(err)
	}
	random_string := randString(100000)
	err = file.WriteString(random_string)
	if err != nil {
		t.Error("write failed", err)
	}

	// file.exists() true
	size, err := file.Size()
	if err != nil {
		t.Error(err)
	}
	if size <= 0 {
		t.Error("File size should be greater than 0")
	}

	// file read
	data, err := file.ReadString()

	if err != nil {
		t.Error("Reading Failed", err)
	}

	subDir := dir.Join("subdir").Dir()
	subDir.CreateIfNotExist()

	if data != random_string {
		t.Error("ReadString failed")
	}

	file.Copy(subDir)

	data, err = file.ReadString()
	if err != nil {
		t.Error("Reading Failed", err)
	}

	if data != random_string {
		t.Error("Copy data is not matched with original failed")
	}

	// file readline
	reader, _, err := file.ChunkReader(100)

	if err != nil {
		t.Error("Reading Failed", err)
	}
	i := 0
	for subData, err, finished := reader(); !finished && err == nil; subData, err, finished = reader() {
		if string(subData) != random_string[i*100:i*100+100] {
			t.Error("line data is not matched with original failed got:", string(subData), "expected:", random_string[i:i+100])
		}
		i++
	}
	if err != nil {
		t.Error("Reading Failed", err)
	}
	//delete dir
	dir.Delete(true)
}
