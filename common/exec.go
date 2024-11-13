package common

import (
	"bytes"
	"os/exec"
	"strings"
)

type ExecAPI interface {
	// 脚本判断服务是否正常。返回状态0为正常，1为异常。
	Bash() error
}

type ExecShell struct {
	Script string
}

func NewExecShell(s string) *ExecShell {
	return &ExecShell{
		Script: s,
	}
}

func (e *ExecShell) Bash(ip string) string {
	e.Script = strings.Replace(e.Script, "IPADDR", ip, -1)
	cmd := exec.Command("cmd", "/C", e.Script)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "1"
	}

	if stderr.String() != "" {
		return "1"
	}

	return "0"
}
