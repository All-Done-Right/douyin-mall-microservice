package config

type Config struct {
	Mysql       Mysql
	Consul      Etcd
	Redis       Redis
	ServiceInfo ServiceInfo
	Logger      Logger
}
