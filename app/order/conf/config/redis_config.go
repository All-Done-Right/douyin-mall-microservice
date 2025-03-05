package config

import (
	"strconv"
)

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

func (r Redis) Addr() string {
	return r.Host + ":" + strconv.Itoa(r.Port)
}
