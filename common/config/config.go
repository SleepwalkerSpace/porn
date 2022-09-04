package config

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
