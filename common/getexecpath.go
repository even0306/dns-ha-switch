package common

import (
	"log/slog"
	"os"
	"path/filepath"
)

func GetExecPath() string {
	dir, err := os.Executable()
	if err != nil {
		slog.Error(err.Error())
	}

	dir, _ = filepath.Split(dir)
	slog.Debug(dir)
	return dir
}
