package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource string
	Etcd       string
	Processor  string
}
