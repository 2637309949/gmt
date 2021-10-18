package fs

import (
	"io"
	"net/http"
	"os"
	pathpkg "path"
	"path/filepath"
	"sort"
)

func openStat(fs http.FileSystem, name string) (http.File, os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return f, fi, nil
}

type WalkFilesFunc func(path string, info os.FileInfo, rs io.ReadSeeker, err error) error

func WalkFiles(fs http.FileSystem, root string, walkFn WalkFilesFunc) error {
	file, info, err := openStat(fs, root)
	if err != nil {
		return walkFn(root, nil, nil, err)
	}
	return walkFiles(fs, root, info, file, walkFn)
}

func ReadDir(fs http.FileSystem, name string) ([]os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdir(0)
}

func readDirNames(fs http.FileSystem, dirname string) ([]string, error) {
	fis, err := ReadDir(fs, dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(fis))
	for i := range fis {
		names[i] = fis[i].Name()
	}
	sort.Strings(names)
	return names, nil
}

func walkFiles(fs http.FileSystem, path string, info os.FileInfo, file http.File, walkFn WalkFilesFunc) error {
	err := walkFn(path, info, file, nil)
	file.Close()
	if err != nil {
		if info.IsDir() && err == filepath.SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}

	names, err := readDirNames(fs, path)
	if err != nil {
		return walkFn(path, info, nil, err)
	}

	for _, name := range names {
		filename := pathpkg.Join(path, name)
		file, fileInfo, err := openStat(fs, filename)
		if err != nil {
			if err := walkFn(filename, nil, nil, err); err != nil && err != filepath.SkipDir {
				return err
			}
		} else {
			err = walkFiles(fs, filename, fileInfo, file, walkFn)
			// file is closed by walkFiles, so we don't need to close it here.
			if err != nil {
				if !fileInfo.IsDir() || err != filepath.SkipDir {
					return err
				}
			}
		}
	}
	return nil
}
