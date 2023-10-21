package newutils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type ElementType int
type CopyrightType int

const (
	Component ElementType = iota
	Base
	App
	EPL2 CopyrightType = iota
)

func IsSafeName(input string) bool {
	pattern := "^[a-z0-9_]+$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

func ShouldNotCreateElement(entity string, elementType ElementType) error {
	if !IsSafeName(entity) {
		return fmt.Errorf("element name %s contains characters other than lowercase letters, numbers, or underscores", entity)
	} else {
		cleanEntityPath, cleanEntityPathErr := CreateCleanElementPath(entity, elementType)
		if cleanEntityPathErr != nil {
			return cleanEntityPathErr
		} else {
			info, infoErr := os.Stat(cleanEntityPath)
			if infoErr != nil {
				if os.IsNotExist(infoErr) {
					return nil
				} else {
					return infoErr
				}
			} else {
				if info.IsDir() {
					return fmt.Errorf("%s is a directory", cleanEntityPath)
				} else {
					return fmt.Errorf("%s is a file", cleanEntityPath)
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

func CreateNewElement(name string, elementType ElementType) error {
	cleanElementPath, cleanElementPathErr := CreateCleanElementPath(name, elementType)
	if cleanElementPathErr != nil {
		return cleanElementPathErr
	} else {
		switch elementType {
		case Component:
			createComponentOrBaseFolderAndInterfaceFiles(cleanElementPath, EPL2)
			createComponentOrBaseInternalFolderAndCoreFiles(cleanElementPath, EPL2)
		case Base:
			createComponentOrBaseFolderAndInterfaceFiles(cleanElementPath, EPL2)
			createComponentOrBaseInternalFolderAndCoreFiles(cleanElementPath, EPL2)
		case App:
			createAppFolderAndMainFile(cleanElementPath, EPL2)
		}
	}
	return fmt.Errorf("implement me")
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
		}
		return nil
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
			return fmt.Errorf("error creating interface files in %s", cleanElementPath)
		} else {
			copyrightComment := copyrightComment(copyrightType)
			packageLine := "package main"
			mainFile.WriteString(copyrightComment + packageLine)
		}
		return nil
	}
}

func copyrightComment(copyrightType CopyrightType) string {
	return ""
}
