package walk

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
func Walk(path string, fn Function) error {
	filepath.Walk(path, newWalkFn(fn))

	return nil
}

func newWalkFn(fn Function) filepath.WalkFunc {
	return func(path string, info os.FileInfo, inputErr error) error {
		if info.IsDir() {
			_, err := fn.Folder(path)
			if err != nil {
				fmt.Print(err)
			}
			return err
		}

		err := fn.File(path)
		if err != nil {
			fmt.Print(err)
		}
		return err
	}
}
