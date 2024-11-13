package logging

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/dns_api_ops/common"
)

var logFile *os.File

func NewLogger(level string) {
	dir := common.GetExecPath()
	logFile, err := os.OpenFile(filepath.Join(dir)+"/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error(err.Error())
	}

	var slogLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warnning":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		slog.Error("", "Panic", "不支持的日志等级")
	}

	logger := slog.New(slog.NewTextHandler(io.MultiWriter(os.Stdout, logFile), &slog.HandlerOptions{Level: slogLevel}))
	slog.SetDefault(logger)
}

func LogFileClose() {
	logFile.Close()
}
