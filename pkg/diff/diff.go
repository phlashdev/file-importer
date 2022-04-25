package diff

import (
	"errors"
	"fmt"
	"os"
)

type directoryDiffer struct {
	src  string
	dest string
}

func NewDirectoryDiffer(src string, dest string) (directoryDiffer, error) {
	if src == "" {
		return directoryDiffer{}, errors.New("source directory must not be empty")
	}

	if dest == "" {
		return directoryDiffer{}, errors.New("destination directory must not be empty")
	}

	return directoryDiffer{
		src:  src,
		dest: dest,
	}, nil
}

func (d directoryDiffer) Diff() error {
	_, err := os.ReadDir(d.src)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("source directory does not exist: %s", d.src)
	} else if err != nil {
		return fmt.Errorf("error while reading source directory: %s", err)
	}

	_, err = os.ReadDir(d.dest)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("destination directory does not exist: %s", d.dest)
	} else if err != nil {
		return fmt.Errorf("error while reading destination directory: %s", err)
	}

	return nil
}
