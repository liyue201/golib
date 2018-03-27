package path

import (
	"os"
	"path/filepath"
	"strings"
	"os/exec"
)

func IsExist(strPath string) bool {
	_, err := os.Stat(strPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

var exeDir string

func GetExeDirPath() string {
	if exeDir != "" {
		return exeDir
	}

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	exeDir = filepath.Dir(path)

	return exeDir
}

func IsRelactivePath(path string) bool {
	if strings.Index(path, ".") == 0 {
		return true
	}
	return false
}

func AbsPath(path string) string {
	if IsRelactivePath(path) {
		path = GetExeDirPath() + string(os.PathSeparator) + path
		path, _ := filepath.Abs(path)
		return path
	}
	return path
}
