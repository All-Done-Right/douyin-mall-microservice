package config

import "strconv"

type ServiceInfo struct {
	Name string
	Host string
	Port int
}

func (i ServiceInfo) Addr() string {
	return i.Host + ":" + strconv.Itoa(i.Port)
}
