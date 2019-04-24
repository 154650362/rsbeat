// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period     time.Duration `config:"period"`
	Redis      []string      `config:"redis"`
	SlowerThan int           `config:"slowerThan"`
	SlowerLen  int           `config:"slowerLen"`
}

var DefaultConfig = Config{
	Period:     1 * time.Second,
	Redis:      []string{"127.0.0.1:6379"},
	SlowerThan: 100, //ms
	SlowerLen: 1000,
}
