package modules

import (
	"net"
	"time"
)

// 检测端口是否可通，传入“ip:port”，返回bool
func CheckPort(ip_port string) bool {
	conn, err := net.DialTimeout("tcp", ip_port, 3*time.Second)
	if err != nil {
		return false
	} else {
		if conn != nil {
			conn.Close()
			return true
		} else {
			return false
		}
	}
}
