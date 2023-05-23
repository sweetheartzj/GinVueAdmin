package config

type Redis struct {
	Addr     string `default:"127.0.0.1:6379" json:"addr" yaml:"addr"`
	DB       int    `default:"0" json:"DB" yaml:"db"`
	User     string `default:"default" json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	IP       string `default:"127.0.0.1" json:"ip" yaml:"ip"`
	Port     string `default:"6379" json:"port" yaml:"port"`
}
