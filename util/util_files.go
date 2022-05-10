package util

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func GetFiles(directory string, libRegEx *regexp.Regexp) *[]fs.FileInfo {
	var fileList []fs.FileInfo
	errWalk := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if libRegEx.MatchString(info.Name()) {
			fileList = append(fileList, info)
		}

		return nil
	})

	if errWalk != nil {
		log.Fatal(errWalk)
	}

	return &fileList
}

func BuildRegexFilterByExtension(argument string) *regexp.Regexp {
	//TODO: Implement auto regex build by extension passed in argument
	libRegEx, errRegex := regexp.Compile(`^.*\.(go|java|py)$`)

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