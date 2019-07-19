package exec

import (
	"io"
	"os"
	"path/filepath"
)

func copyDir(src, dest string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dest, fi.Mode()); err != nil {
		return err
	}
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	infos, err := file.Readdir(0)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	for _, info := range infos {
		s := filepath.Join(src, info.Name())
		d := filepath.Join(dest, info.Name())
		if info.IsDir() {
			if err := copyDir(s, d); err != nil {
				return err
			}
			continue
		}
		if err := copyFile(s, d); err != nil {
			return err
		}
		continue
	}
	return nil
}

func copyFile(src, dest string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	return err
}
