package config

import "fmt"

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (m Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database)
}
