package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func (t Templater) templateFolder(path string) error {
	destPath := filepath.Join(t.destination, path)
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %s", err)
	}
	err = os.Mkdir(destPath, info.Mode())
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}
	return nil
}
