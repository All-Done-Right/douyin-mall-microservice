package config

import "strconv"

type Etcd struct {
	Host string
	Port int
}

func (e Etcd) Addr() string {
	return e.Host + ":" + strconv.Itoa(e.Port)
}
