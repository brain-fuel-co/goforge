package newcomponent

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func isSafe(input string) bool {
	pattern := "^[a-z0-9_]+$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

func createCleanComponentPath(name string) string {
	componentPath := filepath.Join("./pkg/components", name)
	return filepath.Clean(componentPath)
}

func ShouldNotCreateComponent(name string) error {
	if !isSafe(name) {
		return fmt.Errorf("component name %s contains characters other than lowercase letters, numbers, or underscores", name)
	} else {
		cleanComponentPath := createCleanComponentPath(name)
		info, infoErr := os.Stat(cleanComponentPath)
		if infoErr != nil {
			if os.IsNotExist(infoErr) {
				return nil
			} else {
				return infoErr
			}
		} else {
			if info.IsDir() {
				return fmt.Errorf("%s is a directory", cleanComponentPath)
			} else {
				return fmt.Errorf("%s is a file", cleanComponentPath)
			}
		}
	}
}

func CreateNewComponent(name string) error {
	shouldNotCreateComponent := ShouldNotCreateComponent(name)
	if shouldNotCreateComponent != nil {
		return shouldNotCreateComponent
	} else {
		cleanComponentPath := createCleanComponentPath(name)
		cleanComponentInternalPath := filepath.Join(cleanComponentPath, "internal")
		if directoryErr := os.MkdirAll(cleanComponentInternalPath, 0755); directoryErr != nil {
			return directoryErr
		} else {
			interfaceFilePath := filepath.Join(cleanComponentPath, "interface.go")
			interfaceTestFilePath := filepath.Join(cleanComponentPath, "interface_test.go")
			internalCoreFilePath := filepath.Join(cleanComponentInternalPath, "core.go")
			internalCoreTestFilePath := filepath.Join(cleanComponentInternalPath, "core_test.go")
			interfaceFile, interfaceFileErr := os.Create(interfaceFilePath)
			defer interfaceFile.Close()
			interfaceTestFile, interfaceTestFileErr := os.Create(interfaceTestFilePath)
			defer interfaceTestFile.Close()
			internalCoreFile, internalCoreFileErr := os.Create(internalCoreFilePath)
			defer internalCoreFile.Close()
			internalCoreTestFile, internalCoreTestFileErr := os.Create(internalCoreTestFilePath)
			defer internalCoreTestFile.Close()
			if interfaceFileErr != nil ||
				interfaceTestFileErr != nil ||
				internalCoreFileErr != nil ||
				internalCoreTestFileErr != nil {
				removeErr := os.RemoveAll(cleanComponentPath)
				if removeErr != nil {
					return removeErr
				} else {
					return fmt.Errorf("error creating all files for component %s", name)
				}
			} else {
				copyrightComment := ""
				packageLine := fmt.Sprintf("package %s\n", name)
				internalPackageLine := "package internal\n"
				interfaceFile.WriteString(copyrightComment + packageLine)
				interfaceTestFile.WriteString(copyrightComment + packageLine)
				internalCoreFile.WriteString(copyrightComment + internalPackageLine)
				internalCoreTestFile.WriteString(copyrightComment + internalPackageLine)
				return nil
			}
		}
	}
}
