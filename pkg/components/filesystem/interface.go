package filesystem

import (
	"errors"
	"fmt"
	"goforge/pkg/components/result"
	"os"
)

type FileAlreadyExistsError struct {
	Message string
}

func (e FileAlreadyExistsError) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

func CreateDirectoryIfNotExists(path string) *result.Result[string] {
	if mkdirAllErr := os.MkdirAll(path, 0755); mkdirAllErr != nil {
		return result.ErrorResult[string](mkdirAllErr)
	} else {
		return result.Return(path)
	}
}

func CreateAndWriteToFileIfNotExists(path string, contents string) *result.Result[string] {
	if info, statErr := os.Stat(path); statErr != nil {
		if os.IsNotExist(statErr) {
			if file, openErr := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755); openErr != nil {
				return result.ErrorResult[string](openErr)
			} else {
				defer file.Close()
				_, err := file.WriteString(contents)
				if err != nil {
					return result.ErrorResult[string](fmt.Errorf("error writing string \"%s\" to file %s", contents, path))
				}
				return result.Return(path)
			}
		} else {
			return result.ErrorResult[string](fmt.Errorf("error characterizing existing file %s", statErr))
		}
	} else if info.IsDir() {
		return result.ErrorResult[string](fmt.Errorf("cannot create file %s; directory of name already exists", path))
	} else {
		return result.ErrorResult[string](FileAlreadyExistsError{Message: fmt.Sprintf("cannot write to file %s; file already exists", path)})
	}
}

func CreateFileIfNotExists(path string) *result.Result[string] {
	fileResult := CreateAndWriteToFileIfNotExists(path, "")
	if e := fileResult.Left(); e != nil {
		var fileAlreadyExistsError *FileAlreadyExistsError
		switch {
		case errors.As(e, &fileAlreadyExistsError):
			return result.Return(path)
		default:
			return fileResult
		}
	} else {
		return fileResult
	}
}
