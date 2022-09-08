package config

import "fmt"

var Config *Cfg

type Cfg struct {
	Server struct {
		Host string
		Port int
	}
	Cache struct {
		Addr string
		Pswd string
	}
	Mysql struct {
		Dsn string
	}
	Email struct {
		Name     string
		SmtpHost string
		SmtpPort int
		SmtpUser string
		SmtpPswd string
		Count    int
	}
}

func (cfg Cfg) ServerAddress() string {
	return fmt.Sprintf("%s:%v", cfg.Server.Host, cfg.Server.Port)
}
