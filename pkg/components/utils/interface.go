package utils

import (
	"fmt"
	"goforge/pkg/components/filesystem"
	"goforge/pkg/components/result"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ElementType int
type FileType int
type CopyrightType int

const (
	Component ElementType = iota
	Base
	App
)

const (
	EPL2 CopyrightType = iota
)

type SResult result.Result[string]

func (r *SResult) Bind(f result.MonadicFunction[string, string]) *SResult {
	returnedResult := result.Bind((*result.Result[string])(r), f)
	return (*SResult)(returnedResult)
}

func IsSafeName(input string) bool {
	allowedChars := "[a-z][a-z0-9_]+"
	pattern := "^" + allowedChars + "(/" + allowedChars + ")+" + "$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

func IsValidGoFileName(input string) bool {
	allowedChars := "[a-z][a-z0-9_]+"
	pattern := "^" + allowedChars + "(/" + allowedChars + ")+" + "(.go)?" + "$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

func ElementExists(name string, elementType ElementType) bool {
	createElementPath := func(name string) *result.Result[string] {
		return createElementPath(name, elementType)
	}
	doesElementExist := func(elementPath string) *result.Result[bool] {
		_, infoErr := os.Stat(elementPath)
		if infoErr != nil {
			if os.IsNotExist(infoErr) {
				return result.Return(false)
			} else {
				return result.ErrorResult[bool](infoErr)
			}
		} else {
			return result.Return(true)
		}
	}
	nameResult := result.Return(name)
	elementPathResult := result.Bind(nameResult, createElementPath)
	elementExistsResult := result.Bind(elementPathResult, doesElementExist)
	if elementExistsResult.Left() != nil {
		return true
	} else {
		return elementExistsResult.Right()
	}
}

func ShouldNotCreateElement(name string, elementType ElementType) error {
	if !IsSafeName(name) {
		return fmt.Errorf("element name %s contains characters other than lowercase letters, numbers, or underscores", name)
	} else {
		cleanElementPath, cleanElementPathErr := CreateCleanElementPath(name, elementType)
		if cleanElementPathErr != nil {
			return cleanElementPathErr
		} else {
			info, infoErr := os.Stat(cleanElementPath)
			if infoErr != nil {
				if os.IsNotExist(infoErr) {
					return nil
				} else {
					return infoErr
				}
			} else {
				if info.IsDir() {
					return fmt.Errorf("%s is a directory", cleanElementPath)
				} else {
					return fmt.Errorf("%s is a file", cleanElementPath)
				}
			}
		}
	}
}

func CreateCleanElementPath(name string, elementType ElementType) (string, error) {
	switch elementType {
	case Component:
		return filepath.Clean(filepath.Join("./pkg/components", name)), nil
	case Base:
		return filepath.Clean(filepath.Join("./pkg/bases", name)), nil
	case App:
		return filepath.Clean(filepath.Join("./cmd", name)), nil
	default:
		return "", fmt.Errorf("elementType must be of type Component, Base, or App")
	}
}

func createElementPath(name string, elementType ElementType) *result.Result[string] {
	if !IsSafeName(name) {
		return result.ErrorResult[string](fmt.Errorf("element: \"%s\" should start with a letter, contain only lowercase letters and numbers, and be a valid file path", name))
	}
	switch elementType {
	case Component:
		return result.Return(filepath.Clean(filepath.Join("./pkg/components", name)))
	case Base:
		return result.Return(filepath.Clean(filepath.Join("./pkg/bases", name)))
	case App:
		return result.Return(filepath.Clean(filepath.Join("./cmd", name)))
	default:
		return result.ErrorResult[string](fmt.Errorf("elementType must be of type Component, Base, or App"))
	}
}

func CreateNewElement(name string, elementType ElementType, copyrightType CopyrightType) error {
	switch elementType {
	case Component:
		return createComponentOrBase(name, Component, copyrightType)
	case Base:
		return createComponentOrBase(name, Base, copyrightType)
	case App:
		return createAppFolderAndMainFile(name, App, copyrightType)
	default:
		return fmt.Errorf("can only create element of type Component, Base, or App")
	}
}

func createComponentOrBase(name string, elementType ElementType, copyrightType CopyrightType) error {
	createElementPath := func(name string) *result.Result[string] {
		return createElementPath(name, elementType)
	}
	getDirPath := func(path string) *result.Result[string] {
		return result.Return(filepath.Dir(path))
	}
	mkdir := filesystem.CreateDirectoryIfNotExists
	mkwrite := filesystem.CreateAndWriteToFileIfNotExists
	interfaceContents := copyrightComment(copyrightType) + fmt.Sprintf("package %s", name)
	writeInterfaceFile := func(path string) *result.Result[string] {
		interfacePath := filepath.Join(path, "interface.go")
		return mkwrite(interfacePath, interfaceContents)
	}
	writeInterfaceTestFile := func(path string) *result.Result[string] {
		interfaceTestPath := filepath.Join(path, "interface_test.go")
		return mkwrite(interfaceTestPath, interfaceContents)
	}
	internalContents := copyrightComment(copyrightType) + "package internal"
	mkdirInternal := func(path string) *result.Result[string] {
		internalPath := filepath.Join(path, "internal")
		return mkdir(internalPath)
	}
	writeCoreFile := func(path string) *result.Result[string] {
		coreFilePath := filepath.Join(path, "internal", "core.go")
		return mkwrite(coreFilePath, internalContents)
	}
	writeCoreTestFile := func(path string) *result.Result[string] {
		coreFilePath := filepath.Join(path, "internal", "core_test.go")
		return mkwrite(coreFilePath, internalContents)
	}
	finalResult := (*SResult)(result.Return(name)).
		Bind(createElementPath).
		Bind(mkdir).
		Bind(writeInterfaceFile).
		Bind(getDirPath).
		Bind(writeInterfaceTestFile).
		Bind(getDirPath).
		Bind(mkdirInternal).
		Bind(getDirPath).
		Bind(writeCoreFile).
		Bind(getDirPath).
		Bind(getDirPath).
		Bind(writeCoreTestFile).
		Bind(getDirPath).
		Bind(getDirPath)
	return (*result.Result[string])(finalResult).Left()
}

func createAppFolderAndMainFile(name string, elementType ElementType, copyrightType CopyrightType) error {
	createElementPath := func(name string) *result.Result[string] {
		return createElementPath(name, elementType)
	}
	getDirPath := func(path string) *result.Result[string] {
		return result.Return(filepath.Dir(path))
	}
	mkdir := filesystem.CreateDirectoryIfNotExists
	mkwrite := filesystem.CreateAndWriteToFileIfNotExists
	contents := copyrightComment(copyrightType) + "package main"
	writeMainFile := func(path string) *result.Result[string] {
		mainFilePath := filepath.Join(path, "main.go")
		return mkwrite(mainFilePath, contents)
	}
	finalResult := (*SResult)(result.Return(name)).
		Bind(createElementPath).
		Bind(mkdir).
		Bind(writeMainFile).
		Bind(getDirPath)
	return (*result.Result[string])(finalResult).Left()
}

func CreateInternalFile(fileName string, elementName string, elementType ElementType, copyrightType CopyrightType) error {
	createElementPath := func(elementName string) *result.Result[string] {
		return createElementPath(elementName, elementType)
	}
	mkdir := filesystem.CreateDirectoryIfNotExists
	mkwrite := filesystem.CreateAndWriteToFileIfNotExists
	mkdirInternal := func(path string) *result.Result[string] {
		internalPath := filepath.Join(path, "internal")
		return mkdir(internalPath)
	}
	writeInternalFile := func(internalPath string) *result.Result[string] {
		if !IsValidGoFileName(fileName) {
			return result.ErrorResult[string](fmt.Errorf("element: \"%s\" should be a valid path to a go file", fileName))
		}
		if !strings.HasSuffix(fileName, ".go") {
			fileName = fileName + ".go"
		}
		internalFilePath := filepath.Join(internalPath, fileName)
		internalContents := copyrightComment(copyrightType) + "package internal"
		return mkwrite(internalFilePath, internalContents)
	}
	finalResult := (*SResult)(result.Return(elementName)).
		Bind(createElementPath).
		Bind(mkdirInternal).
		Bind(writeInternalFile)
	return (*result.Result[string])(finalResult).Left()
}

func copyrightComment(copyrightType CopyrightType) string {
	return ""
}

func RemoveElement(name string, elementType ElementType) error {
	cleanElementPath, cleanElementPathErr := CreateCleanElementPath(name, elementType)
	if cleanElementPathErr != nil {
		return cleanElementPathErr
	} else {
		return os.RemoveAll(cleanElementPath)
	}
}
