package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func Extract(destination string, src string) error {
	z, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer z.Close()
	for _, file := range z.File {
		f, err := file.Open()
		if err != nil {
			return err
		}
		if !file.Mode().IsRegular() {
			continue
		}

		if err := writeNewFile(filepath.Join(destination, file.Name), f, file.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func writeNewFile(name string, in io.Reader, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(name), 0755)
	if err != nil {
		return err
	}
	dst, err := os.Create(name)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			fmt.Println("warning:", err)
			return nil
		}
		return err
	}
	defer dst.Close()
	if runtime.GOOS != "windows" {
		if err = dst.Chmod(mode); err != nil {
			return err
		}
	}
	if _, err = io.Copy(dst, in); err != nil {
		return err
	}
	return nil
}
