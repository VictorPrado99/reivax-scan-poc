package util

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
)

type FileWrapper interface {
	GetFileContent() []string
	GetFileInfo() fs.FileInfo
	GetPath() string
}

type DefaultFileWrapper struct {
	FileContent []string
	FileInfo    fs.FileInfo
	Path        string
}

func (w *DefaultFileWrapper) GetFileContent() []string {
	return w.FileContent
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
			absPath, err := filepath.Abs(path)

			if err != nil {
				log.Println("Abs path error ", err)
			}

			textFiles := GetFileContent(absPath)

			fileWrapper := DefaultFileWrapper{
				textFiles,
				info,
				path,
			}
			fileList = append(fileList, &fileWrapper)
		}

		return nil
	})

	if errWalk != nil {
		log.Println("Couldn't Walk")
		log.Fatal(errWalk)
	}

	return &fileList
}

func GetFileContent(absPath string) []string {
	file, err := os.Open(absPath)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err == nil {
		log.Fatal("Couldn't read ", absPath, err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func BuildRegexFilterByExtension(arguments ...string) *regexp.Regexp {
	var extensionRegex string

	if len(arguments) == 0 {
		extensionRegex = "go|html|js"
	}

	for _, extension := range arguments {
		if len(extensionRegex) == 0 {
			extensionRegex = extensionRegex + extension
		} else {
			extensionRegex = extensionRegex + "|" + extension
		}
	}

	extensionRegex = `^.*\.(` + extensionRegex + `)$`

	libRegEx, errRegex := regexp.Compile(extensionRegex)

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

func ContainTypeInList[T any](value T, list *[]T) bool {
	for _, valueAlready := range *list {
		if reflect.TypeOf(valueAlready) == reflect.TypeOf(value) {
			return true
		}
	}
	return false
}
