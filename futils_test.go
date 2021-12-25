// go test -v futils.go > test.txt

package futils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseFolder string = "../test/"
var testFolder1 string = baseFolder + "testFolder1"
var testFolder2 string = baseFolder + "testFolder2"
var testFolder3 string = baseFolder + "testFolder3"

var filename1 string = testFolder1 + "/" + "Testfile1.txt"
var filename2 string = testFolder2 + "/" + "Testfile2.txt"
var filename3 string = testFolder3 + "/" + "Testfile3.txt"

func TestCreateFolders(t *testing.T) {

	assert.Equal(t, false, FileOrFolderExists(baseFolder))
	assert.Equal(t, false, FileOrFolderExists(testFolder1))
	assert.Equal(t, false, FileOrFolderExists(testFolder2))
	assert.Equal(t, false, FileOrFolderExists(testFolder3))
	fmt.Printf("baseFolder exists: %v\n", FileOrFolderExists(baseFolder))
	fmt.Printf("Create folder: %s\n", baseFolder)
	err := CreateFolder(baseFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("baseFolder exists: %v\n", FileOrFolderExists(baseFolder))

	fmt.Printf("Create folder: %s\n", testFolder1)
	err = CreateFolder(testFolder1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Create folder: %s\n", testFolder2)
	err = CreateFolder(testFolder2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Create folder: %s\n", testFolder3)
	err = CreateFolder(testFolder3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	assert.Equal(t, true, FileOrFolderExists(baseFolder))
	assert.DirExists(t, testFolder1)
	assert.DirExists(t, testFolder2)
	assert.DirExists(t, testFolder3)
	PrintFoldersContent()
}

func TestCreateEmptyFiles(t *testing.T) {

	assert.Equal(t, false, FileOrFolderExists(filename1))
	assert.Equal(t, false, FileOrFolderExists(filename2))
	assert.Equal(t, false, FileOrFolderExists(filename3))

	fmt.Printf("Create empty file: %s\n", filename1)
	err := CreateEmptyFile(filename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Create empty file: %s\n", filename2)
	err = CreateEmptyFile(filename2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	
	fmt.Printf("Create empty file: %s\n", filename3)
	err = CreateEmptyFile(filename3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	assert.FileExists(t, filename1)
	assert.FileExists(t, filename2)
	assert.FileExists(t, filename3)
	PrintFoldersContent()
}

func TestWriteAndReadFile(t *testing.T) {

	assert.Equal(t, true, FileOrFolderExists(filename1))

	fmt.Printf("Write file: %s\n", filename1)
	data := []byte{}
	stringData := "Hello world"
	data = []byte(stringData)
	fmt.Printf("Write data: %s\n", string(data[:]))
	err := WriteGobEncodedFile(filename1, &data)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	stringData = ""
	data = []byte(stringData)

	fmt.Printf("Read file: %s\n", filename1)
	err = ReadGobEncodedFile(filename1, &data)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Read data: %s\n", string(data[:]))

	assert.FileExists(t, filename1)
	PrintFoldersContent()
}

func TestRenameFile(t *testing.T) {

	newfilename1 := testFolder1 + "/" + "NewTestFile.txt" 
	assert.Equal(t, true, FileOrFolderExists(filename1))
	assert.Equal(t, false, FileOrFolderExists(newfilename1))

	fmt.Printf("Rename file: filename: %s, newfilename: %s\n", filename1, newfilename1)
	err := RenameFile(filename1, newfilename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	filename1 = newfilename1

	assert.FileExists(t, filename1)
	PrintFoldersContent()
}

func TestCopyFile(t *testing.T) {

	newfilename3 := testFolder2 + "/" + "filename3" 
	assert.Equal(t, true, FileOrFolderExists(filename3))
	assert.Equal(t, false, FileOrFolderExists(newfilename3))

	fmt.Printf("Copy file: sourceFilename: %s, targetFilename: %s\n", filename1, newfilename3)
	err := MoveFile(filename3, newfilename3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	filename3 = newfilename3

	assert.FileExists(t, filename3)
	PrintFoldersContent()
}
func TestMoveFile(t *testing.T) {

	newfilename1 := testFolder2 + "/" + "NewTestFile.txt" 
	assert.Equal(t, true, FileOrFolderExists(filename1))
	assert.Equal(t, false, FileOrFolderExists(newfilename1))

	fmt.Printf("Move file: sourceFilename: %s, targetFilename: %s\n", filename1, newfilename1)
	err := MoveFile(filename1, newfilename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	filename1 = newfilename1

	assert.FileExists(t, filename1)
	PrintFoldersContent()
}

func TestDeleteFile(t *testing.T) {

	assert.FileExists(t, filename1)

	fmt.Printf("Delete file: %s\n", filename1)
	err := DeleteFile(filename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	assert.Equal(t, false, FileOrFolderExists(filename1))
	PrintFoldersContent()
}

func TestMoveFolders(t *testing.T) {

	newFolder := testFolder1 + "/testFolder2" 
	assert.DirExists(t, testFolder2)
	assert.Equal(t, false, FileOrFolderExists(newFolder))

	fmt.Printf("Move folder: sourceFolder: %s, targetFolder: %s\n", testFolder2, newFolder)
	err := MoveFolder(testFolder2, newFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	testFolder2 = newFolder

	newFolder = testFolder1 + "/testFolder3" 
	assert.DirExists(t, testFolder3)
	assert.Equal(t, false, FileOrFolderExists(newFolder))

	fmt.Printf("Move folder: sourceFolder: %s, targetFolder: %s\n", testFolder3, newFolder)
	err = MoveFolder(testFolder3, newFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	testFolder3 = newFolder

	assert.DirExists(t, testFolder2)
	assert.DirExists(t, testFolder3)
	PrintFoldersContent()
}

func TestCopyFolder(t *testing.T) {

	newFolder := testFolder3 + "/testFolder2copy" 
	assert.DirExists(t, testFolder2)
	assert.Equal(t, false, FileOrFolderExists(newFolder))

	fmt.Printf("Move folder: sourceFolder: %s, targetFolder: %s\n", testFolder2, newFolder)
	err := MoveFolder(testFolder2, newFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	testFolder2 = newFolder

	assert.DirExists(t, testFolder2)
	PrintFoldersContent()
}

func TestGetFolderContentRecursively(t *testing.T) {

	content, err := GetFolderContentRecursively(baseFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	for _, fileInfo := range content {
		fmt.Printf("dir: %v: name: %s size: %v, modified: %s\n", fileInfo.IsDir, fileInfo.FullFileName, fileInfo.Size, fileInfo.Date)

	}
}

func TestDeleteFolders(t *testing.T) {

	assert.DirExists(t, baseFolder)
	assert.DirExists(t, testFolder1)
	assert.DirExists(t, testFolder2)
	assert.DirExists(t, testFolder3)

	fmt.Printf("Delete folder: %s ", testFolder1)
	err := DeleteFolder(testFolder1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Delete folder: %s ", testFolder2)
	err = DeleteFolder(testFolder2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Delete folder: %s ", testFolder3)
	err = DeleteFolder(testFolder3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Delete folder: %s ", baseFolder)
	err = DeleteFolder(baseFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	assert.Equal(t, false, FileOrFolderExists(testFolder1))
	assert.Equal(t, false, FileOrFolderExists(testFolder2))
	assert.Equal(t, false, FileOrFolderExists(testFolder3))
	assert.Equal(t, false, FileOrFolderExists(baseFolder))
	PrintFoldersContent()
}

func PrintFoldersContent() {
	fmt.Println()
	fmt.Println("Folders content:")
	PrintFolderRecursively(baseFolder)	
	fmt.Println()
}