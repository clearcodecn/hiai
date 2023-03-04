package zipfile

import (
	"archive/zip"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var (
	mu sync.Mutex
)

func ZipFile(srcFile string, uid string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	tmpDir := os.TempDir()
	dir := filepath.Join(tmpDir, uuid.New().String()+".zip")
	c := fmt.Sprintf(conf, uid)

	ioutil.WriteFile(filepath.Join(srcFile, "config.json"), []byte(c), 0666)

	if err := compressPathToZip(srcFile, dir); err != nil {
		return "", err
	}
	return dir, nil
}

// compressPathToZip 压缩文件夹
func compressPathToZip(path, targetFile string) error {
	d, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = compress(f, "", w)
	return err
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if prefix == "" {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
