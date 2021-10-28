package fs

import (
	"io/ioutil"
	"net/http"
)

// ReadFile reads the file named by path from fs and returns the contents.
func ReadFile(fs http.FileSystem, path string) ([]byte, error) {
	rc, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}
