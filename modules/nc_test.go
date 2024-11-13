package modules_test

import (
	"testing"

	"github.com/dns_api_ops/modules"
)

func TestCheckPorts(t *testing.T) {
	abc := "127.0.0.1:80"
	modules.CheckPort(abc)
}
