// go test -v futils.go > test.txt

package futils

import (
	"fmt"
	"testing"
)

var baseDir string = "."
var testFolder1 string = baseDir + "testFolder1"
var testFolder2 string = baseDir + "testFolder2"
var testFolder3 string = baseDir + "testFolder3"

var filename1 string = testFolder1 + "/" + "Testfile1.txt"
var filename2 string = testFolder2 + "/" + "Testfile2.txt"

func TestCreateFolders(t *testing.T) {

	err := CreateFolder(testFolder1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	err = CreateFolder(testFolder2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	err = CreateFolder(testFolder3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	PrintFolderRecursively(baseDir)
}

func TestCreateEmptyFiles(t *testing.T) {

	fmt.Printf("filename1=%s\n", filename1)
	err := CreateEmptyFile(filename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("filename2=%s\n", filename2)
	err = CreateEmptyFile(filename2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	PrintFolderRecursively(baseDir)
}

func TestWriteAndReadFile(t *testing.T) {

	fmt.Printf("filename1=%s\n", filename1)
	data := []byte{}
	stringData := "Hello world"
	data = []byte(stringData)
	fmt.Printf("Data:%s\n", string(data[:]))
	err := WriteGobEncodedFile(filename1, &data)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	stringData = ""
	data = []byte(stringData)

	err = ReadGobEncodedFile(filename1, &data)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Data:%s\n", string(data[:]))
}

func TestRenameFile(t *testing.T) {

	newfilename1 := testFolder1+ "/" + "NewTestFile.txt" 
	fmt.Printf("filename1=%s\n", filename1)
	fmt.Printf("newfilename1=%s\n", newfilename1)
	err := RenameFile(filename1, newfilename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	filename1 = newfilename1
	PrintFolderRecursively(baseDir)
}

func TestMoveFile(t *testing.T) {

	newfilename1 := testFolder2 + "/" + "NewTestFile.txt" 
	err := MoveFile(filename1, newfilename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	filename1 = newfilename1
	PrintFolderRecursively(baseDir)
}

func TestDeleteFile(t *testing.T) {

	err := DeleteFile(filename1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	PrintFolderRecursively(baseDir)
}

func TestMoveFolders(t *testing.T) {

	newFolder := testFolder1 + "/" + testFolder2 
	err := MoveFolder(testFolder2, newFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	testFolder2 = newFolder

	newFolder = testFolder1 + "/" + testFolder3 
	err = MoveFolder(testFolder3, newFolder)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	testFolder3 = newFolder
	PrintFolderRecursively(baseDir)
}

func TestDeleteFolders(t *testing.T) {

	err := DeleteFolder(testFolder1)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	err = DeleteFolder(testFolder2)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	err = DeleteFolder(testFolder3)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	PrintFolderRecursively(baseDir)
}
