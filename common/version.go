package common

import (
	"log/slog"
)

func GetVersion() string {
	var x string
	var y string
	var z string
	x = "1"
	y = "0"
	z = "1"

	slog.Info("dns-api-ops", "version", x+"."+y+"."+z)
	return x + "." + y + "." + z
}
