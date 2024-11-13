package config

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var C *viper.Viper

func SetConfig(filePath string) error {
	c := viper.New()
	dir, name := filepath.Split(filePath)
	c.SetConfigName(name)
	c.SetConfigType("yaml")

	c.AddConfigPath(dir)
	err := c.ReadInConfig()
	if err != nil {
		slog.Error(err.Error())
	}

	c.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed and Reloaded:", e.Name)
		err := c.ReadInConfig()
		if err != nil {
			slog.Error(err.Error())
		}
	})
	c.WatchConfig()
	C = c
	return nil
}
