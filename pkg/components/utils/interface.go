package utils

import (
	"fmt"
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

const (
	DoesNotExist FileType = iota
	Directory
	File
)

func IsSafeName(input string) bool {
	pattern := "^[a-z0-9_]+$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

func ElementExists(name string, elementType ElementType) bool {
	cleanElementPath, cleanElementPathErr := CreateCleanElementPath(name, elementType)
	if cleanElementPathErr != nil {
		return false
	} else {
		info, infoErr := os.Stat(cleanElementPath)
		if infoErr != nil {
			if os.IsNotExist(infoErr) {
				return false
			} else {
				if !info.IsDir() {
					return false
				} else {
					return true
				}
			}
		} else {
			return true
		}
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

func CreateNewElement(name string, elementType ElementType, copyrightType CopyrightType) error {
	cleanElementPath, cleanElementPathErr := CreateCleanElementPath(name, elementType)
	if cleanElementPathErr != nil {
		return cleanElementPathErr
	} else {
		switch elementType {
		case Component:
			return createComponentOrBase(name, Component, copyrightType)
		case Base:
			return createComponentOrBase(name, Base, copyrightType)
		case App:
			return createAppFolderAndMainFile(cleanElementPath, copyrightType)
		default:
			return fmt.Errorf("can only create element of type Component, Base, or App")
		}
	}
}

func createComponentOrBase(name string, elementType ElementType, copyrightType CopyrightType) error {
	cleanElementPath, cleanElementPathError := CreateCleanElementPath(name, elementType)
	if cleanElementPathError != nil {
		return fmt.Errorf("error")
	} else {
		elemErr := createComponentOrBaseFolderAndInterfaceFiles(cleanElementPath, copyrightType)
		if elemErr != nil {
			return elemErr
		} else {
			elemInternalErr := createComponentOrBaseInternalFolderAndCoreFiles(cleanElementPath, copyrightType)
			if elemInternalErr != nil {
				return elemInternalErr
			} else {
				return nil
			}
		}
	}
}

func createComponentOrBaseFolderAndInterfaceFiles(cleanElementPath string, copyrightType CopyrightType) error {
	if directoryErr := os.MkdirAll(cleanElementPath, 0755); directoryErr != nil {
		return directoryErr
	} else {
		interfaceFilePath := filepath.Join(cleanElementPath, "interface.go")
		interfaceFile, interfaceFileErr := os.Create(interfaceFilePath)
		defer interfaceFile.Close()
		interfaceTestFilePath := filepath.Join(cleanElementPath, "interface_test.go")
		interfaceTestFile, interfaceTestFileErr := os.Create(interfaceTestFilePath)
		defer interfaceTestFile.Close()
		if interfaceFileErr != nil || interfaceTestFileErr != nil {
			return fmt.Errorf("error creating interface files in %s", cleanElementPath)
		} else {
			copyrightComment := copyrightComment(copyrightType)
			packageLine := fmt.Sprintf("package %s\n", filepath.Base(cleanElementPath))
			interfaceFile.WriteString(copyrightComment + packageLine)
			interfaceTestFile.WriteString(copyrightComment + packageLine)
			return nil
		}
	}
}

func createComponentOrBaseInternalFolderAndCoreFiles(cleanElementPath string, copyrightType CopyrightType) error {
	cleanElementInternalPath := filepath.Join(cleanElementPath, "internal")
	if directoryErr := os.MkdirAll(cleanElementInternalPath, 0755); directoryErr != nil {
		return directoryErr
	} else {
		internalCoreFilePath := filepath.Join(cleanElementInternalPath, "core.go")
		internalCoreTestFilePath := filepath.Join(cleanElementInternalPath, "core_test.go")
		internalCoreFile, internalCoreFileErr := os.Create(internalCoreFilePath)
		defer internalCoreFile.Close()
		internalCoreTestFile, internalCoreTestFileErr := os.Create(internalCoreTestFilePath)
		defer internalCoreTestFile.Close()
		if internalCoreFileErr != nil ||
			internalCoreTestFileErr != nil {
			return fmt.Errorf("error creating core files in %s", cleanElementInternalPath)
		} else {
			copyrightComment := copyrightComment(copyrightType)
			internalPackageLine := "package internal\n"
			internalCoreFile.WriteString(copyrightComment + internalPackageLine)
			internalCoreTestFile.WriteString(copyrightComment + internalPackageLine)
			return nil
		}
	}
}

func createAppFolderAndMainFile(cleanElementPath string, copyrightType CopyrightType) error {
	if directoryErr := os.MkdirAll(cleanElementPath, 0755); directoryErr != nil {
		return directoryErr
	} else {
		mainFilePath := filepath.Join(cleanElementPath, "main.go")
		mainFile, mainFileErr := os.Create(mainFilePath)
		defer mainFile.Close()
		if mainFileErr != nil {
			return fmt.Errorf("error creating main file in %s", cleanElementPath)
		} else {
			copyrightComment := copyrightComment(copyrightType)
			packageLine := "package main"
			mainFile.WriteString(copyrightComment + packageLine)
			return nil
		}
	}
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

func CreateInternalFile(fileName string, elementName string, elementType ElementType, copyrightType CopyrightType) error {
	if !IsSafeName(fileName) {
		return fmt.Errorf("name %s contains characters other than lowercase letters, numbers, or underscores", fileName)
	} else {
		cleanElementPath, cleanElementPathErr := CreateCleanElementPath(elementName, elementType)
		if cleanElementPathErr != nil {
			return cleanElementPathErr
		} else {
			if !strings.HasSuffix(fileName, ".go") {
				fileName = fileName + ".go"
			}
			internalFolderPath := filepath.Join(cleanElementPath, "internal")
			mkInternalDirErr := os.MkdirAll(internalFolderPath, 0755)
			if mkInternalDirErr != nil {
				return mkInternalDirErr
			} else {
				newInternalFilePath := filepath.Join(internalFolderPath, fileName)
				newInternalFile, newInternalFileErr := os.Create(newInternalFilePath)
				defer newInternalFile.Close()
				if newInternalFileErr != nil {
					return fmt.Errorf("error creating %s file in %s", fileName, internalFolderPath)
				} else {
					copyrightComment := copyrightComment(copyrightType)
					packageLine := "package internal"
					newInternalFile.WriteString(copyrightComment + packageLine)
					return nil
				}
			}
		}
	}
}
