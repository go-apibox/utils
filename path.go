package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	progPath := os.Args[0]
	if strings.IndexByte(progPath, '/') == -1 {
		// 在PATH路径列表中查找命令
		var err error
		progPath, err = exec.LookPath(progPath)
		if err != nil {
			return path
		}
	}
	progPath, _ = filepath.EvalSymlinks(progPath)
	progDir, _ := filepath.Abs(filepath.Dir(progPath))
	return filepath.Join(progDir, path)
}
