package file

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

func GetExt(filename string) string {
	return path.Ext(filename)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func IsNotExitMkDir(src string) error {
	if notExit := CheckNotExist(src); notExit {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}
