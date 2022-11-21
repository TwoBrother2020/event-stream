package util

import (
	"archive/zip"
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Zip 压缩目录下所有文件，不压缩子目录
func Zip(dir string) (*bytes.Buffer, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		create, err := w.Create(info.Name())
		if err != nil {
			return nil, err
		}
		file, err := os.ReadFile(info.Name())
		if err != nil {
			return nil, err
		}
		create.Write(file)
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf, nil

}

// Unzip 解压数据到指定目录
func Unzip(dir string, snapshot io.Reader) error {
	data, err := io.ReadAll(snapshot)
	if err != nil {
		return err
	}

	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		content, err := io.ReadAll(rc)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(dir, f.Name), content, fs.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
