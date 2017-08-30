package template

import (
	"fmt"
	"os"
	"path/filepath"
)

func (t Templater) templateFolder(path string) error {
	templatedPath, err := t.finalPath(path)
	if templatedPath == "" {
		return nil
	}
	sourcePath := filepath.Join(t.source, path)
	destPath := filepath.Join(t.destination, templatedPath)
	if err != nil {
		return fmt.Errorf("failed to process template directory %s -- %s", path, err)
	}

	info, err := os.Stat(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %s", err)
	}

	//TODO: MkdirAll seems like a bad idea... Need some way to reduce the graph...
	err = os.MkdirAll(destPath, info.Mode())
	if err != nil {
		return fmt.Errorf("failed to create directory: %s", err)
	}
	return nil
}
