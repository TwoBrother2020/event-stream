package event_stream

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"github.com/hashicorp/go-msgpack/codec"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Decode reverses the encode operation on a byte slice input
func decodeMsgPack(buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	hd := codec.MsgpackHandle{}
	dec := codec.NewDecoder(r, &hd)
	return dec.Decode(out)
}

// Encode writes an encoded object to a new bytes buffer
func encodeMsgPack(in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	hd := codec.MsgpackHandle{}
	enc := codec.NewEncoder(buf, &hd)
	err := enc.Encode(in)
	return buf, err
}

// Converts bytes to an integer
func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Converts a uint to a byte slice
func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

// 压缩目录下所有文件，不压缩子目录
func zipDir(dir string) (*bytes.Buffer, error) {
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

// 解压数据到指定目录
func unzip(dir string, snapshot io.Reader) error {
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
