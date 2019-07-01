/**
 * @author XieKong
 * @date   2019/7/1 11:16
 */
package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (z *ZipEntry) loadClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.absPath)

	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()

			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()

			data, err := ioutil.ReadAll(rc)

			if err != nil {
				return nil, nil, err
			}

			return data, z, nil
		}
	}

	return nil, nil, errors.New("classpath not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
