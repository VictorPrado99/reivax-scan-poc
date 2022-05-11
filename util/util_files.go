package util

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type FileWrapper interface {
	GetFile() fs.File
	GetFileInfo() fs.FileInfo
	GetPath() string
}

type DefaultFileWrapper struct {
	File     fs.File
	FileInfo fs.FileInfo
	Path     string
}

func (w *DefaultFileWrapper) GetFile() fs.File {
	return w.File
}

func (w *DefaultFileWrapper) GetFileInfo() fs.FileInfo {
	return w.FileInfo
}

func (w *DefaultFileWrapper) GetPath() string {
	return w.Path
}

func GetFiles(directory string, libRegEx *regexp.Regexp) *[]FileWrapper {
	var fileList []FileWrapper

	errWalk := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if libRegEx.MatchString(info.Name()) {
			file, openError := os.Open(directory + path)
			if openError != nil {
				defer file.Close()
				fileWrapper := DefaultFileWrapper{
					file,
					info,
					path,
				}
				fileList = append(fileList, &fileWrapper)
			} else {
				log.Println("Coudn't open the File")
				log.Fatal(openError)
			}
		}

		return nil
	})

	if errWalk != nil {
		log.Println("Couldn't Walk")
		log.Fatal(errWalk)
	}

	return &fileList
}

func BuildRegexFilterByExtension(argument string) *regexp.Regexp {
	//TODO: Implement auto regex build by extension passed in argument
	libRegEx, errRegex := regexp.Compile(`^.*\.(go|java|py|html|js|ts|kt)$`)

	if errRegex != nil {
		log.Fatal(errRegex)
	}

	return libRegEx
}

func CheckDirectory(directory string, closeOnFailure bool) error {
	dir, errDir := os.Stat(directory)
	if errDir != nil {
		if closeOnFailure {
			log.Fatal("failed to open directory, error: %w", errDir)
		}
	}
	if !dir.IsDir() {
		sErro := dir.Name() + " is not a directory"
		if closeOnFailure {
			log.Fatal(sErro)
		}
		errDir = errors.New(sErro)
	}

	return errDir
}
