package futils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateEmptyFile(filename string) error {

	buffer := new(bytes.Buffer)
	err := ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(filename string, buffer *bytes.Buffer) error {

	err := ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func RenameFile(currFilename, newFilename string) error {

	err := os.Rename(currFilename, newFilename)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(filename string) error {

    err := os.Remove(filename)
    if err != nil {
		return err
    }
	return nil
}

func MoveFile(currFilename, newFilename string) error {

	return RenameFile(currFilename, newFilename)
}
/*
func WriteDataInFile(filename string, data interface{}) error {

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filename string, data interface{}) error {
 
	raw, err := ioutil.ReadFile(filename)
	if err!= nil {
		return err
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err!= nil {
		return err
	}
	return nil
}
*/
func WriteGobEncodedFile(filename string, data interface{}) error {
 	
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadGobEncodedFile(filename string, data interface{}) error {

	raw, err := ioutil.ReadFile(filename)
	if err!= nil {
		return err
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err!= nil {
		return err
	}
	return nil
}

func CreateFolder(folderName string) error {

	err := os.Mkdir(folderName, 0755)
    if err != nil {
        return err
    }
	return nil
}

func DeleteFolder(foldername string) error {

	// Delete folder and it's content
	err := os.RemoveAll(foldername)
	if err != nil {
	  return err
	}
	return nil
}

func MoveFolder(currFoldername, newFoldername string) error {

	return RenameFile(currFoldername, newFoldername)
}

func PrintFolder(folder string)  error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	return nil
}
func PrintFolderRecursively(folder string) error {
  
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
            return err
        }
        fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
        return nil
    })
    
	if err != nil {
        fmt.Println(err)
    }
	return nil
}