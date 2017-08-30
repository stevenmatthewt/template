package template

import (
	"fmt"
	"os"
	"path/filepath"
)

type Function interface {
	File(path string) error
	Folder(path string) (ignore bool, err error)
}

// Walk walks the filesystem at `path` applying the Function
// for every file/folder. Additionally it provides the capability
// to stop walking down a particular folder if that folder is
// marked as ignored (coming soon)
//
// TODO: filepath.Walk is lexical, we need a bfs (?) order.
func (t Templater) Walk(path string, fn Function) error {
	filepath.Walk(path, t.newWalkFn(fn))

	return nil
}

func (t Templater) newWalkFn(fn Function) filepath.WalkFunc {
	return func(path string, info os.FileInfo, inputErr error) error {
		relativePath, err := filepath.Rel(t.source, path)
		if err != nil {
			fmt.Print(err)
			return err
		}
		if info.IsDir() {
			_, err := fn.Folder(relativePath)
			if err != nil {
				fmt.Print(err)
			}
			return err
		}

		err = fn.File(relativePath)
		if err != nil {
			fmt.Print(err)
		}
		return err
	}
}
