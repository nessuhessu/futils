package futils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"os"
	"path/filepath"
	"strconv"
)

func FileOrFolderExists(fileOrFolderName string) bool {
	_, err := os.Stat(fileOrFolderName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateEmptyFile(fullFilename string) error {

	newFile, err := os.Create(fullFilename)
	if err != nil {
		return err
	}
	newFile.Close()
	return nil
}

func CreateFile(fullFilename string, buffer *bytes.Buffer) error {

	err := ioutil.WriteFile(fullFilename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func RenameFile(currFullFilename, newFullFilename string) error {

	err := os.Rename(currFullFilename, newFullFilename)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(FullFilename string) error {

    err := os.Remove(FullFilename)
    if err != nil {
		return err
    }
	return nil
}

func CopyFile(SourceFullFilename, targetFullFilename string) error {

	// Open file
	sourceFile, err := os.Open(SourceFullFilename)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
 
	// Create new file
	newFile, err := os.Create(targetFullFilename)
	if err != nil {
		return err
	}
	defer newFile.Close()
 
	_ , err = io.Copy(newFile, sourceFile)
	if err != nil {
		return err
	}	
	return nil
}

func MoveFile(currFullFilename, newFullFilename string) error {

	return RenameFile(currFullFilename, newFullFilename)
}

func WriteGobEncodedFile(fullFilename string, data interface{}) error {
 	
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	// Possible existing file is overwritten
	err = ioutil.WriteFile(fullFilename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadGobEncodedFile(fullFilename string, data interface{}) error {

	raw, err := ioutil.ReadFile(fullFilename)
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

func CopyFolder(sourceFolderName, targerFolderName string) error {

    cmd := exec.Command("cp", "--recursive", sourceFolderName, targerFolderName)
    err := cmd.Run()
	return err
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

func PrintFolderRecursively(fromFolder string) error {
  
	err := filepath.WalkDir(fromFolder, func(fullFileName string, file fs.DirEntry,  err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := file.Info()
		if err!= nil {
			return err
		}

		fmt.Printf("dir: %v: name: %s size: %v, modified: %s\n", fileInfo.IsDir(), fullFileName, fileInfo.Size(), fileInfo.ModTime())

        return nil
    })
    
	if err != nil {
        fmt.Println(err)
    }
	return nil
}

type fileInfo struct {
	IsDir bool
	FullFileName string
	Size string
	Date string
}

func GetFolderContentRecursively(fromFolder string) (map[string]fileInfo, error) {
  
	files := make(map[string]fileInfo)
	err := filepath.WalkDir(fromFolder, func(fullFileName string, file fs.DirEntry,  err error) error {
		if err != nil {
			return err
		}
		
		osFileInfo, err := file.Info()
		if err!= nil {
			return err
		}
		fileInfo := fileInfo{}
		fileInfo.IsDir = osFileInfo.IsDir()
		fileInfo.FullFileName = fullFileName  // Relative path + filename
		fileInfo.Size = strconv.FormatInt(osFileInfo.Size(),10)
		fileInfo.Date = osFileInfo.ModTime().String()
		files[fileInfo.FullFileName] = fileInfo
		return nil
	});
	return files, err
}
