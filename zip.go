package quickzip

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Zipper struct {
	Path string //content path to zip
}

type IZipper interface {
	Zip(filename string) error
	walk(writer *zip.Writer) fs.WalkDirFunc
}

func NewZipper(dir string) *Zipper {
	return &Zipper{
		Path: dir,
	}
}

// compresses the content found in the struct.Path into filename on the filesystem, returns any errors
// and the time to zip
func (z *Zipper) Zip(filename string) error {
	//create compressed file
	nf, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer nf.Close()

	w := zip.NewWriter(nf)
	defer w.Close()

	//walk directories and files to compress
	err = filepath.WalkDir(z.Path, z.walk(w))
	if err != nil {
		return err
	}

	return nil
}

func (z *Zipper) walk(writer *zip.Writer) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		zf, err := writer.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(zf, file)
		if err != nil {
			return err
		}

		return nil
	}
}
