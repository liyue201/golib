package xpath

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func IsRelativePath(path string) bool {
	if strings.Index(path, ".") == 0 {
		return true
	}
	return false
}

func AbsPath(path string) string {
	if IsRelativePath(path) {
		path = filepath.Join(GetExeDirPath(), path)
		path, _ := filepath.Abs(path)
		return path
	}
	return path
}

func Ext(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if len(ext) > 0 {
		ext = ext[1:]
	}
	return ext
}
